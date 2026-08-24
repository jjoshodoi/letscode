package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"letscode/project-01-url-shortener/internal/logging"
	"letscode/project-01-url-shortener/internal/store"
)

type Handler struct {
	store store.Store
	rnd   *rand.Rand
}

func NewHandler(s store.Store) *Handler {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Handler{store: s, rnd: r}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", h.shorten)
	mux.HandleFunc("/health", h.health)
	mux.HandleFunc("/", h.redirect)
	return mux
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Code string `json:"code"`
}

func (h *Handler) shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json payload"})
		return
	}
	// Validate presence
	if strings.TrimSpace(req.URL) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "url is required"})
		return
	}
	// Normalize and validate the URL
	norm, err := NormalizeURL(req.URL)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	// Idempotency: if URL already exists, return existing code
	if existingCode, ok := h.store.FindByURL(norm); ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(shortenResponse{Code: existingCode})
		return
	}
	// Try generating random codes with bounded retries to avoid collisions
	const maxRandomAttempts = 5
	for i := 0; i < maxRandomAttempts; i++ {
		code := h.generateCode(6)
		if err := h.store.Save(code, norm); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(shortenResponse{Code: code})
			return
		} else {
			// collision, try again
			logging.Debugf("code generation collision (attempt %d): %v", i+1, err)
			continue
		}
	}
	// Fallback deterministic approach: base62(hash(url)+counter)
	for c := 0; c < 1000; c++ {
		code := deterministicCode(norm, c, 6)
		if err := h.store.Save(code, norm); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(shortenResponse{Code: code})
			return
		} else {
			// if url already exists, return its code (race)
			if existing, ok := h.store.FindByURL(norm); ok {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(shortenResponse{Code: existing})
				return
			}
			// otherwise continue trying
			continue
		}
	}
	logging.Errorf("failed to obtain unique code for url: %s", norm)
	w.WriteHeader(http.StatusInternalServerError)
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (h *Handler) redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	code := r.URL.Path
	if len(code) > 0 && code[0] == '/' {
		code = code[1:]
	}
	if code == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("URL Shortener: POST /shorten with JSON {\"url\": \"https://...\"}"))
		return
	}
	if dest, ok := h.store.Lookup(code); ok {
		http.Redirect(w, r, dest, http.StatusFound)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (h *Handler) generateCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[h.rnd.Intn(len(letters))]
	}
	return string(b)
}

// deterministicCode produces a base62-looking code derived from sha256(url:counter)
func deterministicCode(url string, counter, n int) string {
	h := sha256.Sum256([]byte(fmt.Sprintf("%s:%d", url, counter)))
	// convert to big.Int for base conversion
	var bi big.Int
	bi.SetBytes(h[:])
	out := make([]byte, n)
	base := big.NewInt(int64(len(letters)))
	for i := 0; i < n; i++ {
		var mod big.Int
		bi.DivMod(&bi, base, &mod)
		out[i] = letters[mod.Int64()]
	}
	return string(out)
}

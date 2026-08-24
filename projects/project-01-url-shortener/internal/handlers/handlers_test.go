package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"letscode/project-01-url-shortener/internal/store"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		in      string
		want    string
		wantErr bool
	}{
		{"example.com", "https://example.com", false},
		{"http://example.com/path", "http://example.com/path", false},
		{"https://example.com", "https://example.com", false},
		{"file:///etc/passwd", "", true},
		{"", "", true},
		{"http://", "", true},
	}
	for _, tc := range tests {
		got, err := NormalizeURL(tc.in)
		if (err != nil) != tc.wantErr {
			t.Fatalf("NormalizeURL(%q) error = %v, wantErr=%v", tc.in, err, tc.wantErr)
		}
		if err == nil && got != tc.want {
			t.Fatalf("NormalizeURL(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestShortenHandler_Validation(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "urls_handlers_test.json")
	defer os.Remove(tmp)
	st := store.NewFileStore(tmp)
	h := NewHandler(st)
	server := h.Routes()

	cases := []struct {
		name      string
		body      interface{}
		wantCode  int
		wantError bool
	}{
		{"valid_missing_scheme", map[string]string{"url": "example.com"}, http.StatusOK, false},
		{"missing_field", map[string]string{}, http.StatusBadRequest, true},
		{"empty_url", map[string]string{"url": ""}, http.StatusBadRequest, true},
		{"invalid_scheme", map[string]string{"url": "file:///etc/passwd"}, http.StatusBadRequest, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b, _ := json.Marshal(c.body)
			req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)
			if rr.Code != c.wantCode {
				t.Fatalf("status: got %d want %d; body=%s", rr.Code, c.wantCode, rr.Body.String())
			}
			if c.wantError {
				var m map[string]string
				if err := json.NewDecoder(rr.Body).Decode(&m); err != nil {
					t.Fatalf("failed to decode error body: %v", err)
				}
				if _, ok := m["error"]; !ok {
					t.Fatalf("expected error field, got: %v", m)
				}
			} else {
				var resp map[string]string
				if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
					t.Fatalf("decode response: %v; body=%s", err, rr.Body.String())
				}
				code, ok := resp["code"]
				if !ok || len(code) == 0 {
					t.Fatalf("expected code in response, got: %v", resp)
				}
			}
		})
	}
}

func TestShortenIdempotency(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "urls_handlers_idempotency.json")
	defer os.Remove(tmp)
	st := store.NewFileStore(tmp)
	h := NewHandler(st)
	server := h.Routes()

	payload := map[string]string{"url": "example.com"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("first request failed: %d body=%s", rr.Code, rr.Body.String())
	}
	var resp1 map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp1); err != nil {
		t.Fatalf("decode first response: %v", err)
	}
	code1, ok := resp1["code"]
	if !ok || code1 == "" {
		t.Fatalf("first response missing code: %v", resp1)
	}

	// Second request
	rr2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(b))
	req2.Header.Set("Content-Type", "application/json")
	server.ServeHTTP(rr2, req2)
	if rr2.Code != http.StatusOK {
		t.Fatalf("second request failed: %d body=%s", rr2.Code, rr2.Body.String())
	}
	var resp2 map[string]string
	if err := json.NewDecoder(rr2.Body).Decode(&resp2); err != nil {
		t.Fatalf("decode second response: %v", err)
	}
	code2, ok := resp2["code"]
	if !ok || code2 == "" {
		t.Fatalf("second response missing code: %v", resp2)
	}
	if code1 != code2 {
		t.Fatalf("expected idempotent code; got %s and %s", code1, code2)
	}
}

func TestRedirectHandler(t *testing.T) {
	st := store.NewFileStore(filepath.Join(t.TempDir(), "urls.json"))
	if err := st.Save("abc123", "https://example.com/path"); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	rr := httptest.NewRecorder()
	NewHandler(st).Routes().ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/abc123", nil))
	if rr.Code != http.StatusFound {
		t.Fatalf("status: got %d want %d", rr.Code, http.StatusFound)
	}
	if location := rr.Header().Get("Location"); location != "https://example.com/path" {
		t.Fatalf("Location: got %q", location)
	}
}

func TestHandlers_MethodNotAllowedAndNotFound(t *testing.T) {
	server := NewHandler(store.NewFileStore(filepath.Join(t.TempDir(), "urls.json"))).Routes()

	methodRR := httptest.NewRecorder()
	server.ServeHTTP(methodRR, httptest.NewRequest(http.MethodGet, "/shorten", nil))
	if methodRR.Code != http.StatusMethodNotAllowed {
		t.Fatalf("GET /shorten status: got %d want %d", methodRR.Code, http.StatusMethodNotAllowed)
	}

	notFoundRR := httptest.NewRecorder()
	server.ServeHTTP(notFoundRR, httptest.NewRequest(http.MethodGet, "/missing", nil))
	if notFoundRR.Code != http.StatusNotFound {
		t.Fatalf("GET /missing status: got %d want %d", notFoundRR.Code, http.StatusNotFound)
	}
}

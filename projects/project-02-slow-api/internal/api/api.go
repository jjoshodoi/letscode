package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"letscode/project-02-slow-api/internal/store"
)

func New(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", products(db))
	mux.HandleFunc("/products/", product(db))
	return mux
}

func products(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		items, err := store.FindProducts(db, r.URL.Query().Get("category"), r.URL.Query().Get("search"))
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		writeJSON(w, map[string]any{"products": items, "count": len(items)})
	}
}

func product(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/products/"))
		if err != nil || id < 1 {
			http.Error(w, "invalid product id", http.StatusBadRequest)
			return
		}
		item, err := store.FindProduct(db, id)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "product not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		writeJSON(w, item)
	}
}

func writeJSON(w http.ResponseWriter, value any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(value)
}

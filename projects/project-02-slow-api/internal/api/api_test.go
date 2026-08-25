package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"letscode/project-02-slow-api/internal/api"
	"letscode/project-02-slow-api/internal/store"
)

func TestProductsEndpoints(t *testing.T) {
	db, err := store.Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err := store.Seed(db, 3); err != nil {
		t.Fatal(err)
	}
	handler := api.New(db)

	for _, tc := range []struct {
		path string
		code int
	}{
		{"/products", http.StatusOK},
		{"/products?category=books", http.StatusOK},
		{"/products/1", http.StatusOK},
		{"/products/999", http.StatusNotFound},
	} {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, tc.path, nil))
		if rec.Code != tc.code {
			t.Errorf("%s: got %d, want %d", tc.path, rec.Code, tc.code)
		}
	}
}

func BenchmarkListProducts(b *testing.B) {
	db, err := store.Open(":memory:")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()
	if err := store.Seed(db, 10000); err != nil {
		b.Fatal(err)
	}
	handler := api.New(db)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/products?search=learning", nil))
	}
}

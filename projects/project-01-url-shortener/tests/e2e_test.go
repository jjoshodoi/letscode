package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"letscode/project-01-url-shortener/internal/handlers"
	"letscode/project-01-url-shortener/internal/store"
)

func TestShortenAndRedirectEndToEnd(t *testing.T) {
	server := httptest.NewServer(
		handlers.NewHandler(store.NewFileStore(filepath.Join(t.TempDir(), "urls.json"))).Routes(),
	)
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+"/shorten",
		bytes.NewBufferString(`{"url":"https://example.com/e2e"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("POST /shorten: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("POST /shorten status: got %d", resp.StatusCode)
	}

	var result struct {
		Code string `json:"code"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("decode shorten response: %v", err)
	}
	if result.Code == "" {
		t.Fatal("shorten response contained no code")
	}

	client := &http.Client{CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	redirect, err := client.Get(server.URL + "/" + result.Code)
	if err != nil {
		t.Fatalf("GET /%s: %v", result.Code, err)
	}
	defer redirect.Body.Close()
	if redirect.StatusCode != http.StatusFound {
		t.Fatalf("GET /%s status: got %d", result.Code, redirect.StatusCode)
	}
	if got := redirect.Header.Get("Location"); got != "https://example.com/e2e" {
		t.Fatalf("redirect Location: got %q", got)
	}
}

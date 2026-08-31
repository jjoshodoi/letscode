package tests

import (
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"

	"letscode/project-01-url-shortener/internal/store"
)

func TestFileStore_SaveLookup(t *testing.T) {
	tmp := filepath.Join(t.TempDir(), "urls.json")
	fs := store.NewFileStore(tmp)
	if err := fs.Save("abc123", "https://example.com"); err != nil {
		t.Fatalf("save failed: %v", err)
	}
	if u, ok := fs.Lookup("abc123"); !ok || u != "https://example.com" {
		t.Fatalf("lookup failed: got %v, ok=%v", u, ok)
	}
}

func TestFileStore_SaveAndLookupConcurrently(t *testing.T) {
	fs := store.NewFileStore(filepath.Join(t.TempDir(), "urls.json"))
	const workers = 100
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		i := i
		go func() {
			defer wg.Done()
			code := "code" + strconv.Itoa(i)
			url := "https://example.com/" + strconv.Itoa(i)
			if err := fs.Save(code, url); err != nil {
				t.Errorf("Save(%q): %v", code, err)
				return
			}
			if got, ok := fs.Lookup(code); !ok || got != url {
				t.Errorf("Lookup(%q) = %q, %v; want %q, true", code, got, ok, url)
			}
		}()
	}
	wg.Wait()

	all, err := fs.All()
	if err != nil {
		t.Fatalf("All: %v", err)
	}
	if len(all) != workers {
		t.Fatalf("All returned %d mappings, want %d", len(all), workers)
	}
}

func TestFileStore_RapidConcurrentSavesDoNotDeadlock(t *testing.T) {
	fs := store.NewFileStore(filepath.Join(t.TempDir(), "urls.json"))
	const workers = 250
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		i := i
		go func() {
			defer wg.Done()
			_ = fs.Save("rapid"+strconv.Itoa(i), "https://example.com/rapid/"+strconv.Itoa(i))
		}()
	}
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("concurrent Save calls did not complete")
	}
}

func TestSQLiteStore_SaveLookupAndUniqueness(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "urls.db")
	st, err := store.NewSQLiteStore(dbPath)
	if err != nil {
		t.Fatalf("NewSQLiteStore: %v", err)
	}
	defer st.Close()

	if err := st.Save("abc123", "https://example.com"); err != nil {
		t.Fatalf("save: %v", err)
	}
	if got, ok := st.Lookup("abc123"); !ok || got != "https://example.com" {
		t.Fatalf("lookup: got %q, ok=%v; want %q, true", got, ok, "https://example.com")
	}
	if got, ok := st.FindByURL("https://example.com"); !ok || got != "abc123" {
		t.Fatalf("FindByURL: got %q, ok=%v; want %q, true", got, ok, "abc123")
	}

	if err := st.Save("abc123", "https://example.com/other"); err == nil {
		t.Fatal("expected duplicate code error")
	}
	if err := st.Save("xyz999", "https://example.com"); err == nil {
		t.Fatal("expected duplicate URL error")
	}
	if err := st.Save("abc123", "https://example.com"); err != nil {
		t.Fatalf("idempotent save should succeed: %v", err)
	}
}

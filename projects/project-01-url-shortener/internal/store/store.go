package store

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"letscode/project-01-url-shortener/internal/model"
)

// Store is a minimal persistence interface for URL mappings.
type Store interface {
	Save(code, url string) error
	Lookup(code string) (string, bool)
	All() ([]model.URLMapping, error)
	FindByURL(url string) (string, bool)
}

// FileStore implements Store using a single JSON file.
type FileStore struct {
	path string
	mu   sync.RWMutex
	m    map[string]string // code -> url
	rev  map[string]string // url -> code
}

// NewFileStore creates or loads a file-backed store.
func NewFileStore(path string) *FileStore {
	fs := &FileStore{path: path, m: make(map[string]string), rev: make(map[string]string)}
	fs.load()
	return fs
}

func (f *FileStore) load() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.m = make(map[string]string)
	f.rev = make(map[string]string)
	file, err := os.Open(f.path)
	if err != nil {
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&f.m)
	// build reverse index
	for k, v := range f.m {
		f.rev[v] = k
	}
}

// persistSnapshot writes the provided snapshot to disk without acquiring
// any locks. Callers must ensure they have created a snapshot while holding
// the appropriate lock (or otherwise ensure snapshot safety).
func (f *FileStore) persistSnapshot(snapshot map[string]string) error {
	file, err := os.Create(f.path)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(snapshot)
}

func (f *FileStore) Save(code, url string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	// If code already exists
	if existing, ok := f.m[code]; ok {
		if existing == url {
			// already stored; idempotent
			return nil
		}
		return fmt.Errorf("code already exists")
	}
	// If URL already present mapped to a different code, return existing code error
	if existingCode, ok := f.rev[url]; ok {
		if existingCode == code {
			return nil
		}
		return fmt.Errorf("url already exists with different code")
	}
	f.m[code] = url
	f.rev[url] = code
	snapshot := make(map[string]string, len(f.m))
	for k, v := range f.m {
		snapshot[k] = v
	}
	return f.persistSnapshot(snapshot)
}

func (f *FileStore) Lookup(code string) (string, bool) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	u, ok := f.m[code]
	return u, ok
}

func (f *FileStore) FindByURL(url string) (string, bool) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	c, ok := f.rev[url]
	return c, ok
}

func (f *FileStore) All() ([]model.URLMapping, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	out := make([]model.URLMapping, 0, len(f.m))
	for k, v := range f.m {
		out = append(out, model.URLMapping{Code: k, URL: v})
	}
	return out, nil
}

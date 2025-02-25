package storage

import "sync"

var (
	docStorage = make(map[string]string) // In-memory storage
	mu         sync.Mutex
)

// SaveFile stores a document in memory
func SaveFile(filename, content string) {
	mu.Lock()
	docStorage[filename] = content
	mu.Unlock()
}

// GetFile retrieves a document from memory
func GetFile(filename string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	content, exists := docStorage[filename]
	return content, exists
}

package redis

import (
	"fmt"
	"sync"
)

type Redis struct {
	mu     sync.Mutex
	values map[string]string
}

func (r *Redis) Set(key string, value string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.values[key]; !ok {
		r.values[key] = value
	} else {
		return fmt.Errorf("Value already exists.")
	}

	return nil
}

func (r *Redis) Get(key string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if value, ok := r.values[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("No value exists there.")
}

func (r *Redis) Delete(key string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.values[key]; !ok {
		return fmt.Errorf("No value exists at key: %s", key)
	}

	delete(r.values, key)

	return nil
}

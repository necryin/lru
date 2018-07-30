package lru

import "testing"

func TestLRU(t *testing.T) {
	cache := New(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	if actual := cache.Get(1); actual != 1 {
		t.Error("Invalid value, expected: 1, got: ", actual)
	}

	cache.Put(3, 3) // evicts key 2
	if actual := cache.Get(2); actual != KeyNotFound {
		t.Error("Invalid value, expected: ", KeyNotFound, ", got: ", actual)
	}

	cache.Put(4, 4) // evicts key 1
	if actual := cache.Get(1); actual != KeyNotFound {
		t.Error("Invalid value, expected: ", KeyNotFound, ", got: ", actual)
	}
	if actual := cache.Get(3); actual != 3 {
		t.Error("Invalid value, expected: 3, got: ", actual)
	}
	if actual := cache.Get(4); actual != 4 {
		t.Error("Invalid value, expected: 4, got: ", actual)
	}
}

func TestLRU_ZeroCap(t *testing.T) {
	cache := New(0)

	cache.Put(1, 1) // evicts instantly after put
	if actual := cache.Get(1); actual != KeyNotFound {
		t.Error("Invalid value, expected: ", KeyNotFound, ", got: ", actual)
	}
}

package main

import "fmt"

type HashMap struct {
	buckets [][][]string
	size    int
}

func (h *HashMap) Set(key string, value string) string {
	index := h.XORHash([]byte(key))
	insertedKeyVal := []string{key, value}
	h.buckets[index] = append(h.buckets[index], insertedKeyVal)
	return key
}

func (h *HashMap) Get(key string) (string, bool) {
	index := h.XORHash([]byte(key))
	if len(h.buckets[index]) == 0 {
		return "", false
	}

	for _, kv := range h.buckets[index] {
		if kv[0] == key {
			return kv[1], true
		}
	}

	return "", false
}

func (h *HashMap) Del(key string) bool {
	index := h.XORHash([]byte(key))

	if len(h.buckets[index]) == 0 {
		return false
	}

	for i, kv := range h.buckets[index] {
		if kv[0] == key {
			h.buckets[index] = append(h.buckets[index][:i], h.buckets[index][i+1:]...)
			return true
		}
	}

	return false
}

func (h *HashMap) XORHash(src []byte) int {
	index := 0
	for _, c := range src {
		index = index ^ int(c)
	}
	return index / h.size
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([][][]string, size),
		size:    size,
	}
}

func main() {
	fmt.Println("vim-go")
	h := NewHashMap(11)
	h.Set("hello", "world")
	h.Set("barab", "123,")
	h.Set("qweqwe", "321")
	h.Set("asdasda", "999")
	h.Set("pooop", "12312")
	val, ok := h.Get("hello")
	fmt.Println(val, ok)
	fmt.Println(h.Get("asdasda"))

	h.Del("barab")
	fmt.Println(h.buckets)
	fmt.Println(h.Get("hello"))
}

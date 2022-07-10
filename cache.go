package cache

import (
	"time"
)

type CachedValue struct {
	data     string
	deadline time.Time
}

type Cache struct {
	hashMap map[string]CachedValue
}

func NewCache() Cache {
	return Cache{make(map[string]CachedValue)}
}

func (receiver Cache) Get(key string) (string, bool) {
	value, ok := receiver.hashMap[key]
	if !ok {
		return "", ok
	}

	if !value.deadline.IsZero() && !value.deadline.After(time.Now()) {
		delete(receiver.hashMap, key)
		return "", false
	}

	return value.data, true
}

func (receiver *Cache) Put(key, value string) {
	receiver.hashMap[key] = CachedValue{data: value}
}

func (receiver Cache) Keys() []string {
	existingKeys := []string{}
	for key, value := range receiver.hashMap {
		if value.deadline.IsZero() || value.deadline.After(time.Now()) {
			existingKeys = append(existingKeys, key)
		} else {
			delete(receiver.hashMap, key)
		}
	}
	return existingKeys
}

func (receiver *Cache) PutTill(key, value string, deadline time.Time) {
	receiver.hashMap[key] = CachedValue{data: value, deadline: deadline}
}

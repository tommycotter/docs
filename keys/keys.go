package keys

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"

	"gitlab.benzinga.io/benzinga/bzgo/pkg/models/apikeys"
)

type Cache struct {
	cache *cache.Cache
}

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrInvalidType = errors.New("invalid type")
)

func NewCache() *Cache {
	return &Cache{cache.New(cache.NoExpiration, 15*time.Minute)}
}

func (c *Cache) SetKey(key *apikeys.Key) {
	c.cache.Set(key.Key, key, cache.NoExpiration)
}

func (c *Cache) GetKeyUser(apiKey string) (*apikeys.Key, error) {
	v, found := c.cache.Get(apiKey)
	if !found {
		return nil, ErrKeyNotFound
	}

	keyDetail, ok := v.(*apikeys.Key)
	if !ok {
		return nil, ErrInvalidType
	}

	return keyDetail, nil
}

// RefreshKeys compares the updated keys to the existing keys and removes those not found in updated keys.
func (c *Cache) RefreshKeys(keys []apikeys.Key) (deleted int) {
	// This avoids using another sync.Mutex
	updatedKeys := make(map[string]struct{})

	for i := 0; i < len(keys); i++ {
		updatedKeys[keys[i].Key] = struct{}{}
	}

	for existingKey := range c.cache.Items() {
		if _, found := updatedKeys[existingKey]; !found {
			c.cache.Delete(existingKey)
			deleted++
		}
	}

	for i := 0; i < len(keys); i++ {
		c.SetKey(&keys[i])
	}

	return deleted
}

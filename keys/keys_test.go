package keys

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.benzinga.io/benzinga/bzgo/pkg/models/apikeys"
)

func TestKeys(t *testing.T) {
	k := NewCache()

	const (
		testUser = "testUser"
		testKey  = "testKey"
	)

	apiKey := apikeys.Key{
		ClientName: testUser,
		Key:        testKey,
	}

	makeTestKeys := func(num int) (res []apikeys.Key) {
		for i := 0; i < num; i++ {
			res = append(res, apikeys.Key{
				ClientName: fake.Company(),
				Key:        fake.CharactersN(20),
			})
		}

		return
	}

	k.SetKey(&apiKey)

	t.Run("Test Key Cache - Expect OK", func(t *testing.T) {
		user, err := k.GetKeyUser(testKey)
		require.NoError(t, err)
		assert.Equal(t, testUser, user.ClientName)
		assert.Equal(t, testKey, user.Key)
	})

	t.Run("Test Key Cache - Expect Not Found", func(t *testing.T) {
		user, err := k.GetKeyUser("invalid")
		assert.EqualError(t, ErrKeyNotFound, err.Error(), "Should return key not found error when key is invalid")
		assert.Empty(t, user, "User returned should be empty when key not found")
	})

	t.Run("Test Key Cache - Expect Invalid Value", func(t *testing.T) {
		k.cache.Set("invalidType", 1, cache.DefaultExpiration)
		user, err := k.GetKeyUser("invalidType")
		assert.EqualError(t, ErrInvalidType, err.Error(), "Should return invalid type error when value returned from cache is model.APIKey")
		assert.Empty(t, user, "Result returned should be empty when value is not model.APIKey")
	})

	t.Run("Test Key Cache Refresh", func(t *testing.T) {
		// Reset
		k = NewCache()
		k.SetKey(&apiKey)

		testKeys := makeTestKeys(4)

		res := k.RefreshKeys(testKeys)
		assert.Equal(t, 1, res, "One key (apiKey) should have been deleted")

		for i := 0; i < len(testKeys); i++ {
			u, err := k.GetKeyUser(testKeys[i].Key)
			assert.NoError(t, err, "Refreshed Keys should be returned without issue")
			assert.Equal(t, testKeys[i].ClientName, u.ClientName, "Client should match API key")
		}

		u, err := k.GetKeyUser(apiKey.Key)
		assert.EqualError(t, ErrKeyNotFound, err.Error(), "Should return key not found error since key should be removed")
		assert.Nil(t, u, "User returned should be nil when key not found")
	})
}

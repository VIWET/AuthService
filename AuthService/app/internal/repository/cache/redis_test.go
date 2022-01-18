package cache_test

import (
	"testing"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/repository/cache"
	"github.com/stretchr/testify/assert"
)

func TestRedisCacheRepository_Set(t *testing.T) {
	c := cache.TestRedisCache(t, config)

	r := cache.NewRedisCacheRepository(c, config.Expires)

	rt := "sdkfjnsijdfnsjdfmjsinfjsndflm"

	rs := &domain.RefreshSession{
		ProfileID:   1,
		Role:        "user",
		UserAgent:   "UserAgent",
		Fingerprint: "hjabsd41561fihsdnfihsdnfih615df6s4df65s1df65s41df65s1df651sf",
	}

	err := r.Set(rt, rs)
	assert.NoError(t, err)
}

func TestRedisCacheRepository_Get(t *testing.T) {
	c := cache.TestRedisCache(t, config)

	r := cache.NewRedisCacheRepository(c, config.Expires)

	rt := "sdkfjnsijdfnsjdfmjsinfjsndflm"

	rs := &domain.RefreshSession{
		ProfileID:   1,
		Role:        "user",
		UserAgent:   "UserAgent",
		Fingerprint: "hjabsd41561fihsdnfihsdnfih615df6s4df65s1df65s41df65s1df651sf",
	}

	err := r.Set(rt, rs)
	if err != nil {
		t.Fatal(err)
	}

	rs_test, err := r.Get(rt)
	assert.NoError(t, err)
	assert.Equal(t, rs, rs_test)
}

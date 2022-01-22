package testcache_test

import (
	"testing"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/repository/testcache"
	"github.com/stretchr/testify/assert"
)

func TestCacheRepository_Set(t *testing.T) {
	r := testcache.NewTestCacheRepository()

	rt := "sdkfjnsijdmjsinfjsndflm"

	rs := &domain.RefreshSession{
		ProfileID:   1,
		Role:        "user",
		UserAgent:   "UserAgent",
		Fingerprint: "hjabsd41561fihsdnfihsdnfih615df6s4df65s1df65s41df65s1df651sf",
	}

	err := r.Set(rt, rs)
	assert.NoError(t, err)
}

func TestCacheRepository_Get(t *testing.T) {
	r := testcache.NewTestCacheRepository()

	rt := "sdkfjnsijmjsinfjsndflm"

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

func TestCacheRepository_Delete(t *testing.T) {
	r := testcache.NewTestCacheRepository()

	rt := "sdkfjnsijmjsinfjsndflm"

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
	err = r.Delete(rt)
	assert.NoError(t, err)
	rs_test, err = r.Get(rt)
	assert.Nil(t, rs_test)
	assert.Error(t, err)
}

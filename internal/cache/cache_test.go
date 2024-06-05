package cache

import (
	"github.com/evanstukalov/wildberries_internship_l0/internal/models"
	"testing"
)

func TestAllEntriesStoredInCache(t *testing.T) {

	cache := &inMemoryCache{
		data: make(map[string]item),
	}

	testData := map[string]models.Order{
		"order1": {OrderUID: "UID1"},
		"order2": {OrderUID: "UID2"},
	}

	err := cache.FillUp(testData)
	if err != nil {
		t.Errorf("Error filling up cache: %v", err)
	}

	if len(cache.data) != len(testData) {
		t.Errorf("Expected cache size %d, got %d", len(testData), len(cache.data))
	}

}

func TestEmptyMapResultsInEmptyCache(t *testing.T) {

	cache := &inMemoryCache{
		data: make(map[string]item),
	}

	emptyData := make(map[string]models.Order)

	err := cache.FillUp(emptyData)
	if err != nil {
		t.Errorf("FillUp returned an error: %v", err)
	}

	if len(cache.data) != 0 {
		t.Errorf("Expected empty cache, but got size %d", len(cache.data))
	}
}

func TestSetNewKeyValuePair(t *testing.T) {
	cache := &inMemoryCache{
		data: make(map[string]item),
	}

	err := cache.Set("key1", "value1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	if val, found := cache.data["key1"]; !found || val.value != "value1" {
		t.Fatalf("expected value1, got %v", val.value)
	}
}

func TestGetNonExistentKey(t *testing.T) {
	cache := &inMemoryCache{
		data: make(map[string]item),
	}

	_, err := cache.Get("nonexistent")
	if err == nil || err.Error() != "item not found" {
		t.Fatalf("expected 'item not found' error, got %v", err)
	}
}

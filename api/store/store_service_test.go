package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://urlame.io"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	SaveUrlMapping(shortURL, initialLink)

	// Retrieve initial URL
	retrievedUrl := *RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)

	failedUrl := RetrieveInitialUrl("derp2do")
	assert.True(t, failedUrl == nil)
}
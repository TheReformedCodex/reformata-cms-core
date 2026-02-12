package utilities

import (
	"fmt"
	"reformata-cms-core/configs"
	"regexp"
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

type VideoSearchId struct {
	Kind    string `json:"kind"`
	VideoId string `json:"videoId"`
}

type VideoSearchSnippet struct {
	PublishTime time.Time `json:"publishTime"`
	Title       string    `json:"title"`
}

type VideoSearchResult struct {
	Kind string             `json:"kind"`
	Id   VideoSearchId      `json:"id"`
	Info VideoSearchSnippet `json:"snippet"`
}

type VideoSearchSnippetResponse struct {
	Items []VideoSearchResult `json:"items"`
}

func FetchRecentVideo() VideoSearchResult {
	client := req.C()

	resultCache := NewTTL[string, VideoSearchResult]()

	cachedResults, cachedResultsState := resultCache.Get("mostRecentVideo")

	if cachedResultsState {
		println("Returning Cached Results")
		return cachedResults
	}

	api_url := configs.Config.ConfigFile.YouTubeApiUrl
	channel_id := configs.Config.ConfigFile.YouTubeChannelId
	api_key := configs.Config.Secrets.YouTubeAPIKey

	fmt.Printf("\n\nAPI Key: %s\n\n", api_key)

	var response VideoSearchSnippetResponse

	resp, err := client.NewRequest().
		SetQueryParam("channelId", channel_id).
		SetQueryParam("key", api_key).
		SetQueryParam("order", "date").
		SetQueryParam("part", "snippet").
		SetSuccessResult(&response).
		Get(api_url)

	if err != nil {
		println("Unable to query the YouTube API", err)
	}

	if resp.Err != nil {
		println("Unable to query the YouTube API", resp.Err)
	}

	if len(response.Items) == 0 {
		fmt.Println("Response from API empty")
		return VideoSearchResult{}
	}

	for index, value := range response.Items {
		match, _ := regexp.MatchString("Sunday Service", value.Info.Title)

		if match {
			resultCache.Set("mostRecentVideo", response.Items[index], time.Minute*15)
			return response.Items[index]
		}
	}
	return response.Items[0]
}

// item represents a cache item with a value and an expiration time.
type item[V any] struct {
	value  V
	expiry time.Time
}

// isExpired checks if the cache item has expired.
func (i item[V]) isExpired() bool {
	return time.Now().After(i.expiry)
}

// TTLCache is a generic cache implementation with support for time-to-live
// (TTL) expiration.
type TTLCache[K comparable, V any] struct {
	items map[K]item[V] // The map storing cache items.
	mu    sync.Mutex    // Mutex for controlling concurrent access to the cache.
}

// NewTTL creates a new TTLCache instance and starts a goroutine to periodically
// remove expired items every 5 seconds.
func NewTTL[K comparable, V any]() *TTLCache[K, V] {
	c := &TTLCache[K, V]{
		items: make(map[K]item[V]),
	}

	go func() {
		for range time.Tick(5 * time.Second) {
			c.mu.Lock()

			// Iterate over the cache items and delete expired ones.
			for key, item := range c.items {
				if item.isExpired() {
					delete(c.items, key)
				}
			}

			c.mu.Unlock()
		}
	}()

	return c
}

// Set adds a new item to the cache with the specified key, value, and
// time-to-live (TTL).
func (c *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item[V]{
		value:  value,
		expiry: time.Now().Add(ttl),
	}
}

// Get retrieves the value associated with the given key from the cache.
func (c *TTLCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		// If the key is not found, return the zero value for V and false.
		return item.value, false
	}

	if item.isExpired() {
		// If the item has expired, remove it from the cache and return the
		// value and false.
		delete(c.items, key)
		return item.value, false
	}

	// Otherwise return the value and true.
	return item.value, true
}

// Remove removes the item with the specified key from the cache.
func (c *TTLCache[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Delete the item with the given key from the cache.
	delete(c.items, key)
}

// Pop removes and returns the item with the specified key from the cache.
func (c *TTLCache[K, V]) Pop(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		// If the key is not found, return the zero value for V and false.
		return item.value, false
	}

	// If the key is found, delete the item from the cache.
	delete(c.items, key)

	if item.isExpired() {
		// If the item has expired, return the value and false.
		return item.value, false
	}

	// Otherwise return the value and true.
	return item.value, true
}

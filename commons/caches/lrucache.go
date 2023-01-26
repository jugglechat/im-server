package caches

import (
	"sync"
	"time"

	"github.com/hashicorp/golang-lru/simplelru"
)

type lruCacheItem struct {
	value        interface{}
	lastReadTime int64
	addedTime    int64
}

type LruCache struct {
	lru                simplelru.LRUCache
	lock               sync.RWMutex
	readTimeoutChecker *time.Ticker
	MaxLifeCycle       time.Duration
	valueCreator       func(key interface{}) interface{}
}

func NewLruCacheWithAddReadTimeout(size int, onEvict simplelru.EvictCallback, timeoutAfterRead time.Duration, timeoutAfterCreate time.Duration) *LruCache {
	cache := NewLruCache(size, onEvict)
	cache.AddTimeoutAfterRead(timeoutAfterRead)
	cache.AddTimeoutAfterCreate(timeoutAfterCreate)
	return cache
}

func NewLruCache(size int, onEvict simplelru.EvictCallback) *LruCache {
	myLru, _ := simplelru.NewLRU(size, onEvict)
	cache := &LruCache{
		lru: myLru,
	}
	return cache
}

func (c *LruCache) AddValueCreator(creator func(interface{}) interface{}) *LruCache {
	c.valueCreator = creator
	return c
}

func (c *LruCache) AddTimeoutAfterCreate(timeout time.Duration) *LruCache {
	c.MaxLifeCycle = timeout
	return c
}

func (c *LruCache) AddTimeoutAfterRead(timeout time.Duration) *LruCache {
	if c.readTimeoutChecker != nil {
		c.readTimeoutChecker.Stop()
	}
	c.readTimeoutChecker = time.NewTicker(time.Second)
	go func() {
		for task := range c.readTimeoutChecker.C {
			current := time.Now().UnixMilli()
			if current-task.UnixMilli() > 500 {
				continue
			}
			timeLine := current - int64(timeout)/(1000*1000)
			c.cleanOdlestByReadTime(timeLine)
		}
	}()
	return c
}

func (c *LruCache) cleanOdlestByReadTime(timeLine int64) {
	for {
		itemKey, itemValue, ok := c.lru.GetOldest()
		if ok {
			lastReadTime := itemValue.(lruCacheItem).lastReadTime
			if lastReadTime < timeLine {
				c.Remove(itemKey)
			} else {
				break
			}
		} else {
			break
		}
	}
}

func (c *LruCache) Add(key, value interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.innerAdd(key, value)
}

func (c *LruCache) AddIfAbsent(key, value interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.innerContains(key) {
		old, ok := c.innerGet(key)
		if ok {
			return old, false
		}
	}
	c.innerAdd(key, value)
	return value, true
}

func (c *LruCache) innerAdd(key, value interface{}) bool {
	nowTime := time.Now().UnixMilli()
	return c.lru.Add(key, lruCacheItem{
		value:        value,
		lastReadTime: nowTime,
		addedTime:    nowTime,
	})
}

func (c *LruCache) Get(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.innerGet(key)
}
func (c *LruCache) innerGet(key interface{}) (interface{}, bool) {
	item, ok := c.lru.Get(key)
	if ok {
		cacheItem := item.(lruCacheItem)
		if c.MaxLifeCycle > 0 {
			timeLine := time.Now().UnixMilli() - int64(c.MaxLifeCycle)/1000/1000
			if cacheItem.addedTime < timeLine { //remove
				c.lru.Remove(key)
				return nil, false
			}
		}
		cacheItem.lastReadTime = time.Now().UnixMilli()
		return cacheItem.value, ok
	} else {
		return nil, ok
	}
}
func (c *LruCache) GetByDefault(key interface{}, defaultValue interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	val, ok := c.innerGet(key)
	if ok {
		return val, ok
	} else {
		return defaultValue, ok
	}
}
func (c *LruCache) GetByCreator(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	val, ok := c.innerGet(key)
	if ok {
		return val, ok
	} else {
		if c.valueCreator != nil {
			newVal := c.valueCreator(key)
			if newVal != nil {
				c.innerAdd(key, newVal)
				return newVal, true
			}
		}
	}
	return nil, ok
}
func (c *LruCache) Contains(key interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Contains(key)
}

func (c *LruCache) innerContains(key interface{}) bool {
	return c.lru.Contains(key)
}

func (c *LruCache) Peek(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	item, ok := c.lru.Peek(key)
	if ok {
		return item.(lruCacheItem).value, ok
	} else {
		return nil, ok
	}
}

func (c *LruCache) Purge() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.lru.Purge()
}
func (c *LruCache) Len() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Len()
}
func (c *LruCache) ReSize(size int) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Resize(size)
}

func (c *LruCache) Remove(key interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Remove(key)
}

func (c *LruCache) Keys() []interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Keys()
}

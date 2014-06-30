package httprouter

import "sync"

// RequestContext represents a request context which contains
// request scope key/value data.
type RequestContext struct {
	kv    map[string]interface{}
	mutex *sync.RWMutex
}

// GetValue returns the request context's value.
func (reqCtx *RequestContext) GetValue(key string) (interface{}, bool) {
	reqCtx.mutex.RLock()
	defer reqCtx.mutex.RUnlock()
	value, ok := reqCtx.kv[key]
	return value, ok
}

// SetValue sets the key/value to the request context.
func (reqCtx *RequestContext) SetValue(key string, value interface{}) {
	reqCtx.mutex.Lock()
	defer reqCtx.mutex.Unlock()
	reqCtx.kv[key] = value
}

// NewRequestContext creates and returns a request context.
func NewRequestContext() *RequestContext {
	return &RequestContext{
		kv:    make(map[string]interface{}),
		mutex: new(sync.RWMutex),
	}
}

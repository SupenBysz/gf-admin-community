package kmap

import (
	"github.com/SupenBysz/gf-admin-community/utility/deepcopy"
	"github.com/SupenBysz/gf-admin-community/utility/empty"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-community/utility/json"
	"github.com/SupenBysz/gf-admin-community/utility/rwmutex"
	"github.com/gogf/gf/v2/util/gconv"
)

type HashMap[TK string | int | int64, TD any] struct {
	mu   rwmutex.RWMutex
	data map[TK]TD
}

func New[TK string | int | int64, TD any](safe ...bool) *HashMap[TK, TD] {
	return &HashMap[TK, TD]{
		mu:   rwmutex.Create(safe...),
		data: make(map[TK]TD),
	}
}

func NewFrom[TK string | int | int64, TD any](mapData map[TK]TD, safe ...bool) *HashMap[TK, TD] {
	return &HashMap[TK, TD]{
		mu:   rwmutex.Create(safe...),
		data: mapData,
	}
}

// Iterator iterates the hash map readonly with custom callback function `f`.
// If `f` returns true, then it continues iterating; or false to stop.
func (m *HashMap[TK, TD]) Iterator(f func(k TK, v TD) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// Clone returns a new hash map with copy of current map data.
func (m *HashMap[TK, TD]) Clone(safe ...bool) *HashMap[TK, TD] {
	return NewFrom[TK, TD](m.MapCopy(), safe...)
}

// Map returns the underlying data map.
// Note that, if it's in concurrent-safe usage, it returns a copy of underlying data,
// or else a pointer to the underlying data.
func (m *HashMap[TK, TD]) Map() map[TK]TD {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if !m.mu.IsSafe() {
		return m.data
	}
	data := make(map[TK]TD, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// MapCopy returns a shallow copy of the underlying data of the hash map.
func (m *HashMap[TK, TD]) MapCopy() map[TK]TD {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[TK]TD, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// MapStrAny returns a copy of the underlying data of the map as map[string]interface{}.
func (m *HashMap[TK, TD]) MapStrAny() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		data[gconv.String(k)] = v
	}
	return data
}

// FilterEmpty deletes all key-value pair of which the value is empty.
// Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.
func (m *HashMap[TK, TD]) FilterEmpty() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsEmpty(v) {
			delete(m.data, k)
		}
	}
}

// FilterNil deletes all key-value pair of which the value is nil.
func (m *HashMap[TK, TD]) FilterNil() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsNil(v) {
			delete(m.data, k)
		}
	}
}

// Set sets key-value to the hash map.
func (m *HashMap[TK, TD]) Set(key TK, value TD) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[TK]TD)
	}
	m.data[key] = value
	m.mu.Unlock()
}

// Sets batch sets key-values to the hash map.
func (m *HashMap[TK, TD]) Sets(data map[TK]TD) {
	m.mu.Lock()
	if m.data == nil {
		m.data = data
	} else {
		for k, v := range data {
			m.data[k] = v
		}
	}
	m.mu.Unlock()
}

// Search searches the map with given `key`.
// Second return parameter `found` is true if key was found, otherwise false.
func (m *HashMap[TK, TD]) Search(key TK) (value TD, found bool) {
	m.mu.RLock()
	if m.data != nil {
		value, found = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Get returns the value by given `key`.
func (m *HashMap[TK, TD]) Get(key TK) (value TD) {
	m.mu.RLock()
	if m.data != nil {
		value = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Pop retrieves and deletes an item from the map.
func (m *HashMap[TK, TD]) Pop() (key TK, value TD) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for key, value = range m.data {
		delete(m.data, key)
		return
	}
	return
}

// Pops retrieves and deletes `size` items from the map.
// It returns all items if size == -1.
func (m *HashMap[TK, TD]) Pops(size int) map[TK]TD {
	m.mu.Lock()
	defer m.mu.Unlock()
	if size > len(m.data) || size == -1 {
		size = len(m.data)
	}
	if size == 0 {
		return nil
	}
	var (
		index  = 0
		newMap = make(map[TK]TD, size)
	)
	for k, v := range m.data {
		delete(m.data, k)
		newMap[k] = v
		index++
		if index == size {
			break
		}
	}
	return newMap
}

// doSetWithLockCheck checks whether value of the key exists with mutex.Lock,
// if not exists, set value to the map with given `key`,
// or else just return the existing value.
//
// When setting value, if `value` is type of `func() interface {}`,
// it will be executed with mutex.Lock of the hash map,
// and its return value will be set to the map with `key`.
//
// It returns value with given `key`.
func (m *HashMap[TK, TD]) doSetWithLockCheck(key TK, value interface{}) TD {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[TK]TD)
	}
	if v, ok := m.data[key]; ok {
		return v
	}
	if f, ok := value.(func() TD); ok {
		value = f()
	}
	if value != nil {
		m.data[key] = funs.AsType[TD](value)
	}
	return funs.AsType[TD](value)
}

// GetOrSet returns the value by key,
// or sets value with given `value` if it does not exist and then returns this value.
func (m *HashMap[TK, TD]) GetOrSet(key TK, value TD) TD {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFunc returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
func (m *HashMap[TK, TD]) GetOrSetFunc(key TK, f func() TD) TD {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
//
// GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f`
// with mutex.Lock of the hash map.
func (m *HashMap[TK, TD]) GetOrSetFuncLock(key TK, f func() TD) TD {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// SetIfNotExist sets `value` to the map if the `key` does not exist, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
func (m *HashMap[TK, TD]) SetIfNotExist(key TK, value TD) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc sets value with return value of callback function `f`, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
func (m *HashMap[TK, TD]) SetIfNotExistFunc(key TK, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f())
		return true
	}
	return false
}

// SetIfNotExistFuncLock sets value with return value of callback function `f`, and then returns true.
// It returns false if `key` exists, and `value` would be ignored.
//
// SetIfNotExistFuncLock differs with SetIfNotExistFunc function is that
// it executes function `f` with mutex.Lock of the hash map.
func (m *HashMap[TK, TD]) SetIfNotExistFuncLock(key TK, f func() interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Remove deletes value from map by given `key`, and return this deleted value.
func (m *HashMap[TK, TD]) Remove(key TK) (value TD) {
	m.mu.Lock()
	if m.data != nil {
		var ok bool
		if value, ok = m.data[key]; ok {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
	return
}

// Removes batch deletes values of the map by keys.
func (m *HashMap[TK, TD]) Removes(keys []TK) {
	m.mu.Lock()
	if m.data != nil {
		for _, key := range keys {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
}

// Keys returns all keys of the map as a slice.
func (m *HashMap[TK, TD]) Keys() []TK {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var (
		keys  = make([]TK, len(m.data))
		index = 0
	)
	for key := range m.data {
		keys[index] = key
		index++
	}
	return keys
}

// Values returns all values of the map as a slice.
func (m *HashMap[TK, TD]) Values() []TD {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var (
		values = make([]TD, len(m.data))
		index  = 0
	)
	for _, value := range m.data {
		values[index] = value
		index++
	}
	return values
}

// Contains checks whether a key exists.
// It returns true if the `key` exists, or else false.
func (m *HashMap[TK, TD]) Contains(key TK) bool {
	var ok bool
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[key]
	}
	m.mu.RUnlock()
	return ok
}

// Size returns the size of the map.
func (m *HashMap[TK, TD]) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// IsEmpty checks whether the map is empty.
// It returns true if map is empty, or else false.
func (m *HashMap[TK, TD]) IsEmpty() bool {
	return m.Size() == 0
}

// Clear deletes all data of the map, it will remake a new underlying data map.
func (m *HashMap[TK, TD]) Clear() {
	m.mu.Lock()
	m.data = make(map[TK]TD)
	m.mu.Unlock()
}

// Replace the data of the map with given `data`.
func (m *HashMap[TK, TD]) Replace(data map[TK]TD) {
	m.mu.Lock()
	m.data = data
	m.mu.Unlock()
}

// LockFunc locks writing with given callback function `f` within RWMutex.Lock.
func (m *HashMap[TK, TD]) LockFunc(f func(m map[TK]TD)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// RLockFunc locks reading with given callback function `f` within RWMutex.RLock.
func (m *HashMap[TK, TD]) RLockFunc(f func(m map[TK]TD)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.data)
}

// Merge merges two hash maps.
// The `other` map will be merged into the map `m`.
func (m *HashMap[TK, TD]) Merge(other *HashMap[TK, TD]) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = other.MapCopy()
		return
	}
	if other != m {
		other.mu.RLock()
		defer other.mu.RUnlock()
	}
	for k, v := range other.data {
		m.data[k] = v
	}
}

// String returns the map as a string.
func (m *HashMap[TK, TD]) String() string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSON()
	return string(b)
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (m HashMap[TK, TD]) MarshalJSON() ([]byte, error) {
	return json.Marshal(gconv.Map(m.Map()))
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (m *HashMap[TK, TD]) UnmarshalJSON(b []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[TK]TD)
	}
	var data map[TK]TD
	if err := json.UnmarshalUseNumber(b, &data); err != nil {
		return err
	}
	for k, v := range data {
		m.data[k] = v
	}
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for map.
func (m *HashMap[TK, TD]) UnmarshalValue(value *HashMap[TK, TD]) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[TK]TD)
	}
	for k, v := range value.data {
		m.data[k] = v
	}
	return
}

// DeepCopy implements interface for deep copy of current type.
func (m *HashMap[TK, TD]) DeepCopy() *HashMap[TK, TD] {
	if m == nil {
		return nil
	}

	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[TK]TD, len(m.data))
	for k, v := range m.data {
		data[k] = funs.AsType[TD](deepcopy.Copy(v))
	}
	return NewFrom(data, m.mu.IsSafe())
}

package threadpool

import (
	"github.com/golang/sync/syncmap"
)

// Set type
// This implementation is faster than slices or arrays because it internally uses map to store the data
// All the values stored as the keys in the map and value is dummy boolean, just a place holder.
// It stores the Unique elements only
type Set struct {
	_map *syncmap.Map
}

// NewSet creates and returns new set
func NewSet() *Set {
	set := new(Set)
	set._map = new(syncmap.Map)
	return set
}

// Add adds the value to the Set
func (s *Set) Add(value interface{}) {
	s._map.Store(value, true)
}

// Remove the value from the Set
func (s *Set) Remove(value interface{}) {
	s._map.Delete(value)
}

//Contains checks if the value exists in the Set and returns true if exists
func (s *Set) Contains(value interface{}) bool {
	_, ok := s._map.Load(value)
	return ok
}

// GetAll returns all the values as Array
func (s *Set) GetAll() []interface{} {
	values := make([]interface{}, 0)
	s._map.Range(func(key interface{}, value interface{}) bool {
		values = append(values, key)
		return true
	})
	return values
}

// GetAllAsString returns the values as Array of string
func (s *Set) GetAllAsString() []string {
	values := make([]string, 0)
	s._map.Range(func(key interface{}, value interface{}) bool {
		values = append(values, key.(string))
		return true
	})
	return values
}

// GetAllWithCap returns the data in set with max data return limit
func (s *Set) GetAllWithCap(cap int) []interface{} {
	values := make([]interface{}, 0)

	s._map.Range(func(key interface{}, value interface{}) bool {
		values = append(values, key)
		if cap--; cap > 0 {
			return true
		} else {
			return false
		}
	})
	return values
}

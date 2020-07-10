package sync

import (
	"fmt"
	"strconv"
	"sync"
)

// Bool type with concurrent access support.
type Bool struct {
	rw    sync.RWMutex
	ch    chan<- bool
	value bool
}

// Get returns the internal value.
func (b *Bool) Get() bool {
	b.rw.RLock()
	defer b.rw.RUnlock()
	return b.value
}

// Set a value.
func (b *Bool) Set(value bool) {
	b.rw.Lock()
	defer b.rw.Unlock()
	b.value = value
	if b.ch != nil {
		b.ch <- value
	}
}

// String returns string representation of value.
func (b *Bool) String() string {
	b.rw.Lock()
	defer b.rw.Unlock()
	if b.value {
		return "true"
	}
	return "false"
}

// SetString parses and sets a value from string type.
func (b *Bool) SetString(val string) error {
	v, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	b.Set(v)
	return nil
}

// Listen sets a channel for value changes notifications.
func (b *Bool) Listen(ch chan<- bool) {
	b.rw.Lock()
	defer b.rw.Unlock()
	b.ch = ch
}

// Int64 type with concurrent access support.
type Int64 struct {
	rw    sync.RWMutex
	ch    chan<- int64
	value int64
}

// Get returns the internal value.
func (i *Int64) Get() int64 {
	i.rw.RLock()
	defer i.rw.RUnlock()
	return i.value
}

// Set a value.
func (i *Int64) Set(value int64) {
	i.rw.Lock()
	defer i.rw.Unlock()
	i.value = value
	if i.ch != nil {
		i.ch <- value
	}
}

// String returns string representation of value.
func (i *Int64) String() string {
	i.rw.RLock()
	defer i.rw.RUnlock()
	return fmt.Sprintf("%d", i.value)
}

// SetString parses and sets a value from string type.
func (i *Int64) SetString(val string) error {
	v, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}
	i.Set(v)
	return nil
}

// Listen sets a channel for value changes notifications.
func (i *Int64) Listen(ch chan<- int64) {
	i.rw.Lock()
	defer i.rw.Unlock()
	i.ch = ch
}

// Float64 type with concurrent access support.
type Float64 struct {
	rw    sync.RWMutex
	ch    chan<- float64
	value float64
}

// Get returns the internal value.
func (f *Float64) Get() float64 {
	f.rw.RLock()
	defer f.rw.RUnlock()
	return f.value
}

// Set a value.
func (f *Float64) Set(value float64) {
	f.rw.Lock()
	defer f.rw.Unlock()
	f.value = value
	if f.ch != nil {
		f.ch <- value
	}
}

// String returns string representation of value.
func (f *Float64) String() string {
	f.rw.RLock()
	defer f.rw.RUnlock()
	return fmt.Sprintf("%f", f.value)
}

// SetString parses and sets a value from string type.
func (f *Float64) SetString(val string) error {
	v, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	f.Set(v)
	return nil
}

// Listen sets a channel for value changes notifications.
func (f *Float64) Listen(ch chan<- float64) {
	f.rw.Lock()
	defer f.rw.Unlock()
	f.ch = ch
}

// String type with concurrent access support.
type String struct {
	rw    sync.RWMutex
	ch    chan<- string
	value string
}

// Get returns the internal value.
func (s *String) Get() string {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.value
}

// Set a value.
func (s *String) Set(value string) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.value = value
	if s.ch != nil {
		s.ch <- value
	}
}

// String returns string representation of value.
func (s *String) String() string {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.value
}

// SetString parses and sets a value from string type.
func (s *String) SetString(val string) error {
	s.Set(val)
	return nil
}

// Listen sets a channel for value changes notifications.
func (s *String) Listen(ch chan<- string) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.ch = ch
}

// Secret string type for secrets with concurrent access support.
type Secret struct {
	rw    sync.RWMutex
	ch    chan<- string
	value string
}

// Get returns the internal value.
func (s *Secret) Get() string {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.value
}

// Set a value.
func (s *Secret) Set(value string) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.value = value
	if s.ch != nil {
		s.ch <- value
	}
}

// String returns obfuscated string representation of value.
func (s *Secret) String() string {
	return "***"
}

// SetString parses and sets a value from string type.
func (s *Secret) SetString(val string) error {
	s.Set(val)
	return nil
}

// Listen sets a channel for value changes notifications.
func (s *Secret) Listen(ch chan<- string) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.ch = ch
}

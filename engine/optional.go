/*
Contains Optional[T] wrapper for struct fields whose natural zero-value differs
meaningfully from its uninitialized state.
*/
package engine

import "errors"

type Optional[T any] struct {
	Value T
	OK    bool
}

func (o Optional[T]) Get() (T, bool) {
	if !o.OK {
		var zero T
		return zero, false
	}
	return o.Value, o.OK
}

// Update() provides safety when creating a new value for an Optional field
// by first checking whether the field is currently initialized.
//
// It provides a new struct to facilitate immutable patterns. Mutation and
// initialization  can be executed via direct o.Value and o.OK field access.
func (o Optional[T]) Update(v T) (Optional[T], error) {
	if !o.OK {
		return Optional[T]{}, ErrUninitialized
	}

	result := Optional[T]{
		Value: v,
		OK:    true,
	}
	return result, nil
}

var ErrUninitialized = errors.New("Attempted to access an uninitialized value")

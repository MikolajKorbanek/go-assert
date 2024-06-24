package assert

import (
	"fmt"
	"reflect"
)

// Assert checks a condition and panics with the provided message if it is false.
func Assert(condition bool, message string) {
	if !condition {
		panic(fmt.Sprintf("Assertion failed: %s", message))
	}
}

// AssertEqual checks if two values are equal and panics if they are not.
func AssertEqual(expected, actual interface{}, message string) {
	if !reflect.DeepEqual(expected, actual) {
		panic(fmt.Sprintf("Assertion failed: %s - Expected: %v, Actual: %v", message, expected, actual))
	}
}

// AssertNotEqual checks if two values are not equal and panics if they are.
func AssertNotEqual(expected, actual interface{}, message string) {
	if reflect.DeepEqual(expected, actual) {
		panic(fmt.Sprintf("Assertion failed: %s - Values should not be equal: %v", message, actual))
	}
}

// AssertNil checks if a value is nil and panics if it is not.
func AssertNil(value interface{}, message string) {
	if !isNil(value) {
		panic(fmt.Sprintf("Assertion failed: %s - Expected nil, but got: %v", message, value))
	}
}

// AssertNotNil checks if a value is not nil and panics if it is.
func AssertNotNil(value interface{}, message string) {
	if isNil(value) {
		panic(fmt.Sprintf("Assertion failed: %s - Expected non-nil value", message))
	}
}

// AssertFalse checks if a value is false and panics if it is not.
func AssertFalse(value bool, message string) {
	if value {
		panic(fmt.Sprintf("Assertion failed: %s - Expected false, but got true", message))
	}
}

// Helper function to check if a value is nil
func isNil(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

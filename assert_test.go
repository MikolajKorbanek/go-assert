package assert_test

import (
	"testing"

	"github.com/MikolajKorbanek/go-assert"
)

func TestAssert_Success(t *testing.T) {
	assert.Assert(true, "This should not panic")
}

func TestAssert_Failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	assert.Assert(false, "This should panic")
}

func TestAssert_CustomMessage(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		} else {
			errMsg := r.(string)
			if errMsg != "Assertion failed: Custom message" {
				t.Errorf("Expected panic message 'Assertion failed: Custom message', got: %s", errMsg)
			}
		}
	}()
	assert.Assert(false, "Custom message")
}

func TestAssertEqual_Success(t *testing.T) {
	assert.AssertEqual(1, 1, "This should not panic")
}

func TestAssertEqual_Failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	assert.AssertEqual(1, 2, "This should panic")
}

func TestAssertEqual_CustomMessage(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		} else {
			errMsg := r.(string)
			if errMsg != "Assertion failed: Custom message - Expected: 1, Actual: 2" {
				t.Errorf("Expected panic message 'Assertion failed: Custom message - Expected: 1, Actual: 2', got: %s", errMsg)
			}
		}
	}()
	assert.AssertEqual(1, 2, "Custom message")
}

func TestAssertEqual_DifferentTypes(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	assert.AssertEqual(1, "1", "This should panic")
}

func TestAssertEqual_Nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	var x *int
	assert.AssertEqual(x, 1, "This should panic")
}

func TestAssertEqual_Structs(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	var x = testStruct{1, "test"}
	var y = testStruct{1, "test"}

	assert.AssertEqual(x, y, "This should not panic")
}

func TestAssertEqual_DifferentStructs(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	var x = testStruct{1, "test"}
	var y = testStruct{2, "test"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	assert.AssertEqual(x, y, "This should panic")
}

func TestAssertNotEqual_Success(t *testing.T) {
	assert.AssertNotEqual(1, 2, "This should not panic")
}

func TestAssertNotEqual_Failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	assert.AssertNotEqual(1, 1, "This should panic")
}

func TestAssertNotEqual_CustomMessage(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		} else {
			errMsg := r.(string)
			if errMsg != "Assertion failed: Custom message - Values should not be equal: 1" {
				t.Errorf("Expected panic message 'Assertion failed: Custom message - Values should not be equal: 1', got: %s", errMsg)
			}
		}
	}()
	assert.AssertNotEqual(1, 1, "Custom message")
}

func TestAssertNotEqual_DifferentTypes(t *testing.T) {
	assert.AssertNotEqual(1, "1", "This should not panic")
}

func TestAssertNotEqual_Nil(t *testing.T) {
	var x *int
	assert.AssertNotEqual(x, 1, "This should not panic")
}

func TestAssertNotEqual_Structs(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	var x = testStruct{1, "test"}
	var y = testStruct{2, "test"}

	assert.AssertNotEqual(x, y, "This should not panic")
}

func TestAssertNotEqual_SameStructs(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	var x = testStruct{1, "test"}
	var y = testStruct{1, "test"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	assert.AssertNotEqual(x, y, "This should panic")
}

func TestAssertNil_Success(t *testing.T) {
	var x *int
	assert.AssertNil(x, "This should not panic")
}

func TestAssertNil_Failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	var x = 1
	assert.AssertNil(x, "This should panic")
}

func TestAssertNil_CustomMessage(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		} else {
			errMsg := r.(string)
			if errMsg != "Assertion failed: Custom message - Expected nil, but got: 1" {
				t.Errorf("Expected panic message 'Assertion failed: Custom message - Expected nil, but got: 1', got: %s", errMsg)
			}
		}
	}()
	var x = 1
	assert.AssertNil(x, "Custom message")
}

func TestAssertNil_Struct(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	var x *testStruct
	assert.AssertNil(x, "This should not panic")
}

func TestAssertNil_StructWithValues(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	var x = testStruct{1, "test"}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	assert.AssertNil(x, "This should panic")
}

func TestAssertNotNil_Success(t *testing.T) {
	var x = 1
	assert.AssertNotNil(x, "This should not panic")
}

func TestAssertNotNil_Failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	var x *int
	assert.AssertNotNil(x, "This should panic")
}

func TestAssertNotNil_CustomMessage(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		} else {
			errMsg := r.(string)
			if errMsg != "Assertion failed: Custom message - Expected non-nil value" {
				t.Errorf("Expected panic message 'Assertion failed: Custom message - Expected non-nil value', got: %s", errMsg)
			}
		}
	}()
	var x *int
	assert.AssertNotNil(x, "Custom message")
}

func TestAssertNotNil_Struct(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	var x = testStruct{1, "test"}
	assert.AssertNotNil(x, "This should not panic")
}

func TestAssertNotNil_StructWithNil(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}

	var x *testStruct
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	assert.AssertNotNil(x, "This should panic")
}

func TestAssertFalse_Success(t *testing.T) {
	assert.AssertFalse(false, "This should not panic")
}

func TestAssertFalse_Failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		}
	}()
	assert.AssertFalse(true, "This should panic")
}

func TestAssertFalse_CustomMessage(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but code did not panic")
		} else {
			errMsg := r.(string)
			if errMsg != "Assertion failed: Custom message - Expected false, but got true" {
				t.Errorf("Expected panic message 'Assertion failed: Custom message - Expected false, but got true', got: %s", errMsg)
			}
		}
	}()
	assert.AssertFalse(true, "Custom message")

}

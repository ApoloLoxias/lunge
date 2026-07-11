package engine

import "testing"

func TestOptionalIntGet(t *testing.T) {
	testCases := []struct {
		name          string
		o             Optional[int]
		expectedValue int
		expectedOK    bool
	}{
		{
			name:          "unitialized",
			o:             Optional[int]{},
			expectedValue: 0,
			expectedOK:    false,
		},
		{
			name:          "unitialized with value",
			o:             Optional[int]{Value: 3, OK: false},
			expectedValue: 0,
			expectedOK:    false,
		},
		{
			name:          "initialized zero-value",
			o:             Optional[int]{OK: true},
			expectedValue: 0,
			expectedOK:    true,
		},
		{
			name:          "non-zero",
			o:             Optional[int]{Value: 3, OK: true},
			expectedValue: 3,
			expectedOK:    true,
		},
	}

	for _, testCase := range testCases {
		gotValue, gotOK := testCase.o.Get()
		if gotValue != testCase.expectedValue || gotOK != testCase.expectedOK {
			t.Errorf(
				"%v.Get() = %v, %v; want %v, %v",
				testCase.o,
				gotValue, gotOK,
				testCase.expectedValue, testCase.expectedOK,
			)
		}
	}
}

func TestOptionalIntUpdate(t *testing.T) {
	testCases := []struct {
		name          string
		o             Optional[int]
		UpdateArg     int
		expectedValue Optional[int]
		expectedErr   error
	}{
		{
			name:          "unitialized",
			o:             Optional[int]{},
			UpdateArg:     3,
			expectedValue: Optional[int]{},
			expectedErr:   ErrUninitialized,
		},
		{
			name:          "uninitialized with value",
			o:             Optional[int]{Value: 3},
			UpdateArg:     4,
			expectedValue: Optional[int]{},
			expectedErr:   ErrUninitialized,
		},
		{
			name:          "initialized zero-value",
			o:             Optional[int]{OK: true},
			UpdateArg:     4,
			expectedValue: Optional[int]{Value: 4, OK: true},
			expectedErr:   nil,
		},
		{
			name:          "non-zero",
			o:             Optional[int]{Value: 3, OK: true},
			UpdateArg:     4,
			expectedValue: Optional[int]{Value: 4, OK: true},
			expectedErr:   nil,
		},
	}

	for _, testCase := range testCases {
		a := testCase.o
		gotValue := testCase.o
		var gotErr error

		gotValue, gotErr = testCase.o.Update(testCase.UpdateArg)

		if &a == &gotValue {
			t.Errorf("updated value points to original value")
		}
		if a.Value != testCase.o.Value || a.OK != testCase.o.OK {
			t.Errorf(
				"Optional.Update() alters original, "+
					"got %v, want %v",
				a,
				testCase.o,
			)
		}
		if gotValue.Value != testCase.expectedValue.Value ||
			gotValue.OK != testCase.expectedValue.OK ||
			gotErr != testCase.expectedErr {
			t.Errorf(
				"%v.Update(%v) = %v, %v; want %v, %v",
				testCase.o,
				testCase.UpdateArg,
				gotValue, gotErr,
				testCase.expectedValue,
				testCase.expectedErr,
			)
		}
	}
}

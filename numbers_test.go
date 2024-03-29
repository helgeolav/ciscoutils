package ciscoutils

import (
	"testing"
)

func TestGetRangefromString(t *testing.T) {
	// first test - a valid string
	low, high, error := GetRangeFromString("4-10")
	if error != nil {
		t.Error("got nil")
		t.Fail()
	}
	if low != 4 {
		t.Error("did not get 4")
	}
	if high != 10 {
		t.Error("did not get 10")
	}
	// second test - an invalid range
	_, _, error = GetRangeFromString("10-4")
	if error == nil {
		t.Error("did not get nil on invalid range")
	}
	// 3rd test - send invalid string in
	_, _, error = GetRangeFromString("no-number")
	if error == nil {
		t.Error("passed on invalid string")
	}
}

func TestGetVLANString1(t *testing.T) {
	expectedResult := []int{1, 10, 11, 15}
	result, err := GetVLANString("1, 10-11,15")
	if err != nil {
		t.Error("got nil")
	}
	if IntArrayEquals(expectedResult, result) == false {
		t.Fail()
	}
}

func TestGetVLANString2(t *testing.T) {
	expectedResult := []int{1, 10, 15}
	result, err := GetVLANString("1,10,15")
	if err != nil {
		t.Error("got nil")
	}
	if IntArrayEquals(expectedResult, result) == false {
		t.Error("output is not correct")
	}
}

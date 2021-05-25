package ciscoutils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// GetRangeFromString parses a string with the syntax #-# and returns the two numbers.
func GetRangeFromString(input string) (low, high int, err error) {
	r := regexp.MustCompile("(\\d+)-(\\d+)")
	parsed := r.FindStringSubmatch(input)
	if parsed == nil {
		return 0, 0, errors.New("GetRangeFromString: input not in #-# syntax")
	}
	low, _ = strconv.Atoi(parsed[1])
	high, _ = strconv.Atoi(parsed[2])
	if low > high {
		return low, high, errors.New("number a higher than number b")
	}
	return low, high, nil
}

// IntArrayEquals support func to verify arrays
// source: https://stackoverflow.com/questions/18561219/comparing-arrays-in-go-language
func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// GetVLANString parses a line of comma-separated VLAN's in Cisco format
// 1,10,12-15,20
// TODO: add support for whitespace in string
func GetVLANString(input string) (vlans []int, err error) {
	var result []int
	s := strings.Split(input, ",")
	// loop through each value
	for _, each := range s {
		if strings.Contains(each, "-") {
			low, high, err := GetRangeFromString(each)
			if err != nil {
				return result, err
			}
			num := high - low
			for count := 0; count <= num; count++ {
				result = append(result, low+count)
			}
		} else {
			i, err := strconv.Atoi(each)
			if err != nil {
				return result, err
			}
			result = append(result, i)
		}
	}
	return result, nil
}

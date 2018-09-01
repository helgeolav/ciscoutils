package ciscoutils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// This func parses a string with the syntax #-# and returns a slice with the numbers in it
// TODO: implement check if Regexp does not parse
func GetRangefromString(input string) (low, high int, err error) {
	r := regexp.MustCompile("(\\d*)-(\\d*)")
	parsed := r.FindStringSubmatch(input)
	low, _ = strconv.Atoi(parsed[1])
	high, _ = strconv.Atoi(parsed[2])
	if low > high {
		return low, high, errors.New("number a higher than number b")
	}
	return low, high, nil
}

// support func to verify arrays
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

// parse a line of comma-separated VLAN's in Cisco format
// 1,10,12-15,20
// TODO: add support for whitespace in string
func GetVLANString(input string) (vlans []int, err error) {
	var result []int
	s := strings.Split(input, ",")
	// loop through each value
	for _, each := range s {
		if strings.Contains(each, "-") {
			low, high, err := GetRangefromString(each)
			if err != nil {
				return result, err
			}
			num := high - low
			for count := 0; count <= num; count++ {
				result = append(result, low+count)
			}
		} else {
			i, err := strconv.Atoi(each)
			if err == nil {
				result = append(result, i)
			}
		}
	}
	return result, nil
}

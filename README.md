# ciscoutils
Various code fragments to make the life easier as a Cisco stack developer in Go. The idea is here to collect small snippets of Cisco related code that is used to work with Cisco devices in many ways.

There are no external dependencies in this code. Only Go builtin libraries. Please note that some of the subpackages may have other external dependencies.

To use this library just download it like this:
```
go get github.com/helgeolav/ciscoutils
```

And then add it into your program
```
package main

import (
 "github.com/helgeolav/ciscoutils"
 "fmt"
)

func main() {
  vlans := ciscoutils.ReadVlanCsv("vlan_test.csv", ';')
  fmt.Println(vlans)
}
```

## Functions

The list below are the functions that are provided for you. There are some others, they are mostly used for testing and development.
### func GetRangeFromString(input string) (low, high int, err error)

The inputstring is a numbered list "from-to", and in return you get either an error or the two numbers in the list.
```
low, high := GetRangefromString("3-10")
```
Will give you low = 3 and high = 10.

### func GetVLANString(input string) (vlans []int, err error)

input a list of vlans in Cisco format and in return you get either a panic, an error or the list of vlan numbers in an array.
```
vlans, _ := GetVLANString("1,10-12,15,20-40,43")
```

Do not add white spaces into the input string.

### func FindVlan(ID int, AllVlans []Vlan) *Vlan

When you have an array of Vlan (see below) you can then search through them to find a VLAN.
```
vlans := MakeTestVlans()
vlan := FindVlan(1, vlans)
```

If the given VLAN ID does not exist in the array, nil is returned.

### func ReadVlanCsv(inputFile string, separator rune) (vlans []Vlan, err error)

Reads an input file and returns either an error or an array of []Vlan that you can work on.

```
vlans, _ := ReadVlanCsv("vlan_test.csv", ';')
```

## Types that are exposed

The Vlan struct is used when you need to work with vlans.
```type Vlan struct {
	ID     int
	Name   string
	Tenant string
	VRF    string
}
```
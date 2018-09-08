package ciscoutils

import (
	"fmt"
	"testing"
)

func TestMakeTestVlans(t *testing.T) {
	vlans := MakeTestVlans()
	if len(vlans) != 2 {
		t.Error("not two VLANs")
	}
}

func TestFindVlan1(t *testing.T) {
	vlan := FindVlan(1, nil)
	if vlan != nil {
		t.Fail()
	}
}

func TestFindVlan2(t *testing.T) {
	vlans := MakeTestVlans()
	// check for first VLAN
	vlan := FindVlan(2, vlans)
	if vlan == nil {
		t.Fail()
	}
	if vlan.ID != 2 {
		t.Error("incorrect VLAN ID returned")
	}
}

func TestFindVlan3(t *testing.T) {
	vlans := MakeTestVlans()
	// check for first VLAN
	vlan := FindVlan(3, vlans)
	if vlan != nil {
		t.Fail()
	}
}

func TestIntArrayEquals(t *testing.T) {
	a := []int{1, 5, 8}
	b := []int{1, 5, 8}
	c := []int{2, 7, 9}
	if IntArrayEquals(a, b) == false {
		t.Fail()
	}
	if IntArrayEquals(a, a) == false {
		t.Fail()
	}
	if IntArrayEquals(a, c) {
		t.Fail()
	}
}

func TestReadVlanCsv(t *testing.T) {
	vlans, err := ReadVlanCsv("vlan_test.csv", ';', "mydomain")
	if err != nil {
		t.Error(err)
	}
	if len(vlans) != 3 {
		t.Error("wrong number of vlans read from file")
	}
	// check VLAN 1
	vlan := FindVlan(1, vlans)
	if vlan == nil {
		t.Error("VLAN 1 not found in CSV")
	} else {
		if vlan.EPG != "MyEPG" {
			t.Error("VLAN 1 missing or incorrect EPG")
		}
	}
	// check VLAN 3
	vlan = FindVlan(10, vlans)
	if vlan == nil {
		t.Error("VLAN 10 not found in CSV")
	} else {
		if vlan.EPG != "" {
			t.Error("VLAN 10 got EPG")
		}
		if vlan.Domain != "mydomain" {
			t.Error("VLAN 10 missing domain")
		}
	}
	fmt.Println(vlans)
}

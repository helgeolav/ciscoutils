package ciscoutils

import (
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
	vlans, err := ReadVlanCsv("vlan_test.csv", ';')
	if err != nil {
		t.Error(err)
	}
	if len(vlans) != 3 {
		t.Error("wrong number of vlans read from file")
	}
}

package ciscoutils

import (
	"encoding/csv"
	"os"
	"strconv"
)

// this is a struct that you can use to keep some information about a VLAN.
type Vlan struct {
	ID     int    // VLAN ID
	Name   string // name of VLAN
	Tenant string // tenant name
	VRF    string // name of VRF
}

// Look through an array of VLANs and find first VLAN with an ID, return nil if not found
func FindVlan(ID int, AllVlans []Vlan) *Vlan {
	if AllVlans == nil {
		return nil
	}
	for _, each := range AllVlans {
		if each.ID == ID {
			return &each
		}
	}
	return nil
}

// this func just returns a list of VLANs that is used for testing code
func MakeTestVlans() []Vlan {
	vl2 := Vlan{2, "VLAN-2-SOMETHING", "TEST-TENANT", "VRF1"}
	vl444 := Vlan{4, "STRANGE-VLAN", "TEST-TENANT", "VRF1"}
	var AllVlans []Vlan
	AllVlans = append(AllVlans, vl2)
	AllVlans = append(AllVlans, vl444)
	return AllVlans
}

// This func reads a CSV file in the format "vlan,name,tenant,vrf"
func ReadVlanCsv(inputFile string, separator rune) (vlans []Vlan, err error) {
	var AllVlans []Vlan
	csvFile, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = separator
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var oneRecord Vlan
	for _, each := range csvData {
		oneRecord.ID, _ = strconv.Atoi(each[0])
		oneRecord.Name = each[1]
		if len(each) > 2 {
			oneRecord.Tenant = each[2]
		} else {
			oneRecord.Tenant = ""
		}
		if len(each) > 3 {
			if each[3] == "" {
				oneRecord.VRF = "VRF1"
			} else {
				oneRecord.VRF = each[3]
			}
		} else {
			oneRecord.VRF = ""
		}
		if oneRecord.ID > 0 && oneRecord.ID < 4095 {
			AllVlans = append(AllVlans, oneRecord)
		}
	}
	return AllVlans, nil
}

package ciscoutils

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Vlan is a struct that you can use to keep some information about a VLAN.
type Vlan struct {
	ID     int    // VLAN ID
	Name   string // name of VLAN
	Tenant string // tenant name
	VRF    string // name of VRF
	Domain string // physical domain
	App    string // Application associated with Vlan
	EPG    string // EPG associated with Vlan
}

// NewVlan return a new Vlan with optional default values
func NewVlan() *Vlan {
	return new(Vlan)
}

// FindVlanDomain Look through an array of Vlans to find a VLAN inside a domain
func FindVlanDomain(ID int, AllVlans *[]Vlan, domain string) *Vlan {
	if AllVlans == nil {
		return nil
	}
	for _, each := range *AllVlans {
		if each.ID == ID && each.Domain == domain {
			return &each
		}
	}
	return nil
}

// FindVlan Look through an array of VLANs and find first VLAN with an ID, return nil if not found
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

// MakeTestVlans this func just returns a list of VLANs that is used for testing code
func MakeTestVlans() []Vlan {
	vl2 := Vlan{2, "VLAN-2-SOMETHING", "TEST-TENANT", "VRF1", "", "", ""}
	vl444 := Vlan{4, "STRANGE-VLAN", "TEST-TENANT", "VRF1", "", "", ""}
	var AllVlans []Vlan
	AllVlans = append(AllVlans, vl2)
	AllVlans = append(AllVlans, vl444)
	return AllVlans
}

// ReadVlanCsv This func reads a CSV file in the format "vlan,name,tenant,vrf,application,EPG"
func ReadVlanCsv(inputFile string, separator rune, domain string) (vlans []Vlan, err error) {
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
	oneRecord.Domain = domain
	for _, each := range csvData {
		oneRecord.ID, _ = strconv.Atoi(each[0])
		oneRecord.Name = each[1]
		// check tenant
		if len(each) > 2 {
			oneRecord.Tenant = each[2]
		} else {
			oneRecord.Tenant = ""
		}
		// check for VRF
		if len(each) > 3 {
			if each[3] == "" {
				oneRecord.VRF = "VRF1"
			} else {
				oneRecord.VRF = each[3]
			}
		} else {
			oneRecord.VRF = ""
		}
		// check for application
		if len(each) > 4 {
			oneRecord.App = each[4]
		} else {
			oneRecord.App = ""
		}
		// check for EPG
		if len(each) > 5 {
			oneRecord.EPG = each[5]
		} else {
			oneRecord.EPG = ""
		}
		// sanity check
		if oneRecord.ID > 0 && oneRecord.ID < 4095 {
			AllVlans = append(AllVlans, oneRecord)
		}
	}
	return AllVlans, nil
}

// GetAppName Return the name of the application
func GetAppName(vlan *Vlan, defaultname string) string {
	// return the default name if no vlan or no App name
	if vlan == nil || vlan.App == "" {
		return defaultname
	}
	return vlan.App
}

// GetEPGName Return the name of the EPG
func GetEPGName(vlan *Vlan) string {
	// return no name if no VLAN
	if vlan == nil {
		return ""
	}
	// return default name if no name
	if vlan.EPG == "" {
		return "VL" + strconv.Itoa(vlan.ID) + "-L2"
	}
	return vlan.EPG
}

// GetVRFName Return the name of the VRF
func GetVRFName(vlan *Vlan) string {
	// return no name if no VLAN
	if vlan == nil {
		return ""
	}
	// return default name if none specified
	if vlan.VRF == "" {
		return "VRF1"
	}
	return vlan.VRF
}

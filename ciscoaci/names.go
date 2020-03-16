package ciscoaci

import "strconv"

// Function to make name to use for the EPG based on a VLAN ID
func MakeEPGName(ID int) string {
	return "VL" + strconv.Itoa(ID) + "-L2"
}

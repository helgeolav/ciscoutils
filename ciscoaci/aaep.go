package ciscoaci

import (
	"fmt"
	"github.com/helgeolav/acigo/aci"
	"github.com/helgeolav/ciscoutils"
	"strconv"
)

// internal function to make name to use for the EPG based on a VLAN ID
func makeEPGName(ID int) string {
	return "VL" + strconv.Itoa(ID) + "-L2"
}

func PutEncapOnAEP(Client *aci.Client, vlan *ciscoutils.Vlan, aep string) {
	// check if we have a VLAN
	if vlan == nil {
		return
	}
	EPG := makeEPGName(vlan.ID)
	encap := aci.GetVLANEncap(vlan.ID)
	Client.AttachableAccessEntityProfileEncapAdd(aep, vlan.Tenant, "L2", EPG, encap)
}

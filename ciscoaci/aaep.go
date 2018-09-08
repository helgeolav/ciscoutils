package ciscoaci

import (
	"github.com/helgeolav/acigo/aci"
	"github.com/helgeolav/ciscoutils"
	"strconv"
)

// internal function to make name to use for the EPG based on a VLAN ID
func makeEPGName(ID int) string {
	return "VL" + strconv.Itoa(ID) + "-L2"
}

// Adds an encap to an AEP from the Vlan array
func PutEncapOnAEP(Client *aci.Client, vlan *ciscoutils.Vlan, aep string) (err error) {
	// check if we have a VLAN
	if vlan == nil {
		return
	}
	EPG := makeEPGName(vlan.ID)
	encap := aci.GetVLANEncap(vlan.ID)
	return Client.AttachableAccessEntityProfileEncapAdd(aep, vlan.Tenant, "L2", EPG, encap)
}

// deletes an encap from an AEP when found in the in the Vlan array
func DeleteEncapOnAEP(Client *aci.Client, vlan *ciscoutils.Vlan, aep string) (err error) {
	// check if we have a VLAN
	if vlan == nil {
		return
	}
	EPG := makeEPGName(vlan.ID)
	encap := aci.GetVLANEncap(vlan.ID)
	return Client.AttachableAccessEntityProfileEncapDel(aep, vlan.Tenant, "L2", EPG, encap)
}

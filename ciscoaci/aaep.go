package ciscoaci

import (
	"github.com/helgeolav/acigo/aci"
	"github.com/helgeolav/ciscoutils"
)


// Adds an encap to an AEP from the Vlan array
func PutEncapOnAEP(Client *aci.Client, vlan *ciscoutils.Vlan, aep string) (err error) {
	// check if we have a VLAN
	if vlan == nil {
		return
	}
	// define APP
	APP := vlan.App
	if APP == "" {
		APP = "L2"
	}
	// define EPG
	EPG := vlan.EPG
	if EPG == "" {
		EPG = MakeEPGName(vlan.ID)
	}

	encap := aci.GetVLANEncap(vlan.ID)
	return Client.AttachableAccessEntityProfileEncapAdd(aep, vlan.Tenant, APP, EPG, encap)
}

// deletes an encap from an AEP when found in the in the Vlan array
func DeleteEncapOnAEP(Client *aci.Client, vlan *ciscoutils.Vlan, aep string) (err error) {
	// check if we have a VLAN
	if vlan == nil {
		return
	}
	EPG := MakeEPGName(vlan.ID)
	encap := aci.GetVLANEncap(vlan.ID)
	return Client.AttachableAccessEntityProfileEncapDel(aep, vlan.Tenant, "L2", EPG, encap)
}

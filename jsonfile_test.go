package ciscoutils

import (
	"os"
	"testing"
)

func TestSaveJsonToFile(t *testing.T) {
	vlan := Vlan{Name: "myname"}
	tmpFile := "dummytestvlan.json"
	err := SaveJsonToFile(tmpFile, &vlan)
	_ = os.Remove(tmpFile)
	if err != nil {
		t.Error(err)
	}
}

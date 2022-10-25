package idforlic_lm

import "testing"

const YOUR_GUID = "5efb766912bd4cb886aabd8800000000"

func TestGetID(t *testing.T) {
	SystemID, SystemIDErr := GetID()
	if SystemIDErr != nil {
		t.Errorf("Err: %s", SystemIDErr)
	} else {
		if SystemID != YOUR_GUID {
			t.Errorf("Got wrong ID: %s", SystemID)
		}
	}
}

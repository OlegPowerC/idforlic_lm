package idforlic_lm

import "testing"

const YOUR_GUID = "1f9c66be-a8d7-4963-0000-3c9db5000000"

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

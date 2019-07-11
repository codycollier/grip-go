package mockload

import "testing"

func TestComputeLoadBasic(t *testing.T) {

	response := "foo"

	response = ComputeLoad(0)
	if response != "0" {
		t.Errorf("Test failed. Expected: 0 Found: %v", response)
	}

	response = ComputeLoad(1)
	if response != "1" {
		t.Errorf("Test failed. Expected: 1 Found: %v", response)
	}
}

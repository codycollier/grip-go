package mockload

import (
	"strconv"
	"time"
)

// Heavily use a cpu/core to represent compute workload
func ComputeLoad(seconds int32) string {

	// TODO: For now we'll sleep, but something better needs to go here that
	// uses compute.
	time.Sleep(time.Duration(seconds) * time.Second)
	response := strconv.Itoa(int(seconds))

	return response
}

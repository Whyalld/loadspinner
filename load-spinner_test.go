package loadspinner

import (
	"time"
	"testing"
)

func doNothing() {
	time.Sleep(5 * time.Second)
}

func TestShowLoading(t *testing.T) {
	stopChan := make(chan bool)

	go ShowLoading(stopChan)

	doNothing()

	stopChan <- true
}
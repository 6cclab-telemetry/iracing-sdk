package winevents

import (
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/winlabs/gowin32/wrappers"
)

var eventHandle syscall.Handle

func OpenEvent(eventName string) error {
	enPtr, err := syscall.UTF16PtrFromString(eventName)
	if err != nil {
		log.Fatalf("Failed to parse eventName %s to UTF16Ptr: %v", eventName, err)
	}

	eventHandle, err = wrappers.OpenEvent(wrappers.SYNCHRONIZE, false, enPtr)
	if err != nil {
		log.Errorf("Failed to open event %s: %v", eventName, err)
		return err
	}
	return nil
}

func WaitForSingleObject(timeout time.Duration) bool {
	absTimeout := time.Now().Add(timeout)
	event, err := wrappers.WaitForSingleObject(eventHandle, uint32(timeout.Milliseconds()))
	if err != nil {
		log.Errorf("Error wait for single Object: %v", err)
		if absTimeout.After(time.Now()) {
			time.Sleep(absTimeout.Sub(time.Now()))
		}
		return false
	}
	log.Debugf("WaitForSingleObject result: %d", event)
	return event == wrappers.WAIT_OBJECT_0
}

func BroadcastMsg(msgName string, msg int, p1 int, p2 interface{}, p3 int) bool {
	// TODO: implement
	return false

}

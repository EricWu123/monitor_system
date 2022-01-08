package testcase

import (
	"monitor_system/test/testtools"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	testtools.TestRunServer()
	time.Sleep(2 * time.Second)
	m.Run()
	testtools.ShutdownServer()
	time.Sleep(time.Second)
}

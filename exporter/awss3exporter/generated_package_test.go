// Code generated by mdatagen. DO NOT EDIT.

package awss3exporter

import (
	"go.uber.org/goleak"
	"testing"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("net/http.(*persistConn).writeLoop"), goleak.IgnoreTopFunction("internal/poll.runtime_pollWait"))
}

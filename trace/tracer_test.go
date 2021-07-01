package trace

import (
    "bytes"
    "testing"
)

func TestNew(t *testing.T) {
    var buf bytes.Buffer
    tracer := New(&buf)
    if tracer == nil {
	t.Error("tracer is nil")
    } else {
	tracer.Trace("hello, tracer package")
	if buf.String() != "hello, tracer package\n" {
	    t.Errorf("'%s' bad string output", buf.String())
	}
    }
}

func TestOff(t *testing.T) {
    var silentTracer Tracer = Off()
    silentTracer.Trace("data")
}

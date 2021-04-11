package protorand

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/sryoya/protoreflet-go-examples/protorand/testproto"
)

func TestNew(t *testing.T) {
	target := &testproto.TestMessage{}

	expected := &testproto.TestMessage{
		SomeInt:   10,
		SomeFloat: 10.1,
		SomeStr:   "Gopher",
		SomeBool:  true,
		SomeMsg: &testproto.ChildMessage{
			SomeInt: 10,
		},
	}

	// inject random generated value to be fixed
	randomInt32 = func() int32 { return 10 }
	randomFloat = func() float32 { return 10.1 }
	randomString = func(int) string { return "Gopher" }
	randomBool = func() bool { return true }

	EmbedValues(target)

	if diff := cmp.Diff(expected, target, protocmp.Transform()); diff != "" {
		t.Errorf("response didn't match (-want / +got)\n%s", diff)
	}
}

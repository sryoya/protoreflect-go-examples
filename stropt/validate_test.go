package stropt

import (
	"testing"

	"github.com/sryoya/protoreflet-go-examples/stropt/testproto"
	"google.golang.org/protobuf/proto"
	// "github.com/sryoya/protoreflet-go-examples/stropt/testproto"
)

func TestIsRefund(t *testing.T) {
	cases := map[string]struct {
		input proto.Message
		isErr bool
	}{
		"Return true if Refund": {
			input: &testproto.TestMessage{
				SomeInt: 10,
				SomeStr: "Gopher",
				SomeMsg: &testproto.ChildMessage{
					SomeInt: 10,
				},
			},
			isErr: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// execution
			res := Validate(tc.input)

			if tc.isErr == (res != nil) {
				t.Errorf("expected error: %v, but got:%v", tc.isErr, res)
			}
		})
	}
}

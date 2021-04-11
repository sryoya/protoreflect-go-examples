package stropt

import (
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/sryoya/protoreflet-go-examples/stropt/testproto"
)

func TestIsRefund(t *testing.T) {
	cases := map[string]struct {
		input proto.Message
		isErr bool
	}{
		"Return true if Refund": {
			input: &testproto.TestMessage{
				SomeStr: "123456",
			},
			isErr: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// execution
			res := Validate(tc.input)

			if tc.isErr != (res == nil) {
				t.Errorf("expecting error is %v, but got:%v", tc.isErr, res)
			}
		})
	}
}

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
		"unmatch with expected length": {
			input: &testproto.MessageFixedLen{
				Value: "12345", // expects 6
			},
			isErr: true,
		},
		"smaller than min length": {
			input: &testproto.MessageMaxMinLen{
				Value: "12", // expects 3~5
			},
			isErr: true,
		},
		"bigger than max length": {
			input: &testproto.MessageMaxMinLen{
				Value: "123456", // expects 3~5
			},
			isErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// execution
			res := Validate(tc.input)

			if tc.isErr == (res == nil) {
				t.Errorf("expecting error is %v, but got:%v", tc.isErr, res)
			}
		})
	}
}

package stropt

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/sryoya/protoreflect-go-examples/stropt/testproto"
)

func TestValidate(t *testing.T) {
	cases := map[string]struct {
		input      proto.Message
		wantErrStr string
	}{
		// length validations
		"valid with the specifed length": {
			input: &testproto.MessageFixedLen{
				Value: "123456", // expects 6
			},
		},
		"unmatch with expected length": {
			input: &testproto.MessageFixedLen{
				Value: "12345", // expects 6
			},
			wantErrStr: "expected: 6, actual: 5",
		},
		"valid within the specified length": {
			input: &testproto.MessageMaxMinLen{
				Value: "1234", // expects 3~5
			},
		},
		"smaller than min length": {
			input: &testproto.MessageMaxMinLen{
				Value: "12", // expects 3~5
			},
			wantErrStr: "longer than or equal to",
		},
		"bigger than max length": {
			input: &testproto.MessageMaxMinLen{
				Value: "123456", // expects 3~5
			},
			wantErrStr: "shorter than or equal to",
		},
		"invalid proto, min length > max length": {
			input: &testproto.MessageInvalidMaxMinLen{
				Value: "", // irrevant
			},
			wantErrStr: "min Length outweighs max Length",
		},
		"invalid value in list against the specified len": {
			input: &testproto.MessageWithList{
				ListValues: []string{"123456", "12"}, // expects 6
			},
			wantErrStr: "invalid length, expected: 6, actual: 2",
		},
		"invalid value in map against the specified": {
			input: &testproto.MessageWithMap{
				MapValues: map[int32]string{ // expects 6
					1: "123456",
					2: "12",
				},
			},
			wantErrStr: "invalid length, expected: 6, actual: 2",
		},
		// recursion check
		"invalid value in child message": {
			input: &testproto.ParentMessage{
				Child: &testproto.ChildMessage{
					Value: "12345", // expects 6
				},
			},
			wantErrStr: "expected: 6, actual: 5",
		},
		// common check
		"invalid proto, stropt appended in non-string field": {
			input: &testproto.MessageStrOptOnInt{
				Value: 1,
			},
			wantErrStr: "stropt is appended to non-string field",
		},
		"invalid proto, stropt appended in non-string map": {
			input: &testproto.MessageStrOptOnIntMap{
				MapValues: map[int32]int32{
					1: 123456, // irrevant
				},
			},
			wantErrStr: "stropt is appended to non-string map",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res := Validate(tc.input)

			if tc.wantErrStr == "" {
				if res != nil {
					t.Errorf("unexpected error was returned: %v", res)
				}
				return
			}
			if tc.wantErrStr != "" && res == nil {
				t.Errorf("error was expected, but got nil")
				return
			}

			if !strings.Contains(res.Error(), tc.wantErrStr) {
				t.Errorf("expecting error containing %v, but got:%v", tc.wantErrStr, res)
			}
		})
	}
}

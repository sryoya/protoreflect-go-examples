package main

import (
	"fmt"

	"github.com/sryoya/protoreflet-go-examples/stropt"
	"github.com/sryoya/protoreflet-go-examples/stropt/testproto"
)

func main() {
	// 1. Length validation
	//
	// The used proto message is as follows:
	//
	// message UpdateAccountIDRequest {
	// 	string iD = 1 [ (stroptpb.opts) = {min_len : 6, max_len : 12} ];
	// }

	// ok
	msg := &testproto.UpdateAccountIDRequest{
		ID: "12345678",
	}
	err := stropt.Validate(msg)
	fmt.Println(err)
	// Output: nil

	// shorted than min length
	msg = &testproto.UpdateAccountIDRequest{
		ID: "12345",
	}
	err = stropt.Validate(msg)
	fmt.Println(err)
	// Output: Field: ID, invalid length, the value must be longer than or equal to 6, but actual: 5

	// longer than max length
	msg = &testproto.UpdateAccountIDRequest{
		ID: "1234567890123",
	}
	err = stropt.Validate(msg)
	fmt.Println(err)
	// Output: Field: ID, invalid length, the value must be shorter than or equal to 12, but actual: 13
}

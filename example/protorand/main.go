package main

import (
	"fmt"

	"github.com/sryoya/protoreflect-go-examples/example/testproto"
	"github.com/sryoya/protoreflect-go-examples/protorand"
)

func main() {
	msg := &testproto.TestMessage{}
	protorand.EmbedValues(msg)
	fmt.Println(msg)
	// It outputs the message with the random value like this
	// some_str:"tj30CLH0Sp"
	// some_int:2039157080
	// some_float:0.86438453
	// some_slice:"O34vxOpuVV"
	// some_msg:{
	//	some_int:146553370
	// }
	// some_msgs:{
	// some_int:1015307036
	//}
	// some_map:{
	//	key:862599133
	//	value:{some_int:691336121}
	//}
}

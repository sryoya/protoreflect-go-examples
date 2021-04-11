package protorand

import (
	"fmt"
	"math/rand"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/sryoya/protoreflet-go-examples/protorand/testproto"
)

var (
	chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randomStr = randomString
	ran       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func New() *testproto.TestMessage {
	input := &testproto.TestMessage{}
	EmbedRandValue(input)
	return input
}

// EmbedRandValue embeds randoms value to fields in the prvoded proto message
func EmbedRandValue(msg proto.Message) error {
	pm := msg.ProtoReflect()
	mds := pm.Descriptor()
	fds := mds.Fields()
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)
		switch fd.Kind() {
		case protoreflect.Int32Kind:
			pm.Set(fd, protoreflect.ValueOfInt32(ran.Int31()))
		case protoreflect.FloatKind:
			pm.Set(fd, protoreflect.ValueOfFloat32(ran.Float32()))
		case protoreflect.StringKind:
			pm.Set(fd, protoreflect.ValueOfString(randomString(10)))
		case protoreflect.BoolKind:
			pm.Set(fd, protoreflect.ValueOfBool(randomBool()))
		case protoreflect.MessageKind:
			// pm.Set(fd, protoreflect.ValueOfMessage(fd.Message()))
		default:
			return fmt.Errorf("unexpected type: %v", fd.Kind())
		}
	}

	return nil
}

// We can implement the above func using "dynamicpb"
//
// // EmbedRandValue embeds randoms value to fields in the prvoded proto message
// func EmbedRandValue(msg proto.Message) error {
// 	mds := msg.ProtoReflect().Descriptor()
// 	dm := dynamicpb.NewMessage(mds)
// 	fds := mds.Fields()
// 	for k := 0; k < fds.Len(); k++ {
// 		fd := fds.Get(k)
// 		switch fd.Kind() {
// 		case protoreflect.Int32Kind:
// 			dm.Set(fd, protoreflect.ValueOfInt32(ran.Int31()))
// 		case protoreflect.FloatKind:
// 			dm.Set(fd, protoreflect.ValueOfFloat32(ran.Float32()))
// 		case protoreflect.StringKind:
// 			dm.Set(fd, protoreflect.ValueOfString(randomString(10)))
// 		case protoreflect.BoolKind:
// 			dm.Set(fd, protoreflect.ValueOfBool(randomBool()))
// 		case protoreflect.MessageKind:
// 			// TODO: make it recursive
// 		default:
// 			return fmt.Errorf("unexpected type: %v", fd.Kind())
// 		}
// 	}

// 	proto.Merge(msg, dm)
// 	return nil
// }

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func randomBool() bool {
	return ran.Int31()%2 == 0
}

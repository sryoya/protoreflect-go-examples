package protorand

import (
	"fmt"
	"math/rand"
	"time"

	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

var (
	chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randomStr = randomString
	ran       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// FuzzMessage embeds randoms value to fields in the prvoded proto message
func FuzzMessage(msg proto.Message) error {
	mds := msg.ProtoReflect().Descriptor()
	dm := dynamicpb.NewMessage(mds)
	fds := mds.Fields()
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)
		switch fd.Kind() {
		case protoreflect.Int32Kind:
			dm.Set(fd, protoreflect.ValueOfInt32(ran.Int31()))
		case protoreflect.FloatKind:
			dm.Set(fd, protoreflect.ValueOfFloat32(ran.Float32()))
		case protoreflect.StringKind:
			dm.Set(fd, protoreflect.ValueOfString(randomString(10)))
		case protoreflect.BoolKind:
			dm.Set(fd, protoreflect.ValueOfBool(randomBool()))
		case protoreflect.MessageKind:
			// TODO: make it recursive
		default:
			return fmt.Errorf("unexpected type: %v", fd.Kind())
		}
	}

	proto.Merge(msg, dm)
	return nil
}

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

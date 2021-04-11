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

	randomInt32  = genRandInt32
	randomFloat  = genRandFloat
	randomString = genRandString
	randomBool   = genRandBool
	// ran       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// EmbedRandValue embeds randoms value to fields in the provided proto message
func EmbedValues(msg proto.Message) error {
	mds := msg.ProtoReflect().Descriptor()
	dm, err := NewDynamicProto(mds)
	if err != nil {
		return nil
	}

	proto.Merge(msg, dm)
	return nil
}

// NewDynamicProto embeds value to Proto
func NewDynamicProto(mds protoreflect.MessageDescriptor) (*dynamicpb.Message, error) {
	dm := dynamicpb.NewMessage(mds)
	fds := mds.Fields()
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)
		switch fd.Kind() {
		case protoreflect.Int32Kind:
			dm.Set(fd, protoreflect.ValueOfInt32(randomInt32()))
		case protoreflect.FloatKind:
			dm.Set(fd, protoreflect.ValueOfFloat32(randomFloat()))
		case protoreflect.StringKind:
			dm.Set(fd, protoreflect.ValueOfString(randomString(10)))
		case protoreflect.BoolKind:
			dm.Set(fd, protoreflect.ValueOfBool(randomBool()))
		case protoreflect.MessageKind:
			// process recursively
			rm, err := NewDynamicProto(fd.Message())
			if err != nil {
				return nil, err
			}
			dm.Set(fd, protoreflect.ValueOfMessage(rm))
		default:
			return nil, fmt.Errorf("unexpected type: %v", fd.Kind())
		}
	}

	return dm, nil
}

func genRandInt32() int32 {
	return rand.Int31()
}

func genRandFloat() float32 {
	return rand.Float32()
}

func genRandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func genRandBool() bool {
	return rand.Int31()%2 == 0
}

// Version without using dynamicpb.
// It's fine, but difficult to execute recursively
//
// // EmbedRandValue embeds randoms value to fields in the prvoded proto message
// func EmbedRandValue(msg proto.Message) error {
// 	pm := msg.ProtoReflect()
// 	mds := pm.Descriptor()
// 	fds := mds.Fields()
// 	for k := 0; k < fds.Len(); k++ {
// 		fd := fds.Get(k)
// 		switch fd.Kind() {
// 		case protoreflect.Int32Kind:
// 			pm.Set(fd, protoreflect.ValueOfInt32(ran.Int31()))
// 		case protoreflect.FloatKind:
// 			pm.Set(fd, protoreflect.ValueOfFloat32(ran.Float32()))
// 		case protoreflect.StringKind:
// 			pm.Set(fd, protoreflect.ValueOfString(randomString(10)))
// 		case protoreflect.BoolKind:
// 			pm.Set(fd, protoreflect.ValueOfBool(randomBool()))
// 		case protoreflect.MessageKind:
// 			// pm.Set(fd, protoreflect.ValueOfMessage(fd.Message()))
// 		default:
// 			return fmt.Errorf("unexpected type: %v", fd.Kind())
// 		}
// 	}

// 	return nil
// }

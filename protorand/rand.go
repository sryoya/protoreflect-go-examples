package protorand

import (
	"fmt"
	"math/rand"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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

// // EmbedValues embeds randoms value to fields in the provided proto message
// func EmbedValues(msg proto.Message) error {
// 	mds := msg.ProtoReflect().Descriptor()
// 	dm, err := NewDynamicProtoRand(mds)
// 	if err != nil {
// 		return nil
// 	}

// 	proto.Merge(msg, dm)
// 	return nil
// }

// // NewDynamicProtoRand created dynamicpb with assiging random value to proto
// func NewDynamicProtoRand(mds protoreflect.MessageDescriptor) (*dynamicpb.Message, error) {
// 	dm := dynamicpb.NewMessage(mds)
// 	fds := mds.Fields()
// 	for k := 0; k < fds.Len(); k++ {
// 		fd := fds.Get(k)

// 		if fd.IsList() {
// 			// TODO
// 			continue
// 		}

// 		if fd.IsMap() {
// 			// TODO
// 			continue
// 		}

// 		switch fd.Kind() {
// 		case protoreflect.Int32Kind:
// 			dm.Set(fd, protoreflect.ValueOfInt32(randomInt32()))
// 		case protoreflect.FloatKind:
// 			dm.Set(fd, protoreflect.ValueOfFloat32(randomFloat()))
// 		case protoreflect.StringKind:
// 			dm.Set(fd, protoreflect.ValueOfString(randomString(10)))
// 		case protoreflect.BoolKind:
// 			dm.Set(fd, protoreflect.ValueOfBool(randomBool()))
// 		case protoreflect.MessageKind:
// 			// process recursively
// 			rm, err := NewDynamicProtoRand(fd.Message())
// 			if err != nil {
// 				return nil, err
// 			}
// 			dm.Set(fd, protoreflect.ValueOfMessage(rm))
// 		default:
// 			return nil, fmt.Errorf("unexpected type: %v", fd.Kind())
// 		}
// 	}

// 	return dm, nil
// }

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

// EmbedValues embeds randoms value to fields in the provided proto message
func EmbedValues(msg proto.Message) error {
	pm := msg.ProtoReflect()
	return embedValues(pm)
}

// embedValues embeds randoms value to fields in the provioded protoreflect message
func embedValues(pm protoreflect.Message) error {
	var getRandValue func(fd protoreflect.FieldDescriptor) (protoreflect.Value, error)
	getRandValue = func(fd protoreflect.FieldDescriptor) (protoreflect.Value, error) {
		if fd.IsList() {
			list := pm.Mutable(fd).List()
			// TODO: implement
			return protoreflect.ValueOfList(list), nil
		}
		if fd.IsMap() {
			mp := pm.Mutable(fd).Map()
			// TODO: make the number of elements randomly
			key, err := getRandValue(fd.MapKey())
			if err != nil {
				return protoreflect.Value{}, err
			}
			value, err := getRandValue(fd.MapValue())
			if err != nil {
				return protoreflect.Value{}, err
			}
			mp.Set(protoreflect.MapKey(key), protoreflect.Value(value))
			return protoreflect.ValueOfMap(mp), nil
		}

		switch fd.Kind() {
		case protoreflect.Int32Kind:
			return protoreflect.ValueOfInt32(randomInt32()), nil
		case protoreflect.FloatKind:
			return protoreflect.ValueOfFloat32(randomFloat()), nil
		case protoreflect.StringKind:
			return protoreflect.ValueOfString(randomString(10)), nil
		case protoreflect.BoolKind:
			return protoreflect.ValueOfBool(randomBool()), nil
		case protoreflect.MessageKind:
			// process recursively
			child := pm.Mutable(fd).Message()
			err := embedValues(child)
			if err != nil {
				return protoreflect.Value{}, err
			}
			return protoreflect.ValueOfMessage(child), nil
		default:
			return protoreflect.Value{}, fmt.Errorf("unexpected type: %v", fd.Kind())
		}
	}

	fds := pm.Descriptor().Fields()
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)

		val, err := getRandValue(fd)
		if err != nil {
			return err
		}
		pm.Set(fd, val)
	}

	return nil
}

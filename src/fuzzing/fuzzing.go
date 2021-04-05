package fuzzing

import (
	"log"
	"math/rand"
	"time"

	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func New() *TestMessage {
	input := &TestMessage{}
	res := Message(input.ProtoReflect())

	proto.Merge(input, res)
	return input
}

// Message is a function that returns a random dyanmicpb.Message constructed from a protoreflect.MessageDescriptor
func Message(m protoreflect.Message) *dynamicpb.Message {
	dm := dynamicpb.NewMessage(m.Descriptor())
	fds := m.Descriptor().Fields()
	// Iterate over *all* fields
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)
		v := func() protoreflect.Value {
			var result protoreflect.Value
			switch fd.Kind() {
			case protoreflect.Int32Kind:
				log.Println("[random:New] Int32")
				result = protoreflect.ValueOfInt32(r.Int31())
			case protoreflect.FloatKind:
				log.Println("[random:New] Float32")
				result = protoreflect.ValueOfFloat32(r.Float32())
			case protoreflect.MessageKind:
				// TODO: make it recursive
			case protoreflect.StringKind:
				log.Println("[random:New] String")
				result = protoreflect.ValueOfString("X")
			default:
				log.Fatal("[random:New] unanticipated kind")
			}
			return result
		}()
		dm.Set(fd, v)
	}
	return dm
}

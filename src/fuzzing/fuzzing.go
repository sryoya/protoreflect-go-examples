package fuzzing

import (
	"fmt"
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
	Message(input)
	return input
}

// Message is a function that returns a random dyanmicpb.Message constructed from a protoreflect.MessageDescriptor
func Message(msg proto.Message) error {
	mds := msg.ProtoReflect().Descriptor()
	dm := dynamicpb.NewMessage(mds)
	fds := mds.Fields()
	// Iterate over *all* fields
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)
		switch fd.Kind() {
		case protoreflect.Int32Kind:
			dm.Set(fd, protoreflect.ValueOfInt32(r.Int31()))
		case protoreflect.FloatKind:
			dm.Set(fd, protoreflect.ValueOfFloat32(r.Float32()))
		case protoreflect.MessageKind:
			// TODO: make it recursive
		case protoreflect.StringKind:
			dm.Set(fd, protoreflect.ValueOfString("X"))
		default:
			return fmt.Errorf("unexpected type: %v", fd.Kind())
		}
	}

	proto.Merge(msg, dm)
	return nil
}

package stropt

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/sryoya/protoreflet-go-examples/protobuf/stroptpb"
)

// Validate checks if the proto message has the valid value compliant to the setting in stropt
func Validate(pb proto.Message) error {
	m := pb.ProtoReflect()
	var errs error
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		opts := fd.Options().(*descriptorpb.FieldOptions)
		so, ok := proto.GetExtension(opts, stroptpb.E_Opts).(*stroptpb.StringOpts)
		if !ok {
			return true
		}

		strVal, ok := m.Get(fd).Interface().(string)
		if !ok {
			return true
		}

		err := validateLength(so, strVal)
		if err != nil {
			errs = fmt.Errorf("%v\n%v", errs, err)
		}

		return true
	})

	return errs
}

func validateLength(opts *stroptpb.StringOpts, v string) error {
	valueLen := len(v)

	// Fixed length check
	// If this check is available, further length checks won't be processed
	fixedLen := opts.Len
	if fixedLen != nil {
		if valueLen != int(*fixedLen) {
			return fmt.Errorf("invalid length, expected: %v, actual: %v", fixedLen, valueLen)
		}
		return nil
	}

	// Min/Max length check
	minLen := opts.MinLen
	maxLen := opts.MaxLen
	checkMinLen := opts.MinLen != nil
	checkMaxLen := opts.MaxLen != nil
	if !checkMinLen && !checkMaxLen {
		return nil
	}
	if checkMinLen && checkMaxLen && *maxLen <= *minLen {
		return fmt.Errorf("invalid length setup in proto file, min Length outweighs max Length")
	}
	if valueLen < int(*minLen) {
		return fmt.Errorf("invalid length, expected longer than or equal to:%v, but actual: %v", minLen, valueLen)
	}
	if valueLen > int(*maxLen) {
		return fmt.Errorf("invalid length, expected shorter than or equal to:%v, but actual: %v", maxLen, valueLen)
	}

	return nil
}

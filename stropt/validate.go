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
	return validate(pb.ProtoReflect())
}

// Validate checks if the proto message has the valid value compliant to the setting in stropt
func validate(m protoreflect.Message) error {
	md := m.Descriptor()
	fds := md.Fields()
	var errs error
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)
		// call recursively
		if fd.Kind() == protoreflect.MessageKind && !fd.IsMap() {
			errs = appendErr(errs, validate(m.Get(fd).Message()))
			continue
		}

		// get Stropt if it's set
		opts := fd.Options().(*descriptorpb.FieldOptions)
		so, ok := proto.GetExtension(opts, stroptpb.E_Opts).(*stroptpb.StringOpts)
		if !ok || so == nil {
			continue
		}

		// check if the field is string
		if fd.Kind() != protoreflect.StringKind && !fd.IsMap() {
			errs = appendErr(errs, fmt.Errorf("invalid proto def, stropt is appended to non-string field, field name:%v", fd.Name()))
			continue
		}

		if fd.IsList() {
			strList := m.Get(fd).List()
			for i := 0; i < strList.Len(); i++ {
				strVal := strList.Get(i).Interface().(string)
				errs = appendErr(errs, validateValue(fd.Name(), so, strVal))
			}
			continue
		}

		if fd.IsMap() {
			strMap := m.Get(fd).Map()
			// strMap := m.Get(fd).Interface().(map[interface{}]string)
			strMap.Range(func(mk protoreflect.MapKey, mv protoreflect.Value) bool {
				strVal, ok := mv.Interface().(string)
				if !ok {
					errs = appendErr(errs, fmt.Errorf("invalid proto def, stropt is appended to non-string map field, field name:%v", fd.Name()))
					return false
				}
				errs = appendErr(errs, validateValue(fd.Name(), so, strVal))
				return true
			})
			// for _, strVal := range strMap {
			// 	errs = appendErr(errs, validateValue(fd.Name(), so, strVal))
			// }
			continue
		}

		strVal := m.Get(fd).Interface().(string)
		errs = appendErr(errs, validateValue(fd.Name(), so, strVal))
	}
	return errs
}

func validateValue(fieldName protoreflect.Name, opts *stroptpb.StringOpts, v string) error {
	var errs error
	errs = appendErr(errs, validateLength(opts, v))
	// TODO: add other validations

	if errs != nil {
		errs = fmt.Errorf("\nField: %v, %v", fieldName, errs)
	}

	return errs
}

func validateLength(opts *stroptpb.StringOpts, v string) error {
	actualLen := len(v)

	// Fixed length check
	// If this check is available, further length checks won't be processed
	fixedLen := opts.Len
	if fixedLen != nil {
		if actualLen != int(*fixedLen) {
			return fmt.Errorf("invalid length, expected: %v, actual: %v", int(*fixedLen), actualLen)
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
	if checkMinLen && checkMaxLen && *maxLen < *minLen {
		return fmt.Errorf("invalid length setup in proto file, min Length outweighs max Length")
	}
	if checkMinLen && actualLen < int(*minLen) {
		return fmt.Errorf("invalid length, the value must be longer than or equal to %v, but actual: %v", int(*minLen), actualLen)
	}
	if checkMaxLen && actualLen > int(*maxLen) {
		return fmt.Errorf("invalid length, the value must be shorter than or equal to %v, but actual: %v", int(*maxLen), actualLen)
	}

	return nil
}

func appendErr(original, new error) error {
	if new == nil {
		return original
	}
	if original == nil {
		return new
	}
	return fmt.Errorf("%v;%v", new, original)
}

// We can implement the function using."m.Range", but m.Range ignores the default value
// // Validate checks if the proto message has the valid value compliant to the setting in stropt
// func Validate(pb proto.Message) error {
// 	m := pb.ProtoReflect()
// 	var errs error
// 	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
// 		// call recursively
// 		if fd.Kind() == protoreflect.MessageKind {
// 			appendErr(errs, Validate(fd.Message().(proto.Message)))
// 		}

// 		// get Stropt if it's set
// 		opts := fd.Options().(*descriptorpb.FieldOptions)
// 		so, ok := proto.GetExtension(opts, stroptpb.E_Opts).(*stroptpb.StringOpts)
// 		if !ok {
// 			return true
// 		}
// 		// get value in the field
// 		strVal, ok := m.Get(fd).Interface().(string)
// 		if !ok {
// 			appendErr(errs, fmt.Errorf("invalid proto def, stropt is appended to non-string field, field name:%v", fd.Name()))
// 			return true
// 		}
// 		// validate
// 		errs = appendErr(errs, validateValue(fd.Name(), so, strVal))
// 		return true
// 	})
// 	return errs
// }

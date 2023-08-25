package parser

import (
	"github.com/vektah/gqlparser/v2/ast"
	"log"
	"strconv"
	"strings"
)

type Value interface {
	String() string
	GoString() string
	JSString() string
	HCLString() string
}

func parseValue(v *ast.Value) Value {
	if v == nil {
		return nil
	}
	switch v.Kind {
	case ast.StringValue:
		return &valMarshallable{val: v.Raw}
	case ast.IntValue, ast.FloatValue, ast.BlockValue, ast.EnumValue:
		return &valRaw{val: v.Raw}
	case ast.BooleanValue:
		parsed, err := strconv.ParseBool(v.Raw)
		if err != nil {
			log.Fatal(err)
		}
		return &valBool{val: parsed}
	case ast.NullValue:
		return &valNull{}
	case ast.ListValue:
		log.Fatal("ListValue currently unsupported")
		return &valList{vals: nil}
	default:
		log.Printf("unhandled parseValue type: %d %s", v.Kind, v.Raw)
		return nil
	}
}

type valList struct {
	vals []Value
	typ  *FieldType
}

func (v *valList) String() string {
	conv := make([]string, 0, len(v.vals))
	for _, c := range v.vals {
		conv = append(conv, c.String())
	}
	return strings.Join(conv, ", ")
}

func (v *valList) GoString() string {
	conv := make([]string, 0, len(v.vals))
	for _, c := range v.vals {
		conv = append(conv, c.GoString())
	}
	return "[]" + v.typ.GoType() + "{" + strings.Join(conv, ", ") + "}"
}

func (v *valList) JSString() string {
	conv := make([]string, 0, len(v.vals))
	for _, c := range v.vals {
		conv = append(conv, c.JSString())
	}
	return "[" + strings.Join(conv, ", ") + "]"
}

func (v *valList) HCLString() string {
	conv := make([]string, 0, len(v.vals))
	for _, c := range v.vals {
		conv = append(conv, c.HCLString())
	}
	return "[" + strings.Join(conv, ", ") + "]"
}

type valMarshallable struct {
	val string
}

func (v *valMarshallable) String() string {
	return v.val
}

func (v *valMarshallable) GoString() string {
	return strconv.Quote(v.val)
}

func (v *valMarshallable) JSString() string {
	return strconv.Quote(v.val)
}

func (v *valMarshallable) HCLString() string {
	return strconv.Quote(v.val)
}

type valRaw struct {
	val string
}

func (v *valRaw) String() string {
	return v.val
}

func (v *valRaw) GoString() string {
	return v.val
}

func (v *valRaw) JSString() string {
	return v.val
}

func (v *valRaw) HCLString() string {
	return v.val
}

type valBool struct {
	val bool
}

func (v *valBool) String() string {
	if v.val {
		return "true"
	} else {
		return "false"
	}
}

func (v *valBool) GoString() string {
	if v.val {
		return "true"
	} else {
		return "false"
	}
}

func (v *valBool) JSString() string {
	if v.val {
		return "true"
	} else {
		return "false"
	}
}

func (v *valBool) HCLString() string {
	if v.val {
		return "true"
	} else {
		return "false"
	}
}

type valNull struct {
}

func (v *valNull) String() string {
	return "null"
}

func (v *valNull) GoString() string {
	return "nil"
}

func (v *valNull) JSString() string {
	return "null"
}

func (v *valNull) HCLString() string {
	return "null"
}

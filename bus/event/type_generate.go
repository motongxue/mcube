// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package event

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	enumTypeShowMap = map[Type]string{
		OperateEventType: "operate",
		AlertEventType:   "alert",
	}

	enumTypeIDMap = map[string]Type{
		"operate": OperateEventType,
		"alert":   AlertEventType,
	}
)

// ParseType Parse Type from string
func ParseType(str string) (Type, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := enumTypeIDMap[key]
	if !ok {
		return 0, fmt.Errorf("unknown Status: %s", str)
	}

	return v, nil
}

// Is todo
func (t Type) Is(target Type) bool {
	return t == target
}

// String stringer
func (t Type) String() string {
	v, ok := enumTypeShowMap[t]
	if !ok {
		return "unknown"
	}

	return v
}

// MarshalJSON todo
func (t Type) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(t.String())
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Type) UnmarshalJSON(b []byte) error {
	ins, err := ParseType(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

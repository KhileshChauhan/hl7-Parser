package hl7

import (
	"bytes"
)

// Encode create an Hl7 message
func Encode(hl7Message Hl7Message) ([]byte, error) {
	var ls [][]byte
	for _, hl7Segment := range hl7Message {
		bSegment := encodeSegment(hl7Segment)
		ls = append(ls, bSegment)
	}
	return bytes.Join(ls, []byte("\n")), nil
}

func encodeSegment(hl7Segment Hl7Segment) []byte {
	var ls [][]byte
	for _, hl7Repeated := range hl7Segment {
		bSegment := encodeRepeated(hl7Repeated)
		ls = append(ls, bSegment)
	}
	return bytes.Join(ls, []byte("|"))
}
func encodeRepeated(hl7Repeated Hl7Repeated) []byte {
	var ls [][]byte
	for _, hl7Component := range hl7Repeated {
		bRepeated := encodeComponent(hl7Component)
		ls = append(ls, bRepeated)
	}
	return bytes.Join(ls, []byte("~"))
}

func encodeComponent(hl7Component Hl7Component) []byte {
	var ls [][]byte
	for _, hl7SubComponent := range hl7Component {
		bComponent := encodeSubcomponent(hl7SubComponent)
		ls = append(ls, bComponent)
	}
	return bytes.Join(ls, []byte("^"))
}

func encodeSubcomponent(hl7SubComponent Hl7SubComponent) []byte {
	return joinFields(hl7SubComponent)
}

func joinFields(s Hl7SubComponent) []byte {
	sep := separators[3]
	if len(s) == 0 {
		return Hl7Field{}
	}
	if len(s) == 1 {
		// Just return a copy.
		return append(Hl7Field(nil), s[0]...)
	}
	n := len(sep) * (len(s) - 1)
	for _, v := range s {
		n += len(v)
	}

	b := make(Hl7Field, n)
	bp := copy(b, s[0])
	for _, v := range s[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], v)
	}
	return b
}

package hl7

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type (
	escapeCouple [][]byte

	Hl7Message      []Hl7Segment
	Hl7Segment      []Hl7Repeated
	Hl7Repeated     []Hl7Component
	Hl7Component    []Hl7SubComponent
	Hl7SubComponent []Hl7Field
	Hl7Field        []byte
)

var (
	separators = [][]byte{
		[]byte("|"), // separate repeated lists in a segment
		[]byte("~"), // separate components in a repeated list
		[]byte("^"), // separate sub-components in a component
		[]byte("&"), // separate fields in a sub-component
	}
	// special characters
	// http://healthstandards.com/blog/2006/11/02/hl7-escape-sequences/
	escapeCouples = []escapeCouple{
		{[]byte("\\F\\"), []byte("|")},
		{[]byte("\\R\\"), []byte("~")},
		{[]byte("\\S\\"), []byte("^")},
		{[]byte("\\T\\"), []byte("&")},
	}
)

// Decode takes the message passed and returns the segments of the hl7 message.
func Decode(message []byte) (Hl7Message, error) {

	if len(message) == 0 {
		return nil, errors.New("[E0002] No data to unmarshal")
	}

	bSegments := LoadMessage(message)

	var hl7Message Hl7Message
	for _, bSegment := range bSegments {
		segment := DecodeSegment(bSegment)
		hl7Message = append(hl7Message, segment)
	}

	return hl7Message, nil
}

// DecodeSegment takes the segment passed in []byte and returns an hl7 segment.
func DecodeSegment(bSegment []byte) Hl7Segment {
	var segment Hl7Segment // contain the whole segment
	for _, bRepeated := range bytes.Split(bSegment, separators[0]) {
		var repeated Hl7Repeated
		for _, bComponent := range bytes.Split(bRepeated, separators[1]) {
			var component Hl7Component
			for _, bSubcomponent := range bytes.Split(bComponent, separators[2]) {
				var subComponent Hl7SubComponent
				for _, bField := range bytes.Split(bSubcomponent, separators[3]) {
					// escape special characters
					if bytes.Contains(bField, []byte("\\")) {
						for _, escapeCouple := range escapeCouples {
							bField = bytes.Replace(bField, escapeCouple[0], escapeCouple[1], -1)
						}
					}
					subComponent = append(subComponent, bField)
				}
				component = append(component, subComponent)
			}
			repeated = append(repeated, component)
		}
		segment = append(segment, repeated)
	}
	return segment
}

func (hl7Segment Hl7Segment) AtIndex(path string) (ret string) {
	defer func() {
		if recover() != nil {
			ret = ""
		}
	}()
	idxs := strings.Split(path, ".")
	idxs_int := make([]int, len(idxs))
	for i, idx := range idxs {
		v, _ := strconv.Atoi(idx)
		idxs_int[i] = v
	}
	return string(hl7Segment[idxs_int[0]][idxs_int[1]][idxs_int[2]][idxs_int[3]])
}

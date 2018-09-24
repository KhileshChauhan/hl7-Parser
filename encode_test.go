package hl7_test

import (
	"testing"

	hl7 "github.com/zquangu112z/hl7-parser"
)

func TestEncode(t *testing.T) {
	// Extract messages from file. Each message has several line
	Hl7Message, err := hl7.Decode(message)
	if err != nil {
		t.Error(err)
	}
	bMessage, err := hl7.Encode(Hl7Message)
	if err != nil {
		t.Error(err)
	}
	if string(bMessage) != string(message) {
		t.Error("[E0005] The created message does not meet the provided message")
	}
}

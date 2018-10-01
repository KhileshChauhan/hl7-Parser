package hl7_test

import (
	"testing"

	hl7 "github.com/zquangu112z/hl7-parser"
)

var (
	message = []byte(`MSH|^~\&|LCAMC|LABCORP|FAKE ACO|FAKE HOSPITAL|20170418093130||ORU^R01|M17108000000000001|P|2.3|||ER|ER
PID|1|11111111111|11111111111|11111111111|Truong^Nicholas||19510610|M|||1511 MONTE VISTA ST^^PASADENA^CA^91106|||||||6337494512380
PV1|1|O||||||1598879918^Khanh Hoang^Minh^^^^^^^^^^NPI|||||||||||||||||||||||||||||||SO
ORC||633749453380|633749453380||||||20161202|||1598879918^Fake name^Fake given name^^^^^^^^^^NPI
OBR|1||633749453380|005009^CBC WITH DIFFERENTIAL/PLATELET^L|||20161202|20161202||||||||1598879918^Fake name^Fake given name^^^^^^^^^^NPI||SO|||||||F
OBX|1|ST|6690-2^LOINC^LN^005025^WBC^L||8.7|X10E3/UL|3.4-10.8||||F|||20161202|SO   ^^L
OBX|1|ST|6690-2^LOINC^LN^005025^WBC^L||8.7|X10E3/UL|3.4-10.8||||F|||20161203|SO   ^^L`)
)

// TestLoadMessage tests if the provided message can be loaded succesful
func TestLoadMessage(t *testing.T) {
	var message = []byte(
		`PV1|1|INPATIENT|WG|||||||||||||||||||||||||||||||||||||||||20171031041000
		DG1|1|I10|Z34.80^Encounter for supervision of other normal pregnancy, unspecified trimester^I10|Encounter for supervision of other normal pregnancy, unspecified trimester||^10150;EPT`)
	segments := hl7.LoadMessage(message)

	if len(segments) != 2 {
		t.Error(len(segments))
		t.Error(string(segments[0]))
		t.Error(string(segments[1]))
		t.Error(string(segments[2]))
		t.Error("[E0001] Cannot load message")
	}
}
func TestDecode(t *testing.T) {
	// Extract messages from file. Each message has several line
	Hl7Message, err := hl7.Decode(message)
	if err != nil {
		t.Error(err)
	}
	// Count the number of segments
	if len(Hl7Message) != 7 {
		t.Error("[E0003] Wrong number of segments")
	}
	// Find the patient name:
	// first name: PID segment, 5th repeated, component index 0, subcomponent index 1, field index 0
	// last name: PID segment, 5th repeated, component index 1, subcomponent index 1, field index 0
	var lastName hl7.Hl7Field
	var firstName hl7.Hl7Field
	for _, segment := range Hl7Message {
		// Find the segment name
		segmentName := segment.AtIndex("0.0.0.0")
		if segmentName == "PID" {
			lastName = segment[5][0][0][0]
			firstName = segment[5][0][1][0]
		}
	}
	if string(lastName) != "Truong" && string(firstName) != "Nicholas" {
		t.Error("[E0004] Wrong parsing names")
		t.Error(string(lastName))
	}
}

package main

import (
	"fmt"

	hl7 "github.com/zquangu112z/hl7-parser"
)

var (
	message = []byte(`MSH|^~\&|LCAMC|LABCORP|FAKE ACO|FAKE HOSPITAL|20170418093130||ORU^R01|M17108000000000001|P|2.3|||ER|ER
	PID|1|11111111111|11111111111|11111111111|Truong^Nicholas||19510610|M|||1511 MONTE VISTA ST^^PASADENA^CA^91106|||||||6337494512380
	PV1|1|O||||||1598879918^Khanh Hoang^Minh^^^^^^^^^^NPI|||||||||||||||||||||||||||||||SO
	ORC||633749453380|633749453380||||||20161202|||1598879918^Fake name^Fake given name^^^^^^^^^^NPI
	OBR|1||633749453380|005009^CBC WITH DIFFERENTIAL/PLATELET^L|||20161202|20161202||||||||1598879918^Fake name^Fake given name^^^^^^^^^^NPI||SO|||||||F
	OBX|1|ST|6690-2^LOINC^LN^005025^WBC^L||8.7|X10E3/UL|3.4-10.8||||F|||20161202|SO   ^^L
	`)
)

func main() {
	// Extract messages from file. Each message has several line
	Hl7Message, err := hl7.Unmarshal(message)
	if err != nil {
		panic(err)
	}
	// Count the number of segments
	fmt.Printf("There are totally %d segments.\n", len(Hl7Message))
	// Find the patient name:
	// first name: PID segment, 5th repeated, component index 0, subcomponent index 1, field index 0
	// last name: PID segment, 5th repeated, component index 1, subcomponent index 1, field index 0
	for _, segment := range Hl7Message {
		// Find the segment name
		segmentName := string(segment[0][0][0][0])
		if segmentName == "PID" {
			lastName := segment[5][0][0][0]
			firstName := segment[5][0][1][0]
			fmt.Printf("The patient is %s %s.\n", firstName, lastName)
		}
	}
}

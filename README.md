# HL7 parse by Golang
## About
The parser takes an Hl7 message passed and return a multi-dimensional array as:
* A *message* is a list of segments which are separated by line
* A *segment* is a list of *repeated* which are separated by the symbol `|`
* A *repeated* is a list of *component* which are separated by the symbol `~`
* A *component* is a list of *sub-component* which are separated by the symbol `^`
* A *sub-component* is a list of *field* which are separated by the symbol `&`

*Note:* In reality, the smallest unit of an HL7 message is *field* and an Hl7 segment are not required to have *repeated*, *componet*, or *sub-component*. Therefore, between 2 symbol `|` can be a *field*, a *componet*, or a *sub-component*. In these cases, the list will have length of 1.

## Implementation
```
type (
	Hl7Message      []Hl7Segment
	Hl7Segment      []Hl7Repeated
	Hl7Repeated     []Hl7Component
	Hl7Component    []Hl7SubComponent
	Hl7SubComponent []Hl7Field
	Hl7Field        []byte
)
```

## Usage
Assume obxSegment represents the observation segment:
```
OBX|1|ST|6690-2^LOINC^LN^005025^WBC^L||8.7|X10E3/UL|3.4-10.8||||F|||20161202|SO   ^^L
```

We can access the fields by the following way:
```
segmentName     := obxSegment[0][0][0][0]
testDate        := obxSegment[14][0][0][0]
testCode        := obxSegment[3][0][0][0]
testName	:= obxSegment[3][0][1][0]
testCodeSystemt	:= obxSegment[3][0][2][0]
testResult  	:= obxSegment[5][0][0][0]
```




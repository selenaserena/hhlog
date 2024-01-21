package main

import (
	"fmt"
)

type CodedChoice struct {
	name string
	code string
}

func askString(prompt string) string {
	answerString := ""
	fmt.Printf("Enter %s:\n", prompt)
	fmt.Scanf("%s\n", &answerString)
	return answerString
}

func askChoiceList(fieldName string, options []string) string {
	fmt.Printf("Enter %s:\n", fieldName)
	choiceNum := 0
	for i := 0; i < len(options); i++ {
		fmt.Printf("%d - %s\n", i, options[i])
	}
	fmt.Scanf("%d\n", &choiceNum)
	return options[choiceNum]
}

func askChoiceCode(fieldName string, options []CodedChoice) string {
	fmt.Printf("Enter %s:\n", fieldName)
	choiceNum := 0
	for i := 0; i < len(options); i++ {
		fmt.Printf("%d - %s\n", i, options[i].name)
	}

	fmt.Scanf("%d\n", &choiceNum)
	return options[choiceNum].code

}

func askCallsign() string {
	return askString("callsign")
}

func askEmail() string {
	return askString("email")
}

func askStationType() string {
	return askChoiceList("category-station", []string{"FIXED", "MOBILE", "SCHOOL"})
}

func askMode() string {
	return askChoiceList("category-mode", []string{"CW", "SSB", "MIXED"})
}

func askAssistedCat() string {
	return askChoiceList("category-assisted", []string{"ASSISTED", "NON-ASSISTED"})
}

func askLocationCode() string {
	return askChoiceCode("location", arrl_sections)
}

func askClass(classTypes []string) string {
	fmt.Printf("Enter class:\n")
	classNum := 0
	for i := 0; i < len(classTypes); i++ {
		fmt.Printf("%d - %s\n", i, classTypes[i])
	}
	fmt.Scanf("%d\n", &classNum)
	class := classTypes[classNum]
	return class
}

func askPower() string {
	return askChoiceList("category-power", []string{"HIGH", "LOW", "QRP"})
}

func askTransmitter() string {
	return askChoiceList("category-transmitter", []string{"ONE", "TWO"})
}

func Aprl10MeterContest() {
	contestCode := "ARRL-10"
	email := askEmail()
	callSign := askCallsign()
	locationCode := askLocationCode()
	catOpAs := askAssistedCat()
	catMode := askMode()
	power := askPower()
	station := askStationType()
	transmitter := askTransmitter()

	fmt.Printf("START-OF-LOG: 3.0\n")
	fmt.Printf("CONTEST: %s\n", contestCode)
	fmt.Printf("CALLSIGN: %s\n", callSign)
	fmt.Printf("LOCATION: %s\n", locationCode)
	fmt.Printf("CATEGORY-ASSISTED: %s\n", catOpAs)
	fmt.Printf("CATEGORY-BAND: 10M\n")
	fmt.Printf("CATEGORY-MODE: %s\n", catMode)
	fmt.Printf("CATEGORY-POWER: %s\n", power)
	fmt.Printf("CATEGORY-STATION: %s\n", station)
	fmt.Printf("CATEGORY-TRANSMITTER: %s\n", transmitter)
	fmt.Printf("EMAIL: %s\n", email)
	//fmt.Printf("END-OF-LOG:\n")

}

func NovSweepCW() {
	contestCode := "ARRL-SS-CW"
	callSign := askCallsign()
	locationCode := askLocationCode()
	catOp := askClass([]string{"SINGLE-OP", "MULTI-OP", "CHECKLOG"})

	fmt.Printf("START-OF-LOG: 3.0\n")
	fmt.Printf("CONTEST: %s\n", contestCode)
	fmt.Printf("CALLSIGN: %s\n", callSign)
	fmt.Printf("LOCATION: %s\n", locationCode)
	fmt.Printf("CATEGORY-OPERATOR: %s\n", catOp)
	//fmt.Printf("END-OF-LOG:\n")

}

func NorthAmQSOCW() {
	contestCode := "NAQP-CW"
	callSign := askCallsign()
	class := askClass([]string{"SINGLE-OP", "MULTI-TWO", "CHECKLOG"})
	power := askPower()

	fmt.Printf("START-OF-LOG: 3.0\n")
	fmt.Printf("CONTEST: %s\n", contestCode)
	fmt.Printf("CALLSIGN: %s\n", callSign)
	fmt.Printf("CATEGORY: %s ALL %s\n", class, power)
	//fmt.Printf("END-OF-LOG:\n")

}

const (
	ARRL10      = "ARRL-10"
	NAQP_CW     = "NAQP-CW"
	ARRL_SS_CW  = "ARRL-SS-CW"
	ARRL_SS_SSB = "ARRL-SS-SSB"
)

func PrintCabrilloHeader() {

	contestCode := askChoiceCode(
		"contest number",
		[]CodedChoice{
			{"ARRL 10-meter contest", ARRL10},
			{"North American QSO Party - CW", NAQP_CW},
			{"November sweepstakes CW", ARRL_SS_CW},
			{"November sweepstakes SSB", ARRL_SS_SSB},
		},
	)

	switch contestCode {
	case ARRL10:
		Aprl10MeterContest()
	case NAQP_CW:
		NorthAmQSOCW()
	case ARRL_SS_CW:
		NovSweepCW()
	}

}

var arrl_sections = []CodedChoice{
	{"Colorado", "CO"},
	{"Iowa", "IA"},
	{"Kansas", "KS"},
	{"Minnesota", "MN"},
	{"Missouri", "MO"},
	{"North Dakota", "ND"},
	{"Nebraska", "NE"},
	{"South Dakota", "SD"},
	{"Connecticut", "CT"},
	{"Eastern Massachusetts", "EMA"},
	{"Maine", "ME"},
	{"New Hampshire", "NH"},
	{"Rhode Island", "RI"},
	{"Vermont", "VT"},
	{"Western Massachusetts", "WMA"},
	{"Eastern New York", "ENY"},
	{"New York City-Long Island", "NLI"},
	{"Northern New Jersey", "NNJ"},
	{"Northern New York", "NNY"},
	{"Southern New Jersey", "SNJ"},
	{"Western New York", "WNY"},
	{"Delaware", "DE"},
	{"Eastern Pennsylvania", "EPA"},
	{"Maryland-DC", "MDC"},
	{"Western Pennsylvania", "WPA"},
	{"Alabama", "AL"},
	{"Georgia", "GA"},
	{"Kentucky", "KY"},
	{"North Carolina", "NC"},
	{"Northern Florida", "NFL"},
	{"Puerto Rico", "PR"},
	{"South Carolina", "SC"},
	{"Southern Florida", "SFL"},
	{"Tennessee", "TN"},
	{"Virginia", "VA"},
	{"US Virgin Islands", "VI"},
	{"West Central Florida", "WCF"},
	{"Arkansas", "AR"},
	{"Louisiana", "LA"},
	{"Mississippi", "MS"},
	{"New Mexico", "NM"},
	{"North Texas", "NTX"},
	{"Oklahoma", "OK"},
	{"South Texas", "STX"},
	{"West Texas", "WTX"},
	{"East Bay", "EB"},
	{"Los Angeles", "LAX"},
	{"Orange", "ORG"},
	{"Pacific", "PAC"},
	{"Santa Barbara", "SB"},
	{"Santa Clara Valley", "SCV"},
	{"San Diego", "SDG"},
	{"San Francisco", "SF"},
	{"San Joaquin Valley", "SJV"},
	{"Sacramento Valley", "SV"},
	{"Alaska", "AK"},
	{"Arizona", "AZ"},
	{"Eastern Washington", "EWA"},
	{"Idaho", "ID"},
	{"Montana", "MT"},
	{"Nevada", "NV"},
	{"Oregon", "OR"},
	{"Utah", "UT"},
	{"Western Washington", "WWA"},
	{"Wyoming", "WY"},
	{"Michigan", "MI"},
	{"Ohio", "OH"},
	{"West Virginia", "WV"},
	{"Illinois", "IL"},
	{"Indiana", "IN"},
	{"Wisconsin", "WI"},
	{"Alberta", "AB"},
	{"British Columbia", "BC"},
	{"Ontario Golden Horseshoe", "GH"},
	{"Manitoba", "MB"},
	{"New Brunswick", "NB"},
	{"Newfoundland and Labrador", "NL"},
	{"Nova Scotia", "NS"},
	{"Ontario East", "ONE"},
	{"Ontario North", "ONN"},
	{"Ontario South", "ONS"},
	{"Prince Edward Island", "PE"},
	{"Quebec", "QC"},
	{"Saskatchewan", "SK"},
	{"Other", "DX"},
}

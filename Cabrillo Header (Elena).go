package main

import (
	"fmt"
)

type CodedChoice struct {
	name string
	code string
}

func askCallsign() string {
	callSign := ""
	fmt.Printf("Enter call sign:\n")
	fmt.Scanf("%s\n", &callSign)
	return callSign
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
	return askChoiceCode("location", []CodedChoice{{"East Bay", "EB"}, {"Los Angeles", "LAX"}, {"Orange", "ORG"}, {"Santa Barbara", "SB"}, {"Santa Clara Valley", "SCV"}, {"San Diego", "SDG"}, {"San Francisco", "SF"}, {"San Joaquin Valley", "SJV"}, {"Sacramento Valley", "SV"}, {"Pacific", "PAC"}})
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

func Aprl10MeterContest() {
	contestCode := "ARRL-10"
	callSign := askCallsign()
	locationCode := askLocationCode()
	catOpAs := askAssistedCat()
	catMode := askMode()
	power := askPower()
	station := askStationType()

	fmt.Printf("START-OF-LOG: 3.0\n")
	fmt.Printf("CONTEST: %s\n", contestCode)
	fmt.Printf("CALLSIGN: %s\n", callSign)
	fmt.Printf("LOCATION: %s\n", locationCode)
	fmt.Printf("CATEGORY-ASSISTED: %s\n", catOpAs)
	fmt.Printf("CATEGORY-BAND: 10M\n")
	fmt.Printf("CATEGORY-MODE: %s\n", catMode)
	fmt.Printf("CATEGORY-POWER: %s\n", power)
	fmt.Printf("CATEGORY-STATION: %s\n", station)
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

	if contestCode == ARRL10 {
		Aprl10MeterContest()
	}

	if contestCode == NAQP_CW {
		NorthAmQSOCW()
	}
	if contestCode == ARRL_SS_CW {
		NovSweepCW()
	}
}

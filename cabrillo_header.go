package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type CodedChoice struct {
	name string
	code string
}

type ContestChoice struct {
	name string
	code string
	file string
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

func askContestChoice(options []ContestChoice) string {
	fmt.Printf("Enter contest number:\n")
	choiceNum := 0
	for i := 0; i < len(options); i++ {
		fmt.Printf("%d - %s\n", i, options[i].name)
	}

	fmt.Scanf("%d\n", &choiceNum)
	return options[choiceNum].file
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

func askCatOp() string {
	return askClass([]string{"SINGLE-OP", "MULTI-OP", "CHECKLOG"})
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

func askOverlay() string {
	return askChoiceList("category-overlay", []string{"NONE", "LIMITED-ANTENNAS"})
}

func askClubName() string {
	return askString("club name")
}

func askFullName() string {
	return askString("entrant's full name")
}

func askStreetAddress() string {
	return askString("street address")
}

func askCityAddress() string {
	return askString("city")
}

func askStateAddress() string {
	return askString("state-province")
}

func askZipAddress() string {
	return askString("postal code")
}

func askCountryAddress() string {
	return askString("country")
}

func Arrl10MeterContest() {
	contestCode := "ARRL-10"
	email := askEmail()
	callSign := askCallsign()
	locationCode := askLocationCode()
	catOpAs := askAssistedCat()
	catMode := askMode()
	power := askPower()
	station := askStationType()
	transmitter := askTransmitter()
	overlay := askOverlay()
	clubName := askClubName()
	fullName := askFullName()
	streetAddress := askStreetAddress()
	cityAddress := askCityAddress()
	stateAddress := askStateAddress()
	zipAddress := askZipAddress()
	countryAddress := askCountryAddress()

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
	fmt.Printf("CATEGORY-OVERLAY: %s\n", overlay)
	fmt.Printf("CLUB: %s\n", clubName)
	fmt.Printf("NAME: %s\n", fullName)
	fmt.Printf("ADDRESS: %s\n", streetAddress)
	fmt.Printf("ADDRESS-CITY: %s\n", cityAddress)
	fmt.Printf("ADDRESS-STATE-PROVINCE: %s\n", stateAddress)
	fmt.Printf("ADDRESS-POSTALCODE: %s\n", zipAddress)
	fmt.Printf("ADDRESS-COUNTRY: %s\n", countryAddress)
	fmt.Printf("EMAIL: %s\n", email)
	//fmt.Printf("END-OF-LOG:\n")

}

func NovSweepCW() {
	contestCode := "ARRL-SS-CW"
	callSign := askCallsign()
	locationCode := askLocationCode()
	catOp := askClass([]string{"SINGLE-OP", "MULTI-OP", "CHECKLOG"})
	catOpAs := askAssistedCat()
	power := askPower()
	clubName := askClubName()
	fullName := askFullName()
	streetAddress := askStreetAddress()
	cityAddress := askCityAddress()
	stateAddress := askStateAddress()
	zipAddress := askZipAddress()
	countryAddress := askCountryAddress()
	email := askEmail()

	fmt.Printf("START-OF-LOG: 3.0\n")
	fmt.Printf("CONTEST: %s\n", contestCode)
	fmt.Printf("LOCATION: %s\n", locationCode)
	fmt.Printf("CALLSIGN: %s\n", callSign)
	fmt.Printf("CATEGORY-ASSISTED: %s\n", catOpAs)
	fmt.Printf("CATEGORY-BAND: ALL\n")
	fmt.Printf("CATEGORY-MODE: CW\n")
	fmt.Printf("CATEGORY-POWER: %s\n", power)
	fmt.Printf("CATEGORY-OPERATOR: %s\n", catOp)
	fmt.Printf("CLUB: %s\n", clubName)
	fmt.Printf("NAME: %s\n", fullName)
	fmt.Printf("ADDRESS: %s\n", streetAddress)
	fmt.Printf("ADDRESS-CITY: %s\n", cityAddress)
	fmt.Printf("ADDRESS-STATE-PROVINCE: %s\n", stateAddress)
	fmt.Printf("ADDRESS-POSTALCODE: %s\n", zipAddress)
	fmt.Printf("ADDRESS-COUNTRY: %s\n", countryAddress)
	fmt.Printf("EMAIL: %s\n", email)
	//fmt.Printf("END-OF-LOG:\n")

}

func NorthAmQSOCW() {
	contestCode := "NAQP-CW"
	callSign := askCallsign()
	class := askClass([]string{"SINGLE-OP", "MULTI-TWO", "CHECKLOG"})
	power := askPower()
	fullName := askFullName()
	streetAddress := askStreetAddress()
	cityAddress := askCityAddress()
	stateAddress := askStateAddress()
	zipAddress := askZipAddress()
	countryAddress := askCountryAddress()
	email := askEmail()

	fmt.Printf("START-OF-LOG: 3.0\n")
	fmt.Printf("CONTEST: %s\n", contestCode)
	fmt.Printf("CALLSIGN: %s\n", callSign)
	fmt.Printf("CATEGORY: %s ALL %s\n", class, power)
	fmt.Printf("NAME: %s\n", fullName)
	fmt.Printf("ADDRESS: %s\n", streetAddress)
	fmt.Printf("ADDRESS-CITY: %s\n", cityAddress)
	fmt.Printf("ADDRESS-STATE-PROVINCE: %s\n", stateAddress)
	fmt.Printf("ADDRESS-POSTALCODE: %s\n", zipAddress)
	fmt.Printf("ADDRESS-COUNTRY: %s\n", countryAddress)
	fmt.Printf("EMAIL: %s\n", email)
	//fmt.Printf("END-OF-LOG:\n")

}

const (
	ARRL10      = "ARRL-10"
	NAQP_CW     = "NAQP-CW"
	ARRL_SS_CW  = "ARRL-SS-CW"
	ARRL_SS_SSB = "ARRL-SS-SSB"
)

func PrintCabrilloHeader(outF *os.File) {
	fileName := askContestChoice(parseContestChoices())
	parseContestFile(fileName, outF)

}

func parseContestFile(fileName string, outF *os.File) {
	f, err := os.Open("contest_specs/" + fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open %s: %s", fileName, err)
		os.Exit(1)
	}
	fileReader := bufio.NewReader(f)
	//fmt.Printf("--------------\n")
	readLine(fileReader)
	fmt.Fprintf(outF, "START-OF-LOG: 3.0\n")
	code, _ := readLine(fileReader)
	fmt.Fprintf(outF, "CONTEST: %s\n", code)
	for {
		fieldName, err := readLine(fileReader)
		if err == io.EOF {
			break
		}
		switch fieldName {
		case "CALLSIGN":
			callSign := askCallsign()
			fmt.Fprintf(outF, "CALLSIGN: %s\n", callSign)
		case "LOCATION":
			location := askLocationCode()
			fmt.Fprintf(outF, "LOCATION: %s\n", location)
		case "CATEGORY-ASSISTED":
			catOpAs := askAssistedCat()
			fmt.Fprintf(outF, "CATEGORY-ASSISTED: %s\n", catOpAs)
		case "CATEGORY-MODE":
			mode := askMode()
			fmt.Fprintf(outF, "CATEGORY-MODE: %s\n", mode)
		case "CATEGORY-POWER":
			power := askPower()
			fmt.Fprintf(outF, "CATEGORY-POWER: %s\n", power)
		case "CATEGORY-OPERATOR":
			catOp := askCatOp()
			fmt.Fprintf(outF, "CATEGORY-OPERATOR: %s\n", catOp)
		case "CATEGORY-STATION":
			station := askStationType()
			fmt.Fprintf(outF, "CATEGORY-STATION: %s\n", station)

		}

		//fmt.Printf("%s\n", fieldName)

	}
	//fmt.Printf("--------------\n")
}

func parseContestChoices() (choices []ContestChoice) {
	f, err := os.Open("contest_specs/index")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open index file: %s", err)
		os.Exit(1)
	}
	fileReader := bufio.NewReader(f)
	for {
		contestName, err := readLine(fileReader)
		if err == io.EOF {
			break
		}

		contestCode, err := readLine(fileReader)

		fileName, _ := readLine(fileReader)

		choices = append(choices, ContestChoice{contestName, contestCode, fileName})
	}
	return choices
}

func readLine(fileReader *bufio.Reader) (string, error) {
	line, err := fileReader.ReadString('\n')
	if err == io.EOF {
		return "", err
	}
	colonPos := strings.Index(line, ":")
	return strings.TrimSpace(line[colonPos+1:]), nil
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

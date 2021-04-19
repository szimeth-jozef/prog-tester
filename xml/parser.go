package xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type TestCases struct {
	XMLName xml.Name   `xml:"tester"`
	All     []TestCase `xml:"testcase"`
}

type TestCase struct {
	// XMLName  xml.Name `xml:"testcase"`
	Args     string `xml:"args"`
	Input    string `xml:"input"`
	Expect   string `xml:"expect"`
	ExitCode int    `xml:"exitcode"`
}

func (t *TestCase) equals(testCase TestCase) bool {
	return (t.Args == testCase.Args &&
		t.Input == testCase.Input &&
		t.Expect == testCase.Expect &&
		t.ExitCode == testCase.ExitCode)
}

func ParseXmlSchema(name string) []TestCase {
	xmlFile, err := os.Open(name)
	if err != nil {
		log.Fatalf("[ERROR]: Couldn't open %s. Make sure that the name is corrent.\n", name)
	}
	defer xmlFile.Close()

	xmlFileByteValue, _ := ioutil.ReadAll(xmlFile)

	testCases, err := getTestCasesFromXml(xmlFileByteValue)
	if err != nil {
		fmt.Printf("[ERROR]: %s.\n", err)
		os.Exit(1)
	}

	return testCases
}

func getTestCasesFromXml(xmlBytes []byte) ([]TestCase, error) {
	var testCases TestCases
	err := xml.Unmarshal(xmlBytes, &testCases)
	if err != nil {
		return nil, err
	}
	return filterInvalidCases(testCases.All), nil
}

func filterInvalidCases(testCases []TestCase) []TestCase {
	var validTests []TestCase
	for i, t := range testCases {
		if len(t.Input) == 0 {
			log.Fatalf("[ERROR]: TestCase %d has empty input, to run the test remove or fill it in.", i+1)
		} else if len(t.Expect) == 0 {
			log.Fatalf("[ERROR]: TestCase %d has empty expect, to run the test remove or fill it in.", i+1)
		} else {
			validTests = append(validTests, t)
		}
	}
	return validTests
}

package crawler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// fetchMajor downloads a major's website and returns a populated major
func fetchMajor(id int) (*major, error) {
	major := &major{}
	major.externalID = id
	document, err := getDocument(major.getURL())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch major %d; %v", id, err)
	}
	rawName := document.Find("h1").First().Text()
	major.name = strings.Title(strings.TrimSpace(strings.Split(rawName, "(")[0]))
	return major, nil
}

// fetchProgram downloads a major's website and returns a populated program
func fetchProgram(id int) (*program, error) {
	p := &program{externalID: id}
	document, err := getDocument(p.getURL())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch program %d; %v", id, err)
	}
	parenthesisRegex := regexp.MustCompile(`\(([^\)]+)\)`)
	digitRegex := regexp.MustCompile(`\d+`)
	programRawName := document.Find("h1").First().Text()
	programRawName = parenthesisRegex.FindString(programRawName)
	p.name = strings.Title(programRawName[1 : len(programRawName)-1])
	p.year, err = strconv.Atoi(digitRegex.FindString(p.name))
	if err != nil {
		return nil, fmt.Errorf("failed parsing program year for program %d; %v", id, err)
	}
	return p, nil
}

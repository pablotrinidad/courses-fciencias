package crawler

import (
	"fmt"
	"strings"
)

// fetchMajor download a major's website and parse the input
func fetchMajor(id int) (*major, error) {
	major := &major{}
	major.externalID = id
	document, err := getDocument(major.getURL())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch major; %v", err)
	}
	rawName := document.Find("h1").First().Text()
	major.name = strings.Title(strings.TrimSpace(strings.Split(rawName, "(")[0]))
	return major, nil
}

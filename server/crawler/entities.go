package crawler

import (
	"fmt"
	"strconv"
)

type major struct {
	externalID int
	name       string
}

// getURL returns the major's public URL
func (m *major) getURL() string {
	return fmt.Sprintf("%s/%s/%d", baseURL, majorsURL, m.externalID)
}

// toProto returns the protobuf representation of the major
func (m *major) toProto() *Major {
	return &Major{
		Id:   strconv.Itoa(m.externalID),
		Name: m.name,
		Url:  m.getURL(),
	}
}

type program struct {
	externalID int
	major      int
	name       string
	year       int
}

// getURL of the program's main page
func (p *program) getURL() string {
	return fmt.Sprintf("%s/%s/%d/%d", baseURL, programsURL, p.major, p.externalID)
}

// toProto returns the protobuf representation of the program
func (p *program) toProto() *Program {
	return &Program{
		Id:   strconv.Itoa(p.externalID),
		Name: p.name,
		Year: uint32(p.year),
		Url:  p.getURL(),
	}
}

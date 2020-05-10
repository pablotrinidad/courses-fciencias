package crawler

import (
	"fmt"
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
		Id:   uint32(m.externalID),
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
		Id:   uint32(p.externalID),
		Name: p.name,
		Year: uint32(p.year),
		Url:  p.getURL(),
	}
}

type course struct {
	externalID  int
	program     int
	name        string
	semester    uint
	mandatory   bool
	credits     uint
	syllabusURL string
}

// getURL of the course's main page
func (c *course) getURL() string {
	return fmt.Sprintf("%s/%s/%d/%d", baseURL, coursesURL, c.program, c.externalID)
}

// toProto returns the protobuf representation of the course
func (c *course) toProto() *Course {
	return &Course{
		Id:        uint32(c.externalID),
		Name:      c.name,
		Semester:  uint32(c.semester),
		Mandatory: c.mandatory,
		Credits:   uint32(c.credits),
		Syllabus:  c.syllabusURL,
		Url:       c.getURL(),
	}
}

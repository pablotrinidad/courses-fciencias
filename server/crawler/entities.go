package crawler

import "fmt"

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
		Id:   string(m.externalID),
		Name: m.name,
		Url:  m.getURL(),
	}
}

package main

import (
	"fmt"
	"strings"
)

// parseMethod receives a gRPC full method name and returns the
// service name and the method name. Returns error if it cannot
// parse the given full name.
func parseMethod(fullName string) (string, string, error) {
	sections := strings.SplitN(fullName, "/", 3)
	if len(sections) != 3 {
		return "", "", fmt.Errorf("split by / returned %d, want %d", len(sections), 3)
	}
	serviceSections := strings.SplitN(sections[1], ".", 2)
	if len(serviceSections) != 2 {
		return "", "", fmt.Errorf("split by . returned %d, want d", len(serviceSections), 2)
	}
	return serviceSections[1], sections[2], nil
}

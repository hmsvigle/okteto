package model

import (
	"fmt"
	"regexp"
)

var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9][-a-zA-Z0-9]*$`).MatchString

//Space represents a dev space
type Space struct {
	ID      string   `json:"id,omitempty" yaml:"id,omitempty"`
	Name    string   `json:"name,omitempty" yaml:"name,omitempty"`
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`
}

func (s *Space) validate() error {
	if s.ID == "" {
		return fmt.Errorf("Name cannot be empty")
	}

	if !isAlphaNumeric(s.ID) {
		return fmt.Errorf("Name must be alphanumeric")
	}

	return nil
}

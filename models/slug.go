package models

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

type Slug struct {
	value string
}

func NewSlug(value string) *Slug {
	slugValue := parseSlug(value)
	return &Slug{value: slugValue}
}

func (s *Slug) Cast(optionsAlreadyAvailable []string) {
	slugsSet := hashset.New()
	for _, value := range optionsAlreadyAvailable {
		slugsSet.Add(value)
	}

	currentString := s.value
	index := 0
	for slugsSet.Contains(currentString) {
		index += 1
		currentString = fmt.Sprintf("%s-%d", s.value, index)
	}

	s.value = currentString
}

func (s *Slug) String() string {
	return s.value
}

func parseSlug(value string) string {
	var nonSlugRegex = regexp.MustCompile(`[^a-z0-9-]`)
	var duplicatedDash = regexp.MustCompile(`[-]+`)

	v := value
	v = strings.ToLower(v)
	v = strings.TrimSpace(v)
	v = nonSlugRegex.ReplaceAllLiteralString(v, "-")
	v = duplicatedDash.ReplaceAllLiteralString(v, "-")

	return v
}

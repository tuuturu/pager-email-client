package filtering

import (
	"fmt"
	"regexp"

	"github.com/tuuturu/pager-email-client/pkg/core"
)

func newSubjectFilter(pattern string) (core.Filter, error) {
	matcher, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("compiling pattern: %w", err)
	}

	return &subjectFilter{matcher: matcher}, nil
}

type subjectFilter struct {
	matcher *regexp.Regexp
}

func (s subjectFilter) Test(email core.Email) bool {
	return s.matcher.Match([]byte(email.Subject))
}

func newSenderFilter(pattern string) (core.Filter, error) {
	matcher, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("compiling pattern: %w", err)
	}

	return &fromFilter{matcher: matcher}, nil
}

type fromFilter struct {
	matcher *regexp.Regexp
}

func (f fromFilter) Test(email core.Email) bool {
	return f.matcher.Match([]byte(email.From))
}

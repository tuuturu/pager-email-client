package filtering

import (
	"regexp"
	"testing"

	"github.com/tuuturu/pager-email-client/pkg/core"
	"gotest.tools/assert"
)

func TestSubjectFilter(t *testing.T) {
	testCases := []struct {
		name string

		withEmail   core.Email
		withMatcher *regexp.Regexp

		expect bool
	}{
		{
			name: "Should match with correct subject filter",

			withEmail:   core.Email{Subject: "Donate something please"},
			withMatcher: regexp.MustCompile("[Dd]onate"),

			expect: true,
		},
		{
			name: "Should fail with incorrect subject filter",

			withEmail:   core.Email{Subject: "Donate something please"},
			withMatcher: regexp.MustCompile("[Ww]ash"),

			expect: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			filter := subjectFilter{matcher: tc.withMatcher}

			assert.Equal(t, tc.expect, filter.Test(tc.withEmail))
		})
	}
}

func TestFromFilter(t *testing.T) {
	testCases := []struct {
		name string

		withEmail   core.Email
		withMatcher *regexp.Regexp

		expect bool
	}{
		{
			name: "Should match with correct subject filter",

			withEmail:   core.Email{From: "tease@hooters.com"},
			withMatcher: regexp.MustCompile(`.+@hooters.com`),

			expect: true,
		},
		{
			name: "Should fail with incorrect subject filter",

			withEmail:   core.Email{Subject: "tease@hooters.com"},
			withMatcher: regexp.MustCompile(".+@example.com"),

			expect: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			filter := fromFilter{matcher: tc.withMatcher}

			assert.Equal(t, tc.expect, filter.Test(tc.withEmail))
		})
	}
}

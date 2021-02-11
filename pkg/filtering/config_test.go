package filtering

import (
	"testing"

	"github.com/tuuturu/pager-email-client/pkg/core"
	"gotest.tools/assert"
)

func TestFilterItem(t *testing.T) {
	testCases := []struct {
		name string

		withFilterItem filterItem
		withEmail      core.Email

		expectTestResult bool
	}{
		{
			name: "Should test true with matching subject pattern",

			withFilterItem: filterItem{
				SubjectPattern: "[dD]onate",
			},
			withEmail: core.Email{
				Subject: "Would you want to donate some cash?",
			},

			expectTestResult: true,
		},
		{
			name: "Should test true with matching from pattern",

			withFilterItem: filterItem{
				FromPattern: ".*@soulless.io",
			},
			withEmail: core.Email{
				From: "johnny@soulless.io",
			},

			expectTestResult: true,
		},
		{
			name: "Should test false with no pattern",

			withFilterItem: filterItem{},
			withEmail: core.Email{
				From:    "johnny@soulless.io",
				Subject: "Donate please",
			},

			expectTestResult: false,
		},
		{
			name: "Should test false with no matches",

			withFilterItem: filterItem{
				FromPattern:    ".*@example.com",
				SubjectPattern: "[wW]ish",
			},
			withEmail: core.Email{
				From:    "johnny@soulless.io",
				Subject: "Donate please",
			},

			expectTestResult: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			filter, err := tc.withFilterItem.Filter()
			assert.NilError(t, err)

			assert.Equal(t, tc.expectTestResult, filter.Test(tc.withEmail))
		})
	}
}

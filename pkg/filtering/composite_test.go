package filtering

import (
	"testing"

	"github.com/tuuturu/pager-email-client/pkg/core"
	"gotest.tools/assert"
)

type trueFilter struct{}

func (t trueFilter) Test(_ core.Email) bool { return true }

type falseFilter struct{}

func (f falseFilter) Test(_ core.Email) bool { return false }

func TestCompositeFilter(t *testing.T) {
	testCases := []struct {
		name string

		withType    compositeFilterType
		withFilters []core.Filter

		expectTestResult bool
	}{
		{
			name: "Should test true with AND type and only trues",

			withType:    compositeFilterTypeAND,
			withFilters: []core.Filter{trueFilter{}, trueFilter{}},

			expectTestResult: true,
		},
		{
			name: "Should test false with AND type and one false",

			withType:    compositeFilterTypeAND,
			withFilters: []core.Filter{trueFilter{}, falseFilter{}, trueFilter{}},

			expectTestResult: false,
		},
		{
			name: "Should test true with OR type and one true",

			withType:    compositeFilterTypeOR,
			withFilters: []core.Filter{falseFilter{}, trueFilter{}, falseFilter{}},

			expectTestResult: true,
		},
		{
			name: "Should test false with OR type and all false",

			withType:    compositeFilterTypeOR,
			withFilters: []core.Filter{falseFilter{}, falseFilter{}, falseFilter{}},

			expectTestResult: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			filter := compositeFilter{
				filterType: tc.withType,
				filters:    tc.withFilters,
			}

			assert.Equal(t, tc.expectTestResult, filter.Test(core.Email{}))
		})
	}
}

package filtering

import (
	"github.com/tuuturu/pager-email-client/pkg/core"
)

type compositeFilterType int

const (
	compositeFilterTypeAND compositeFilterType = iota
	compositeFilterTypeOR
)

type compositeFilter struct {
	filterType compositeFilterType
	filters    []core.Filter
}

func (c compositeFilter) andTest(email core.Email) bool {
	if len(c.filters) == 0 {
		return false
	}

	for _, f := range c.filters {
		if !f.Test(email) {
			return false
		}
	}

	return true
}

func (c compositeFilter) orTest(email core.Email) bool {
	for _, f := range c.filters {
		if f.Test(email) {
			return true
		}
	}

	return false
}

func (c compositeFilter) Test(email core.Email) bool {
	if c.filterType == compositeFilterTypeAND {
		return c.andTest(email)
	}

	if c.filterType == compositeFilterTypeOR {
		return c.orTest(email)
	}

	return false
}

package filtering

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/tuuturu/pager-email-client/pkg/core"
	"gopkg.in/yaml.v3"
)

func ParseFilterConfig(reader io.Reader) (core.Filter, error) {
	filterConfig := filterConfig{}

	result, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	err = yaml.Unmarshal(result, &filterConfig)
	if err != nil {
		return nil, fmt.Errorf("parsing filter config: %w", err)
	}

	return filterConfig.Filter()
}

type filterConfig struct {
	On []filterItem `json:"on" yaml:"on"`
}

func (c filterConfig) Filter() (result core.Filter, err error) {
	filters := make([]core.Filter, len(c.On))

	for index, filter := range c.On {
		filters[index], err = filter.Filter()

		if err != nil {
			return nil, err
		}
	}

	return compositeFilter{filterType: compositeFilterTypeOR, filters: filters}, nil
}

type filterItem struct {
	SubjectPattern string `json:"subject" yaml:"subject"`
	FromPattern    string `json:"from" yaml:"from"`
}

func (f filterItem) Filter() (result core.Filter, err error) {
	filters := make([]core.Filter, 0)

	if f.FromPattern != "" {
		potentialFilter, err := newSenderFilter(f.FromPattern)
		if err != nil {
			return nil, fmt.Errorf("parsing sender filter: %w", err)
		}

		filters = append(filters, potentialFilter)
	}

	if f.SubjectPattern != "" {
		potentialFilter, err := newSubjectFilter(f.SubjectPattern)
		if err != nil {
			return nil, fmt.Errorf("parsing subject filter: %w", err)
		}

		filters = append(filters, potentialFilter)
	}

	return compositeFilter{filterType: compositeFilterTypeAND, filters: filters}, nil
}

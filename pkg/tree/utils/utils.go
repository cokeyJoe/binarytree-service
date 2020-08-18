package utils

import (
	"binarytree/pkg/tree/binary"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

var (
	ErrIntsCountIsNotEnough = errors.New("ints count is not enough")
)

const (
	defaultCount = 30
)

type opts struct {
	minCount int
}

type Option func(*opts)

func WithMinCount(minCount int) Option {
	return func(option *opts) {
		option.minCount = minCount
	}
}

func FromReader(r io.Reader, options ...Option) (*binary.Tree, error) {

	ints, err := decodeJSON(r)
	if err != nil {
		return nil, err
	}

	defaultOpts := opts{
		minCount: defaultCount,
	}

	for _, options := range options {
		options(&defaultOpts)
	}

	if !isIntsEnough(ints, defaultOpts.minCount) {
		return nil, errors.Wrapf(ErrIntsCountIsNotEnough, "got %d, required %d", len(ints), defaultOpts.minCount)
	}

	tree := binary.New(ints...)

	return tree, nil

}

func isIntsEnough(ints []int, minCount int) bool {
	if ints == nil {
		return false
	}

	if len(ints) < minCount {
		return false
	}

	return true
}

func decodeJSON(reader io.Reader) ([]int, error) {
	var vals initValues

	if err := json.NewDecoder(reader).Decode(&vals); err != nil {
		return nil, err
	}

	return vals.Ints, nil
}

type initValues struct {
	Ints []int `json:"ints"`
}

package pipe_filter

import (
	"errors"
)

var SumFilterWrongFormatError = errors.New("unsupported field type found in sum filter")

type SumFilter struct {
	delimiter string
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]int) //检查数据类型
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, part := range parts {
		ret += part
	}
	return ret, nil
}

package pipe_filter

import (
	"errors"
	"strconv"
)

var IntegerFilterWrongFormatError = errors.New("unsupported field type found in Integer filter")

type IntegerFilter struct {
	delimiter string
}

func NewIntegerFilter() *IntegerFilter {
	return &IntegerFilter{}
}

func (sf *IntegerFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string) //检查数据类型
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	ret := []int{}
	for _, part := range parts {
		//convert str to int
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}

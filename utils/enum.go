package utils

import (
	"errors"
	"fmt"
)

type CommonEnum struct {
	Enums map[int]EnumItem
	Name  string
}

func NewEnum(name string) *CommonEnum {
	resp := &CommonEnum{
		make(map[int]EnumItem),
		name,
	}
	return resp
}

func (ep *CommonEnum) Set(code int, msg string) (EnumItem, error) {
	// TODO check error

	item := EnumItem{
		code, msg,
	}

	if _, ok := ep.Enums[code]; ok {
		return item, errors.New(fmt.Sprintf("Code: %d has already in enums", code))
	}

	ep.Enums[code] = item
	return item, nil
}

func (ep *CommonEnum) GetEnumsMsgs() (resp []string) {
	for _, ei := range ep.Enums {
		resp = append(resp, ei.Msg)
	}

	return resp
}

type EnumItem struct {
	Code int
	Msg  string
}

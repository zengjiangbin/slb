package core

import "github.com/zengjiangbin/driver/structs/service"

type slbCore struct{}

func (s *slbCore) Init() {}

func (s *slbCore) Start() {}

func NewService() service.Service {
	return new(slbCore)
}

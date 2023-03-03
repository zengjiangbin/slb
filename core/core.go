package core

import "github.com/zengjiangbin/driver/structs/service"

func NewService() service.Service {
	return new(slbCore)
}

type slbCore struct{}

func (s *slbCore) Init() {}

func (s *slbCore) Start() {}

func (s *slbCore) Name() service.ServiceName {
	return service.SlbServiceName
}

func (s *slbCore) Version() string {
	return version
}

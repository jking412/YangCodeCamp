package service

import (
	"sync"
)

type IClassService interface {
	//GetAllClasses() ([]*model.Class, error)
}

type ClassService struct {
}

var classServiceSync sync.Once
var _ IClassService = (*ClassService)(nil)

func NewClassService() IClassService {
	var s *ClassService
	classServiceSync.Do(func() {
		s = &ClassService{}
	})
	return s
}

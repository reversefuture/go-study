package main

import "github.com/pkg/errors"

// 编译期接口实现检查：确保 *study 类型确实实现了 Study 接口所有方法
var _ Study = (*study)(nil)

type Study interface {
	Listen(msg string) string
	Speak(msg string) string
	Read(msg string) string
	Write(msg string) string
}

type study struct {
	Name string
}

func (s *study) Listen(msg string) string {
	return s.Name + " 听 " + msg
}

func (s *study) Speak(msg string) string {
	return s.Name + " 说 " + msg
}

func (s *study) Read(msg string) string {
	return s.Name + " 读 " + msg
}

func (s *study) Write(msg string) string {
	return s.Name + " 写 " + msg
}

func NewStudy(name string) (Study, error) {
	if name == "" {
		return nil, errors.New("name required")
	}

	return &study{
		Name: name,
	}, nil
}

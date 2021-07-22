package ru

type ru struct{}

func (r ru) Code() string {
	return "ru"
}

func (r ru) Local() string {
	return "русский"
}

func (r ru) Name() string {
	return "Russian"
}

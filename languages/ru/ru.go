package ru

type Ru struct{}

func (r Ru) Code() string {
	return "ru"
}

func (r Ru) Local() string {
	return "русский"
}

func (r Ru) Name() string {
	return "Russian"
}

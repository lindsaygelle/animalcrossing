package otter

type Otter struct{}

func (o Otter) Id() string    { return "otter" }
func (o Otter) Special() bool { return true }

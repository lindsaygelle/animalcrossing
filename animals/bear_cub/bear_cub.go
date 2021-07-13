package bear_cub

type BearCub struct{}

func (b BearCub) Id() string    { return "bear_cub" }
func (b BearCub) Special() bool { return false }

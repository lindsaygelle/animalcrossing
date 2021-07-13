package ko

type Ko struct{}

func (k Ko) Code() string  { return "ko" }
func (k Ko) Local() string { return "한국인" }
func (k Ko) Name() string  { return "Korean" }

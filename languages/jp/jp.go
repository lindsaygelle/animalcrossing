package jp

type Jp struct{}

func (j Jp) Code() string  { return "jp" }
func (j Jp) Local() string { return "日本語" }
func (j Jp) Name() string  { return "Japanese" }

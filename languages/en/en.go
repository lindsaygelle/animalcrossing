package en

type En struct{}

func (e En) Code() string {
	return "en"
}

func (e En) Local() string {
	return "English"
}

func (e En) Name() string {
	return e.Local()
}

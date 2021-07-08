package translations

import "golang.org/x/text/language"

type ChineseTraditional struct{}

func (ct ChineseTraditional) Code() string {
	return language.TraditionalChinese.String()
}

func (ct ChineseTraditional) Name() string {
	return "Chinese Traditional"
}

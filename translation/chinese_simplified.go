package translation

import "golang.org/x/text/language"

type ChineseSimplified struct{}

func (cs ChineseSimplified) Code() string {
	return language.SimplifiedChinese.String()
}

func (cs ChineseSimplified) Name() string {
	return "Chinese Simplified"
}

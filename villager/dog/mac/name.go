package mac

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Mac"
	nameCanadianFrench       string = "Brutus"
	nameDutch                string = "Mac"
	nameFrench               string = "Brutus"
	nameGerman               string = "Wuffi"
	nameItalian              string = "Teo"
	nameJapanese             string = "チャンプ"
	nameLatinAmericanSpanish string = "Pit"
	nameKorean               string = "챔프"
	nameRussian              string = "Мак"
	nameSpanish              string = "Pit"
	nameSimplifiedChinese    string = "金牌"
	nameTraditionalChinese   string = "金牌"
)

var (
	name = map[language.Tag]string{
		language.AmericanEnglish:      nameAmericanEnglish,
		language.CanadianFrench:       nameCanadianFrench,
		language.Dutch:                nameDutch,
		language.French:               nameFrench,
		language.German:               nameGerman,
		language.Italian:              nameItalian,
		language.Japanese:             nameJapanese,
		language.Korean:               nameKorean,
		language.LatinAmericanSpanish: nameLatinAmericanSpanish,
		language.Russian:              nameRussian,
		language.Spanish:              nameSpanish,
		language.SimplifiedChinese:    nameSimplifiedChinese,
		language.TraditionalChinese:   nameTraditionalChinese}
)

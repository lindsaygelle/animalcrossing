package penelope

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Penelope"
	nameCanadianFrench       string = "Missy"
	nameDutch                string = "Penelope"
	nameFrench               string = "Missy"
	nameGerman               string = "Penelope"
	nameItalian              string = "Penelope"
	nameJapanese             string = "チューこ"
	nameLatinAmericanSpanish string = "Adela"
	nameKorean               string = "찍순이"
	nameRussian              string = "Адела"
	nameSpanish              string = "Adela"
	nameSimplifiedChinese    string = "小啾"
	nameTraditionalChinese   string = "小啾"
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

package coach

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Coach"
	nameCanadianFrench       string = "Arnold"
	nameDutch                string = "Coach"
	nameFrench               string = "Arnold"
	nameGerman               string = "Arnold"
	nameItalian              string = "Ercole"
	nameJapanese             string = "テッチャン"
	nameLatinAmericanSpanish string = "Cacho"
	nameKorean               string = "철소"
	nameRussian              string = "Коуч"
	nameSpanish              string = "Cacho"
	nameSimplifiedChinese    string = "大常"
	nameTraditionalChinese   string = "大常"
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

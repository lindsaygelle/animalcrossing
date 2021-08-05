package vladimir

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Vladimir"
	nameCanadianFrench       string = "Vladimir"
	nameDutch                string = "Vladimir"
	nameFrench               string = "Vladimir"
	nameGerman               string = "Vladimir"
	nameItalian              string = "Vladimir"
	nameJapanese             string = "ガビ"
	nameLatinAmericanSpanish string = "Vladimir"
	nameKorean               string = "곰비"
	nameRussian              string = "Владимир"
	nameSpanish              string = "Vladimir"
	nameSimplifiedChinese    string = "嘉弼"
	nameTraditionalChinese   string = "嘉弼"
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

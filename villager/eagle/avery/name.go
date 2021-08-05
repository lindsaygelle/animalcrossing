package avery

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Avery"
	nameCanadianFrench       string = "Faust"
	nameDutch                string = "Avery"
	nameFrench               string = "Faust"
	nameGerman               string = "Quetzal"
	nameItalian              string = "Falco"
	nameJapanese             string = "クスケチャ"
	nameLatinAmericanSpanish string = "Cuzco"
	nameKorean               string = "쿠스케처"
	nameRussian              string = "Авери"
	nameSpanish              string = "Cuzco"
	nameSimplifiedChinese    string = "安谷斯"
	nameTraditionalChinese   string = "安谷斯"
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

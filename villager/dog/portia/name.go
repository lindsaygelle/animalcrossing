package portia

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Portia"
	nameCanadianFrench       string = "Dalma"
	nameDutch                string = "Portia"
	nameFrench               string = "Dalma"
	nameGerman               string = "Isolde"
	nameItalian              string = "Porzia"
	nameJapanese             string = "ブレンダ"
	nameLatinAmericanSpanish string = "Amanda"
	nameKorean               string = "블랜더"
	nameRussian              string = "Порция"
	nameSpanish              string = "Amanda"
	nameSimplifiedChinese    string = "傅兰"
	nameTraditionalChinese   string = "傅蘭"
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

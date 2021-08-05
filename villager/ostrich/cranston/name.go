package cranston

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cranston"
	nameCanadianFrench       string = "Gabin"
	nameDutch                string = "Cranston"
	nameFrench               string = "Gabin"
	nameGerman               string = "Guido"
	nameItalian              string = "Carmine"
	nameJapanese             string = "トキオ"
	nameLatinAmericanSpanish string = "Carmelo"
	nameKorean               string = "타키"
	nameRussian              string = "Крэнстон"
	nameSpanish              string = "Carmelo"
	nameSimplifiedChinese    string = "朱禄"
	nameTraditionalChinese   string = "朱祿"
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

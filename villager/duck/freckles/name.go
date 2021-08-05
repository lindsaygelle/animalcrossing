package freckles

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Freckles"
	nameCanadianFrench       string = "Caro"
	nameDutch                string = "Freckles"
	nameFrench               string = "Caro"
	nameGerman               string = "Quacks"
	nameItalian              string = "Semola"
	nameJapanese             string = "マグロ"
	nameLatinAmericanSpanish string = "Rebeca"
	nameKorean               string = "다랑어"
	nameRussian              string = "Фреклс"
	nameSpanish              string = "Rebeca"
	nameSimplifiedChinese    string = "小鲔"
	nameTraditionalChinese   string = "小鮪"
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

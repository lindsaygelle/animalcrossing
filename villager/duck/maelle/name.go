package maelle

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Maelle"
	nameCanadianFrench       string = "Maëlle"
	nameDutch                string = "Maelle"
	nameFrench               string = "Maëlle"
	nameGerman               string = "Sissi"
	nameItalian              string = "Palmita"
	nameJapanese             string = "アンヌ"
	nameLatinAmericanSpanish string = "Patidifú"
	nameKorean               string = "앤"
	nameRussian              string = "Маэль"
	nameSpanish              string = "Patidifú"
	nameSimplifiedChinese    string = "安妮"
	nameTraditionalChinese   string = "安妮"
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

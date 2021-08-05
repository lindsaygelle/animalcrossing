package bubbles

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bubbles"
	nameCanadianFrench       string = "Hippy"
	nameDutch                string = "Bubbles"
	nameFrench               string = "Hippy"
	nameGerman               string = "Christin"
	nameItalian              string = "Ippolita"
	nameJapanese             string = "チャコ"
	nameLatinAmericanSpanish string = "Hipólita"
	nameKorean               string = "차코"
	nameRussian              string = "Баблс"
	nameSpanish              string = "Hipólita"
	nameSimplifiedChinese    string = "曹珂"
	nameTraditionalChinese   string = "曹珂"
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

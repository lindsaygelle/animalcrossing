package raymond

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Raymond"
	nameCanadianFrench       string = "Raymond"
	nameDutch                string = "Raymond"
	nameFrench               string = "Raymond"
	nameGerman               string = "Gunnar"
	nameItalian              string = "Raimondo"
	nameJapanese             string = "ジャック"
	nameLatinAmericanSpanish string = "Narciso"
	nameKorean               string = "잭슨"
	nameRussian              string = "Реймонд"
	nameSpanish              string = "Narciso"
	nameSimplifiedChinese    string = "杰克"
	nameTraditionalChinese   string = "傑克"
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

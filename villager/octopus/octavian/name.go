package octavian

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Octavian"
	nameCanadianFrench       string = "Octave"
	nameDutch                string = "Octavian"
	nameFrench               string = "Octave"
	nameGerman               string = "Ottfried"
	nameItalian              string = "Ottavio"
	nameJapanese             string = "おくたろう"
	nameLatinAmericanSpanish string = "Octavio"
	nameKorean               string = "문복"
	nameRussian              string = "Октавиан"
	nameSpanish              string = "Octavio"
	nameSimplifiedChinese    string = "张瑜"
	nameTraditionalChinese   string = "張瑜"
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

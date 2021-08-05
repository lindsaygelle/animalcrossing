package joey

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Joey"
	nameCanadianFrench       string = "Joseph"
	nameDutch                string = "Joey"
	nameFrench               string = "Joseph"
	nameGerman               string = "Kalle"
	nameItalian              string = "Pino"
	nameJapanese             string = "リチャード"
	nameLatinAmericanSpanish string = "Pascual"
	nameKorean               string = "리처드"
	nameRussian              string = "Джоуи"
	nameSpanish              string = "Pascual"
	nameSimplifiedChinese    string = "李察"
	nameTraditionalChinese   string = "李察"
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

package wade

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Wade"
	nameCanadianFrench       string = "Miglou"
	nameDutch                string = "Wade"
	nameFrench               string = "Miglou"
	nameGerman               string = "Staksi"
	nameItalian              string = "Plinio"
	nameJapanese             string = "カマボコ"
	nameLatinAmericanSpanish string = "Petín"
	nameKorean               string = "호떡"
	nameRussian              string = "Уэйд"
	nameSpanish              string = "Petín"
	nameSimplifiedChinese    string = "竹轮"
	nameTraditionalChinese   string = "竹輪"
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

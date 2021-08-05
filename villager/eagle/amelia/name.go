package amelia

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Amelia"
	nameCanadianFrench       string = "Aurélie"
	nameDutch                string = "Amelia"
	nameFrench               string = "Aurélie"
	nameGerman               string = "Adelheid"
	nameItalian              string = "Amelia"
	nameJapanese             string = "アンデス"
	nameLatinAmericanSpanish string = "Amelia"
	nameKorean               string = "안데스"
	nameRussian              string = "Амелия"
	nameSpanish              string = "Amelia"
	nameSimplifiedChinese    string = "安地斯"
	nameTraditionalChinese   string = "安地斯"
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

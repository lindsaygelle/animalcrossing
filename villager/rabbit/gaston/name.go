package gaston

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gaston"
	nameCanadianFrench       string = "Gaston"
	nameDutch                string = "Gaston"
	nameFrench               string = "Gaston"
	nameGerman               string = "Gaston"
	nameItalian              string = "Gaston"
	nameJapanese             string = "モサキチ"
	nameLatinAmericanSpanish string = "Gastón"
	nameKorean               string = "대길"
	nameRussian              string = "Гастон"
	nameSpanish              string = "Gastón"
	nameSimplifiedChinese    string = "猛兔"
	nameTraditionalChinese   string = "猛兔"
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

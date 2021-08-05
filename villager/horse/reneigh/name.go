package reneigh

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Reneigh"
	nameCanadianFrench       string = "Jennifer"
	nameDutch                string = "Reneigh"
	nameFrench               string = "Jennifer"
	nameGerman               string = "Andrea"
	nameItalian              string = "Requina"
	nameJapanese             string = "リアーナ"
	nameLatinAmericanSpanish string = "Corcelia"
	nameKorean               string = "리아나"
	nameRussian              string = "Рени"
	nameSpanish              string = "Corcelia"
	nameSimplifiedChinese    string = "雷哈娜"
	nameTraditionalChinese   string = "雷哈娜"
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

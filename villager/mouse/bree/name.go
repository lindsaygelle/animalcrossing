package bree

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bree"
	nameCanadianFrench       string = "Quenotte"
	nameDutch                string = "Bree"
	nameFrench               string = "Quenotte"
	nameGerman               string = "Jenny"
	nameItalian              string = "Ginella"
	nameJapanese             string = "サラ"
	nameLatinAmericanSpanish string = "Brie"
	nameKorean               string = "사라"
	nameRussian              string = "Бри"
	nameSpanish              string = "Brie"
	nameSimplifiedChinese    string = "西瓜皮"
	nameTraditionalChinese   string = "西瓜皮"
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

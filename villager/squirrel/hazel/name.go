package hazel

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hazel"
	nameCanadianFrench       string = "Pamela"
	nameDutch                string = "Hazel"
	nameFrench               string = "Pamela"
	nameGerman               string = "Nadine"
	nameItalian              string = "Cigliola"
	nameJapanese             string = "アイリス"
	nameLatinAmericanSpanish string = "Anacarda"
	nameKorean               string = "아이리스"
	nameRussian              string = "Хейзел"
	nameSpanish              string = "Anacarda"
	nameSimplifiedChinese    string = "艾栗丝"
	nameTraditionalChinese   string = "艾栗絲"
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

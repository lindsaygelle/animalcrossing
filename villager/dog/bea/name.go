package bea

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bea"
	nameCanadianFrench       string = "Béa"
	nameDutch                string = "Bea"
	nameFrench               string = "Béa"
	nameGerman               string = "Bea"
	nameItalian              string = "Cucciola"
	nameJapanese             string = "ベーグル"
	nameLatinAmericanSpanish string = "Bea"
	nameKorean               string = "베이글"
	nameRussian              string = "Беа"
	nameSpanish              string = "Bea"
	nameSimplifiedChinese    string = "贝果"
	nameTraditionalChinese   string = "貝果"
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

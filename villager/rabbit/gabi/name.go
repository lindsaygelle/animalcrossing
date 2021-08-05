package gabi

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gabi"
	nameCanadianFrench       string = "Gaby"
	nameDutch                string = "Gabi"
	nameFrench               string = "Gaby"
	nameGerman               string = "Gabi"
	nameItalian              string = "Gabri"
	nameJapanese             string = "ペチカ"
	nameLatinAmericanSpanish string = "Piluca"
	nameKorean               string = "패티카"
	nameRussian              string = "Габи"
	nameSpanish              string = "Piluca"
	nameSimplifiedChinese    string = "珮琪"
	nameTraditionalChinese   string = "珮琪"
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

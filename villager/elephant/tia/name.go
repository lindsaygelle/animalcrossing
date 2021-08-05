package tia

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tia"
	nameCanadianFrench       string = "Fanny"
	nameDutch                string = "Tia"
	nameFrench               string = "Fanny"
	nameGerman               string = "Eleonore"
	nameItalian              string = "Fanny"
	nameJapanese             string = "ティーナ"
	nameLatinAmericanSpanish string = "Meralda"
	nameKorean               string = "티나"
	nameRussian              string = "Тиа"
	nameSpanish              string = "Meralda"
	nameSimplifiedChinese    string = "茉莉"
	nameTraditionalChinese   string = "茉莉"
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

package pate

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pate"
	nameCanadianFrench       string = "Terrine"
	nameDutch                string = "Pate"
	nameFrench               string = "Terrine"
	nameGerman               string = "Daune"
	nameItalian              string = "Camilla"
	nameJapanese             string = "ナッキー"
	nameLatinAmericanSpanish string = "Pati"
	nameKorean               string = "나키"
	nameRussian              string = "Пэйт"
	nameSpanish              string = "Pati"
	nameSimplifiedChinese    string = "爱哭鬼"
	nameTraditionalChinese   string = "愛哭鬼"
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

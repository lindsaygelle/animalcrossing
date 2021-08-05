package anabelle

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Anabelle"
	nameCanadianFrench       string = "Anabelle"
	nameDutch                string = "Anabelle"
	nameFrench               string = "Anabelle"
	nameGerman               string = "Annabell"
	nameItalian              string = "Annalisa"
	nameJapanese             string = "あるみ"
	nameLatinAmericanSpanish string = "Anabel"
	nameKorean               string = "아롱이"
	nameRussian              string = "Анабель"
	nameSpanish              string = "Anabel"
	nameSimplifiedChinese    string = "有美"
	nameTraditionalChinese   string = "有美"
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

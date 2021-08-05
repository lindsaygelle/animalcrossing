package tutu

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tutu"
	nameCanadianFrench       string = "Tutu"
	nameDutch                string = "Tutu"
	nameFrench               string = "Tutu"
	nameGerman               string = "Mandy"
	nameItalian              string = "Lola"
	nameJapanese             string = "れんにゅう"
	nameLatinAmericanSpanish string = "Tutú"
	nameKorean               string = "연유"
	nameRussian              string = "Туту"
	nameSpanish              string = "Tutú"
	nameSimplifiedChinese    string = "恋恋"
	nameTraditionalChinese   string = "戀戀"
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

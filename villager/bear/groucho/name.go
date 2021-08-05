package groucho

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Groucho"
	nameCanadianFrench       string = "Ronchon"
	nameDutch                string = "Groucho"
	nameFrench               string = "Ronchon"
	nameGerman               string = "Ernst"
	nameItalian              string = "Groucho"
	nameJapanese             string = "クロー"
	nameLatinAmericanSpanish string = "Groucho"
	nameKorean               string = "거무틱"
	nameRussian              string = "Граучо"
	nameSpanish              string = "Groucho"
	nameSimplifiedChinese    string = "爪爪"
	nameTraditionalChinese   string = "爪爪"
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

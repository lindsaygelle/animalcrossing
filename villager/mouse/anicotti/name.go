package anicotti

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Anicotti"
	nameCanadianFrench       string = "Annie"
	nameDutch                string = "Anicotti"
	nameFrench               string = "Annie"
	nameGerman               string = "Eva"
	nameItalian              string = "Squitta"
	nameJapanese             string = "ラザニア"
	nameLatinAmericanSpanish string = "Clorinda"
	nameKorean               string = "라자냐"
	nameRussian              string = "Аникотти"
	nameSpanish              string = "Clorinda"
	nameSimplifiedChinese    string = "罗萱儿"
	nameTraditionalChinese   string = "羅萱兒"
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

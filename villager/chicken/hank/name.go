package hank

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hank"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "cocorico"
	nameGerman               string = "bakbak"
	nameItalian              string = "boh boh"
	nameJapanese             string = "じゃん"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "cocorico"
	nameSimplifiedChinese    string = "啾啾"
	nameTraditionalChinese   string = ""
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

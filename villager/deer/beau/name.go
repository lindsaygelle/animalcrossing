package beau

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Beau"
	nameCanadianFrench       string = "Stefaon"
	nameDutch                string = "Beau"
	nameFrench               string = "Stefaon"
	nameGerman               string = "Martin"
	nameItalian              string = "Gigi"
	nameJapanese             string = "ペーター"
	nameLatinAmericanSpanish string = "Lope"
	nameKorean               string = "피터"
	nameRussian              string = "Бью"
	nameSpanish              string = "Lope"
	nameSimplifiedChinese    string = "彼得"
	nameTraditionalChinese   string = "彼得"
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

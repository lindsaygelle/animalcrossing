package biff

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Biff"
	nameCanadianFrench       string = "Biff"
	nameDutch                string = "Biff"
	nameFrench               string = "Biff"
	nameGerman               string = "Norbert"
	nameItalian              string = "Pippo"
	nameJapanese             string = "ガブリエル"
	nameLatinAmericanSpanish string = "Pipo"
	nameKorean               string = "가브리엘"
	nameRussian              string = "Бифф"
	nameSpanish              string = "Pipo"
	nameSimplifiedChinese    string = "贾宝为"
	nameTraditionalChinese   string = "賈寶為"
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

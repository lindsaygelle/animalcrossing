package pudge

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pudge"
	nameCanadianFrench       string = "Gradub"
	nameDutch                string = "Pudge"
	nameFrench               string = "Gradub"
	nameGerman               string = "Bertram"
	nameItalian              string = "Tombolo"
	nameJapanese             string = "きんぞう"
	nameLatinAmericanSpanish string = "Bollito"
	nameKorean               string = "우띠"
	nameRussian              string = "Падж"
	nameSpanish              string = "Bollito"
	nameSimplifiedChinese    string = "阿钦"
	nameTraditionalChinese   string = "阿欽"
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

package phoebe

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Phoebe"
	nameCanadianFrench       string = "Faustine"
	nameDutch                string = "Phoebe"
	nameFrench               string = "Faustine"
	nameGerman               string = "Philippa"
	nameItalian              string = "Fenicia"
	nameJapanese             string = "ヒノコ"
	nameLatinAmericanSpanish string = "Avelina"
	nameKorean               string = "휘니"
	nameRussian              string = "Фиби"
	nameSpanish              string = "Avelina"
	nameSimplifiedChinese    string = "火鸟"
	nameTraditionalChinese   string = "火鳥"
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

package annalise

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Annalise"
	nameCanadianFrench       string = "Âne-lise"
	nameDutch                string = "Annalise"
	nameFrench               string = "Âne-lise"
	nameGerman               string = "Annerose"
	nameItalian              string = "Lisa"
	nameJapanese             string = "シルブプレ"
	nameLatinAmericanSpanish string = "Isadora"
	nameKorean               string = "실부플레"
	nameRussian              string = "Лиза"
	nameSpanish              string = "Isadora"
	nameSimplifiedChinese    string = "舒芙蕾"
	nameTraditionalChinese   string = "舒芙蕾"
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

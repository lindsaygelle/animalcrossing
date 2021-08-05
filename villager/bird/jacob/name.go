package jacob

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Jacob"
	nameCanadianFrench       string = "Jacob"
	nameDutch                string = "Jacob"
	nameFrench               string = "Jacob"
	nameGerman               string = "Jesko"
	nameItalian              string = "Valerio"
	nameJapanese             string = "ジャコテン"
	nameLatinAmericanSpanish string = "Repollo"
	nameKorean               string = "야곱"
	nameRussian              string = "Джейкоб"
	nameSpanish              string = "Repollo"
	nameSimplifiedChinese    string = "余板"
	nameTraditionalChinese   string = "余板"
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

package monty

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Monty"
	nameCanadianFrench       string = "Lourant"
	nameDutch                string = "Monty"
	nameFrench               string = "Lourant"
	nameGerman               string = "Daniel"
	nameItalian              string = "Mimmo"
	nameJapanese             string = "サルモンティ"
	nameLatinAmericanSpanish string = "Burton"
	nameKorean               string = "몽티"
	nameRussian              string = "Монти"
	nameSpanish              string = "Burton"
	nameSimplifiedChinese    string = "皮猴"
	nameTraditionalChinese   string = "皮猴"
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

package dizzy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Dizzy"
	nameCanadianFrench       string = "Pachy"
	nameDutch                string = "Dizzy"
	nameFrench               string = "Pachy"
	nameGerman               string = "Bastian"
	nameItalian              string = "Annibale"
	nameJapanese             string = "ヒュージ"
	nameLatinAmericanSpanish string = "Quique"
	nameKorean               string = "휴지"
	nameRussian              string = "Диззи"
	nameSpanish              string = "Quique"
	nameSimplifiedChinese    string = "巨巨"
	nameTraditionalChinese   string = "巨巨"
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

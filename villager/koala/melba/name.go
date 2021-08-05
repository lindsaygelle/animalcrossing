package melba

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Melba"
	nameCanadianFrench       string = "Melba"
	nameDutch                string = "Melba"
	nameFrench               string = "Melba"
	nameGerman               string = "Kornelia"
	nameItalian              string = "Melba"
	nameJapanese             string = "アデレード"
	nameLatinAmericanSpanish string = "Melba"
	nameKorean               string = "아델레이드"
	nameRussian              string = "Мельба"
	nameSpanish              string = "Melba"
	nameSimplifiedChinese    string = "阿得"
	nameTraditionalChinese   string = "阿得"
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

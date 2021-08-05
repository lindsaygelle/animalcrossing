package pietro

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pietro"
	nameCanadianFrench       string = "Pietro"
	nameDutch                string = "Pietro"
	nameFrench               string = "Pietro"
	nameGerman               string = "Pietro"
	nameItalian              string = "Giulivo"
	nameJapanese             string = "ジュペッティ"
	nameLatinAmericanSpanish string = "Kikolor"
	nameKorean               string = "피엘"
	nameRussian              string = "Пьетро"
	nameSpanish              string = "Kikolor"
	nameSimplifiedChinese    string = "皮耶罗"
	nameTraditionalChinese   string = "皮耶羅"
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

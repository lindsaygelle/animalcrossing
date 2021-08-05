package ed

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ed"
	nameCanadianFrench       string = "Édouard"
	nameDutch                string = "Ed"
	nameFrench               string = "Édouard"
	nameGerman               string = "Hermann"
	nameItalian              string = "Codino"
	nameJapanese             string = "キザノホマレ"
	nameLatinAmericanSpanish string = "Crinaldo"
	nameKorean               string = "꺼벙"
	nameRussian              string = "Эд"
	nameSpanish              string = "Crinaldo"
	nameSimplifiedChinese    string = "马誉"
	nameTraditionalChinese   string = "馬譽"
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

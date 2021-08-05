package genji

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Genji"
	nameCanadianFrench       string = "Kali"
	nameDutch                string = "Genji"
	nameFrench               string = "Kali"
	nameGerman               string = "Aki"
	nameItalian              string = "Genji"
	nameJapanese             string = "ゲンジ"
	nameLatinAmericanSpanish string = "Sumo"
	nameKorean               string = "토시"
	nameRussian              string = "Гэндзи"
	nameSpanish              string = "Sumo"
	nameSimplifiedChinese    string = "儒林"
	nameTraditionalChinese   string = "儒林"
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

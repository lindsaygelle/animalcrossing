package naomi

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Naomi"
	nameCanadianFrench       string = "Maiko"
	nameDutch                string = "Naomi"
	nameFrench               string = "Maiko"
	nameGerman               string = "Jolanda"
	nameItalian              string = "Melina"
	nameJapanese             string = "ハナコ"
	nameLatinAmericanSpanish string = "Regina"
	nameKorean               string = "화자"
	nameRussian              string = "Наоми"
	nameSpanish              string = "Regina"
	nameSimplifiedChinese    string = "玉兰"
	nameTraditionalChinese   string = "玉蘭"
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

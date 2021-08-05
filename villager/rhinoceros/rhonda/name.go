package rhonda

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rhonda"
	nameCanadianFrench       string = "Renée"
	nameDutch                string = "Rhonda"
	nameFrench               string = "Renée"
	nameGerman               string = "Regina"
	nameItalian              string = "Roby"
	nameJapanese             string = "ユメコ"
	nameLatinAmericanSpanish string = "Ronda"
	nameKorean               string = "론다"
	nameRussian              string = "Ронда"
	nameSpanish              string = "Ronda"
	nameSimplifiedChinese    string = "梦梦"
	nameTraditionalChinese   string = "夢夢"
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

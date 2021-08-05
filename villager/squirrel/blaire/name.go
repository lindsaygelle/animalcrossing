package blaire

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Blaire"
	nameCanadianFrench       string = "Cachou"
	nameDutch                string = "Blaire"
	nameFrench               string = "Cachou"
	nameGerman               string = "Klara"
	nameItalian              string = "Ghianda"
	nameJapanese             string = "シルエット"
	nameLatinAmericanSpanish string = "Azabache"
	nameKorean               string = "실루엣"
	nameRussian              string = "Блер"
	nameSpanish              string = "Azabache"
	nameSimplifiedChinese    string = "小影"
	nameTraditionalChinese   string = "小影"
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

package vic

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Vic"
	nameCanadianFrench       string = "Toto"
	nameDutch                string = "Vic"
	nameFrench               string = "Toto"
	nameGerman               string = "Klaus"
	nameItalian              string = "Paco"
	nameJapanese             string = "ノルマン"
	nameLatinAmericanSpanish string = "Artorito"
	nameKorean               string = "노르망"
	nameRussian              string = "Вик"
	nameSpanish              string = "Artorito"
	nameSimplifiedChinese    string = "卢尔曼"
	nameTraditionalChinese   string = "盧爾曼"
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

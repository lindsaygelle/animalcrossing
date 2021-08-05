package poncho

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Poncho"
	nameCanadianFrench       string = "Théo"
	nameDutch                string = "Poncho"
	nameFrench               string = "Théo"
	nameGerman               string = "Toni"
	nameItalian              string = "Poncho"
	nameJapanese             string = "ポンチョ"
	nameLatinAmericanSpanish string = "Poncho"
	nameKorean               string = "봉추"
	nameRussian              string = "Пончо"
	nameSpanish              string = "Poncho"
	nameSimplifiedChinese    string = "蓬蓬"
	nameTraditionalChinese   string = "蓬蓬"
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

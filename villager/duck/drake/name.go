package drake

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Drake"
	nameCanadianFrench       string = "Colvert"
	nameDutch                string = "Drake"
	nameFrench               string = "Colvert"
	nameGerman               string = "Gustav"
	nameItalian              string = "Drago"
	nameJapanese             string = "フォアグラ"
	nameLatinAmericanSpanish string = "Draco"
	nameKorean               string = "푸아그라"
	nameRussian              string = "Дрейк"
	nameSpanish              string = "Draco"
	nameSimplifiedChinese    string = "肥肝"
	nameTraditionalChinese   string = "肥肝"
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

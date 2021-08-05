package rudy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rudy"
	nameCanadianFrench       string = "Rougepif"
	nameDutch                string = "Rudy"
	nameFrench               string = "Rougepif"
	nameGerman               string = "Heinz"
	nameItalian              string = "Gomitolo"
	nameJapanese             string = "チャス"
	nameLatinAmericanSpanish string = "Rufino"
	nameKorean               string = "찰스"
	nameRussian              string = "Руди"
	nameSpanish              string = "Rufino"
	nameSimplifiedChinese    string = "茶茶"
	nameTraditionalChinese   string = "茶茶"
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

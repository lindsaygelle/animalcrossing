package freya

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Freya"
	nameCanadianFrench       string = "Luppa"
	nameDutch                string = "Freya"
	nameFrench               string = "Luppa"
	nameGerman               string = "Freya"
	nameItalian              string = "Freya"
	nameJapanese             string = "ツンドラ"
	nameLatinAmericanSpanish string = "Lupita"
	nameKorean               string = "산드라"
	nameRussian              string = "Фрея"
	nameSpanish              string = "Lupita"
	nameSimplifiedChinese    string = "冰冰"
	nameTraditionalChinese   string = "冰冰"
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

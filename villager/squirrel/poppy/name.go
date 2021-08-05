package poppy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Poppy"
	nameCanadianFrench       string = "Irène"
	nameDutch                string = "Poppy"
	nameFrench               string = "Irène"
	nameGerman               string = "Trita"
	nameItalian              string = "Granella"
	nameJapanese             string = "グミ"
	nameLatinAmericanSpanish string = "Encina"
	nameKorean               string = "다람"
	nameRussian              string = "Поппи"
	nameSpanish              string = "Encina"
	nameSimplifiedChinese    string = "软糖"
	nameTraditionalChinese   string = "軟糖"
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

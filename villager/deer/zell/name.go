package zell

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Zell"
	nameCanadianFrench       string = "Régis"
	nameDutch                string = "Zell"
	nameFrench               string = "Régis"
	nameGerman               string = "Walter"
	nameItalian              string = "Antilio"
	nameJapanese             string = "ネルソン"
	nameLatinAmericanSpanish string = "Corvilo"
	nameKorean               string = "넬슨"
	nameRussian              string = "Зелл"
	nameSpanish              string = "Corvilo"
	nameSimplifiedChinese    string = "森森"
	nameTraditionalChinese   string = "森森"
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

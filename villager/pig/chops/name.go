package chops

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chops"
	nameCanadianFrench       string = "Aaron"
	nameDutch                string = "Chops"
	nameFrench               string = "Aaron"
	nameGerman               string = "Clemens"
	nameItalian              string = "Lino"
	nameJapanese             string = "トンファン"
	nameLatinAmericanSpanish string = "Solomín"
	nameKorean               string = "돈후앙"
	nameRussian              string = "Чопс"
	nameSpanish              string = "Solomín"
	nameSimplifiedChinese    string = "豚皇"
	nameTraditionalChinese   string = "豚皇"
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

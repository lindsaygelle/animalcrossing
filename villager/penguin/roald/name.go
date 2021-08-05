package roald

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Roald"
	nameCanadianFrench       string = "Reynald"
	nameDutch                string = "Roald"
	nameFrench               string = "Reynald"
	nameGerman               string = "Roland"
	nameItalian              string = "Angelino"
	nameJapanese             string = "ペンタ"
	nameLatinAmericanSpanish string = "Bobi"
	nameKorean               string = "펭수"
	nameRussian              string = "Роальд"
	nameSpanish              string = "Bobi"
	nameSimplifiedChinese    string = "企鹅达"
	nameTraditionalChinese   string = "企鵝達"
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

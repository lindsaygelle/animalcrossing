package nibbles

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Nibbles"
	nameCanadianFrench       string = "Lola"
	nameDutch                string = "Nibbles"
	nameFrench               string = "Lola"
	nameGerman               string = "Knuspi"
	nameItalian              string = "Rododea"
	nameJapanese             string = "ガリガリ"
	nameLatinAmericanSpanish string = "Dentina"
	nameKorean               string = "그리미"
	nameRussian              string = "Ниблз"
	nameSpanish              string = "Dentina"
	nameSimplifiedChinese    string = "咬咬"
	nameTraditionalChinese   string = "咬咬"
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

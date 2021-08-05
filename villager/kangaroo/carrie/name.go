package carrie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Carrie"
	nameCanadianFrench       string = "Kanga"
	nameDutch                string = "Carrie"
	nameFrench               string = "Kanga"
	nameGerman               string = "Carola"
	nameItalian              string = "Taschina"
	nameJapanese             string = "マミィ"
	nameLatinAmericanSpanish string = "Lola"
	nameKorean               string = "마미"
	nameRussian              string = "Кэрри"
	nameSpanish              string = "Lola"
	nameSimplifiedChinese    string = "妈咪"
	nameTraditionalChinese   string = "媽咪"
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

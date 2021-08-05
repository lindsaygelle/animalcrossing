package pigleg

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pigleg"
	nameCanadianFrench       string = "Unknown"
	nameDutch                string = ""
	nameFrench               string = "Unknown"
	nameGerman               string = "Unknown"
	nameItalian              string = "grugno"
	nameJapanese             string = "ヨーソロ"
	nameLatinAmericanSpanish string = "Unknown"
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "Unknown"
	nameSimplifiedChinese    string = ""
	nameTraditionalChinese   string = ""
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

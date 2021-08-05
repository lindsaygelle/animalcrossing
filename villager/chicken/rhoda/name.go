package rhoda

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rhoda"
	nameCanadianFrench       string = "Unknown"
	nameDutch                string = ""
	nameFrench               string = "Unknown"
	nameGerman               string = "glucker"
	nameItalian              string = "skiok"
	nameJapanese             string = "コッコ"
	nameLatinAmericanSpanish string = "Unknown"
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "Unknown"
	nameSimplifiedChinese    string = "Unknown"
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

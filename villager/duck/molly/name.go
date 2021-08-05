package molly

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Molly"
	nameCanadianFrench       string = "Molly"
	nameDutch                string = "Molly"
	nameFrench               string = "Molly"
	nameGerman               string = "Monika"
	nameItalian              string = "Molly"
	nameJapanese             string = "カモミ"
	nameLatinAmericanSpanish string = "Deira"
	nameKorean               string = "귀오미"
	nameRussian              string = "Молли"
	nameSpanish              string = "Deira"
	nameSimplifiedChinese    string = "亚美"
	nameTraditionalChinese   string = "亞美"
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

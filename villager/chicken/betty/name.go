package betty

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Betty"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "côt côt"
	nameGerman               string = "klackl"
	nameItalian              string = "coccodè"
	nameJapanese             string = "だがね"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "clo-clo"
	nameSimplifiedChinese    string = "ok"
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

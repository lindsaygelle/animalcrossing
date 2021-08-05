package del

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Del"
	nameCanadianFrench       string = "Hector"
	nameDutch                string = "Del"
	nameFrench               string = "Hector"
	nameGerman               string = "Krokki"
	nameItalian              string = "Krokki"
	nameJapanese             string = "ヤマト"
	nameLatinAmericanSpanish string = "Croco"
	nameKorean               string = "파도맨"
	nameRussian              string = "Дел"
	nameSpanish              string = "Croco"
	nameSimplifiedChinese    string = "大和"
	nameTraditionalChinese   string = "大和"
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

package olivia

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Olivia"
	nameCanadianFrench       string = "Olivia"
	nameDutch                string = "Olivia"
	nameFrench               string = "Olivia"
	nameGerman               string = "Bianca"
	nameItalian              string = "Micina"
	nameJapanese             string = "オリビア"
	nameLatinAmericanSpanish string = "Olivia"
	nameKorean               string = "올리비아"
	nameRussian              string = "Оливия"
	nameSpanish              string = "Olivia"
	nameSimplifiedChinese    string = "奥莉"
	nameTraditionalChinese   string = "奧莉"
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

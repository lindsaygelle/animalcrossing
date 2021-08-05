package flurry

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Flurry"
	nameCanadianFrench       string = "Emma"
	nameDutch                string = "Flurry"
	nameFrench               string = "Emma"
	nameGerman               string = "Emilie"
	nameItalian              string = "Meringa"
	nameJapanese             string = "ゆきみ"
	nameLatinAmericanSpanish string = "Lluvia"
	nameKorean               string = "뽀야미"
	nameRussian              string = "Фларри"
	nameSpanish              string = "Lluvia"
	nameSimplifiedChinese    string = "雪美"
	nameTraditionalChinese   string = "雪美"
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

package mallary

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Mallary"
	nameCanadianFrench       string = "Mallory"
	nameDutch                string = "Mallary"
	nameFrench               string = "Mallory"
	nameGerman               string = "Marina"
	nameItalian              string = "Sofia"
	nameJapanese             string = "スミモモ"
	nameLatinAmericanSpanish string = "Mirta"
	nameKorean               string = "스미모"
	nameRussian              string = "Мэллари"
	nameSpanish              string = "Mirta"
	nameSimplifiedChinese    string = "苏米"
	nameTraditionalChinese   string = "蘇米"
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

package piper

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Piper"
	nameCanadianFrench       string = "Neige"
	nameDutch                string = "Piper"
	nameFrench               string = "Neige"
	nameGerman               string = "Iris"
	nameItalian              string = "Giuliva"
	nameJapanese             string = "レイコ"
	nameLatinAmericanSpanish string = "Bárbara"
	nameKorean               string = "파이프"
	nameRussian              string = "Пайпер"
	nameSpanish              string = "Bárbara"
	nameSimplifiedChinese    string = "丽婷"
	nameTraditionalChinese   string = "麗婷"
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

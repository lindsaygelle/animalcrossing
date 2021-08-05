package elise

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Elise"
	nameCanadianFrench       string = "Élise"
	nameDutch                string = "Elise"
	nameFrench               string = "Élise"
	nameGerman               string = "Steffi"
	nameItalian              string = "Elisa"
	nameJapanese             string = "モンこ"
	nameLatinAmericanSpanish string = "Mayra"
	nameKorean               string = "몽자"
	nameRussian              string = "Элиза"
	nameSpanish              string = "Mayra"
	nameSimplifiedChinese    string = "孟琪"
	nameTraditionalChinese   string = "孟琪"
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

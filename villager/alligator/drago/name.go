package drago

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Drago"
	nameCanadianFrench       string = "Drago"
	nameDutch                string = "Drago"
	nameFrench               string = "Drago"
	nameGerman               string = "Frederik"
	nameItalian              string = "Dragonio"
	nameJapanese             string = "タツオ"
	nameLatinAmericanSpanish string = "Dragonio"
	nameKorean               string = "용남이"
	nameRussian              string = "Драго"
	nameSpanish              string = "Dragonio"
	nameSimplifiedChinese    string = "阿龙"
	nameTraditionalChinese   string = "阿龍"
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

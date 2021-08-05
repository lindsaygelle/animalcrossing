package peanut

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Peanut"
	nameCanadianFrench       string = "Rachida"
	nameDutch                string = "Peanut"
	nameFrench               string = "Rachida"
	nameGerman               string = "Paulina"
	nameItalian              string = "Rachele"
	nameJapanese             string = "ももこ"
	nameLatinAmericanSpanish string = "Belinda"
	nameKorean               string = "핑키"
	nameRussian              string = "Пинат"
	nameSpanish              string = "Belinda"
	nameSimplifiedChinese    string = "小桃"
	nameTraditionalChinese   string = "小桃"
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

package ursala

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ursala"
	nameCanadianFrench       string = "Oursula"
	nameDutch                string = "Ursala"
	nameFrench               string = "Oursula"
	nameGerman               string = "Ursula"
	nameItalian              string = "Ursula"
	nameJapanese             string = "ネーヤ"
	nameLatinAmericanSpanish string = "Úrsula"
	nameKorean               string = "네이아"
	nameRussian              string = "Урсала"
	nameSpanish              string = "Úrsula"
	nameSimplifiedChinese    string = "妮雅"
	nameTraditionalChinese   string = "妮雅"
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

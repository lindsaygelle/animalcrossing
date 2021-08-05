package jambette

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Jambette"
	nameCanadianFrench       string = "Gambette"
	nameDutch                string = "Jambette"
	nameFrench               string = "Gambette"
	nameGerman               string = "Jeanette"
	nameItalian              string = "Giada"
	nameJapanese             string = "エスメラルダ"
	nameLatinAmericanSpanish string = "Anquita"
	nameKorean               string = "에스메랄다"
	nameRussian              string = "Джембет"
	nameSpanish              string = "Anquita"
	nameSimplifiedChinese    string = "翡翠"
	nameTraditionalChinese   string = "翡翠"
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

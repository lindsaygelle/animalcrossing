package merry

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Merry"
	nameCanadianFrench       string = "Suzy"
	nameDutch                string = "Merry"
	nameFrench               string = "Suzy"
	nameGerman               string = "Mischka"
	nameItalian              string = "Katy"
	nameJapanese             string = "さっち"
	nameLatinAmericanSpanish string = "Susi"
	nameKorean               string = "유네찌"
	nameRussian              string = "Мерри"
	nameSpanish              string = "Susi"
	nameSimplifiedChinese    string = "莎莎"
	nameTraditionalChinese   string = "莎莎"
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

package peck

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Peck"
	nameCanadianFrench       string = "Pec"
	nameDutch                string = "Peck"
	nameFrench               string = "Pec"
	nameGerman               string = "Picko"
	nameItalian              string = "Gheppio"
	nameJapanese             string = "ふみたろう"
	nameLatinAmericanSpanish string = "Picuet"
	nameKorean               string = "문대"
	nameRussian              string = "Пек"
	nameSpanish              string = "Picuet"
	nameSimplifiedChinese    string = "文雄"
	nameTraditionalChinese   string = "文雄"
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

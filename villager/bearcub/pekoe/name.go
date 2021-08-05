package pekoe

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pekoe"
	nameCanadianFrench       string = "Pauline"
	nameDutch                string = "Pekoe"
	nameFrench               string = "Pauline"
	nameGerman               string = "Sandrine"
	nameItalian              string = "Mina"
	nameJapanese             string = "ジャスミン"
	nameLatinAmericanSpanish string = "Vera"
	nameKorean               string = "재스민"
	nameRussian              string = "Пеко"
	nameSpanish              string = "Vera"
	nameSimplifiedChinese    string = "佳敏"
	nameTraditionalChinese   string = "佳敏"
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

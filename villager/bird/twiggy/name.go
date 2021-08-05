package twiggy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Twiggy"
	nameCanadianFrench       string = "Titi"
	nameDutch                string = "Twiggy"
	nameFrench               string = "Titi"
	nameGerman               string = "Twiggy"
	nameItalian              string = "Titti"
	nameJapanese             string = "ピーチク"
	nameLatinAmericanSpanish string = "Tití"
	nameKorean               string = "핀틱"
	nameRussian              string = "Твигги"
	nameSpanish              string = "Tití"
	nameSimplifiedChinese    string = "叽叽"
	nameTraditionalChinese   string = "嘰嘰"
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

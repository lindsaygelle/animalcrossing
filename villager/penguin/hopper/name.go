package hopper

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hopper"
	nameCanadianFrench       string = "Victor"
	nameDutch                string = "Hopper"
	nameFrench               string = "Victor"
	nameGerman               string = "Hauke"
	nameItalian              string = "Pinguaio"
	nameJapanese             string = "ダルマン"
	nameLatinAmericanSpanish string = "Güiñón"
	nameKorean               string = "달만이"
	nameRussian              string = "Хоппер"
	nameSpanish              string = "Güiñón"
	nameSimplifiedChinese    string = "达满"
	nameTraditionalChinese   string = "達滿"
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

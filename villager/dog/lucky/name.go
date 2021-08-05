package lucky

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Lucky"
	nameCanadianFrench       string = "Ramsès"
	nameDutch                string = "Lucky"
	nameFrench               string = "Ramsès"
	nameGerman               string = "Viktor"
	nameItalian              string = "Felice"
	nameJapanese             string = "ラッキー"
	nameLatinAmericanSpanish string = "Pupas"
	nameKorean               string = "럭키"
	nameRussian              string = "Лаки"
	nameSpanish              string = "Pupas"
	nameSimplifiedChinese    string = "大吉"
	nameTraditionalChinese   string = "大吉"
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

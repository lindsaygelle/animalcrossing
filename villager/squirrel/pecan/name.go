package pecan

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pecan"
	nameCanadianFrench       string = "Pécan"
	nameDutch                string = "Pecan"
	nameFrench               string = "Pécan"
	nameGerman               string = "Noisette"
	nameItalian              string = "Nocina"
	nameJapanese             string = "レベッカ"
	nameLatinAmericanSpanish string = "Camila"
	nameKorean               string = "레베카"
	nameRussian              string = "Пекан"
	nameSpanish              string = "Camila"
	nameSimplifiedChinese    string = "雷贝嘉"
	nameTraditionalChinese   string = "雷貝嘉"
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

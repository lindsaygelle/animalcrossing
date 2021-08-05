package peewee

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Peewee"
	nameCanadianFrench       string = "Gogo"
	nameDutch                string = "Peewee"
	nameFrench               string = "Gogo"
	nameGerman               string = "Manfred"
	nameItalian              string = "Kong"
	nameJapanese             string = "ダンベル"
	nameLatinAmericanSpanish string = "Lotar"
	nameKorean               string = "덤벨"
	nameRussian              string = "Пиви"
	nameSpanish              string = "Lotar"
	nameSimplifiedChinese    string = "哑铃"
	nameTraditionalChinese   string = "啞鈴"
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

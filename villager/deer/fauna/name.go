package fauna

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Fauna"
	nameCanadianFrench       string = "Bibi"
	nameDutch                string = "Fauna"
	nameFrench               string = "Bibi"
	nameGerman               string = "Fatima"
	nameItalian              string = "Cervina"
	nameJapanese             string = "ドレミ"
	nameLatinAmericanSpanish string = "Fauna"
	nameKorean               string = "솔미"
	nameRussian              string = "Фауна"
	nameSpanish              string = "Fauna"
	nameSimplifiedChinese    string = "音音"
	nameTraditionalChinese   string = "音音"
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

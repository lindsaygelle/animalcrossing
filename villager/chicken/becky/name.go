package becky

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Becky"
	nameCanadianFrench       string = "Sonya"
	nameDutch                string = "Becky"
	nameFrench               string = "Sonya"
	nameGerman               string = "Inga"
	nameItalian              string = "Annetta"
	nameJapanese             string = "アリア"
	nameLatinAmericanSpanish string = "Ramina"
	nameKorean               string = "아리아"
	nameRussian              string = "Бекки"
	nameSpanish              string = "Ramina"
	nameSimplifiedChinese    string = "咏旋"
	nameTraditionalChinese   string = "詠旋"
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

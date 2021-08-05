package stella

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Stella"
	nameCanadianFrench       string = "Bigoudi"
	nameDutch                string = "Stella"
	nameFrench               string = "Bigoudi"
	nameGerman               string = "Stella"
	nameItalian              string = "Merina"
	nameJapanese             string = "アクリル"
	nameLatinAmericanSpanish string = "Merina"
	nameKorean               string = "아크릴"
	nameRussian              string = "Стелла"
	nameSpanish              string = "Merina"
	nameSimplifiedChinese    string = "毡毡"
	nameTraditionalChinese   string = "氈氈"
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

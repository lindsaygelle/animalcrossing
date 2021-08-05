package eunice

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Eunice"
	nameCanadianFrench       string = "Bérénice"
	nameDutch                string = "Eunice"
	nameFrench               string = "Bérénice"
	nameGerman               string = "Edith"
	nameItalian              string = "Tenerina"
	nameJapanese             string = "モヘア"
	nameLatinAmericanSpanish string = "Lanolina"
	nameKorean               string = "곱슬이"
	nameRussian              string = "Юнис"
	nameSpanish              string = "Lanolina"
	nameSimplifiedChinese    string = "毛海儿"
	nameTraditionalChinese   string = "毛海兒"
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

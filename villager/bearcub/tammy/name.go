package tammy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tammy"
	nameCanadianFrench       string = "Vanessa"
	nameDutch                string = "Tammy"
	nameFrench               string = "Vanessa"
	nameGerman               string = "Tatjana"
	nameItalian              string = "Tamara"
	nameJapanese             string = "アネッサ"
	nameLatinAmericanSpanish string = "Aída"
	nameKorean               string = "아네사"
	nameRussian              string = "Тамми"
	nameSpanish              string = "Aída"
	nameSimplifiedChinese    string = "然姐"
	nameTraditionalChinese   string = "然姐"
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

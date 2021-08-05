package mathilda

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Mathilda"
	nameCanadianFrench       string = "Mathilde"
	nameDutch                string = "Mathilda"
	nameFrench               string = "Mathilde"
	nameGerman               string = "Marga"
	nameItalian              string = "Matilda"
	nameJapanese             string = "アザラク"
	nameLatinAmericanSpanish string = "Pugilda"
	nameKorean               string = "안젤라"
	nameRussian              string = "Матильда"
	nameSpanish              string = "Pugilda"
	nameSimplifiedChinese    string = "亚莎"
	nameTraditionalChinese   string = "亞莎"
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

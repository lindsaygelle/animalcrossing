package francine

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Francine"
	nameCanadianFrench       string = "Nadine"
	nameDutch                string = "Francine"
	nameFrench               string = "Nadine"
	nameGerman               string = "Manu"
	nameItalian              string = "Franca"
	nameJapanese             string = "フランソワ"
	nameLatinAmericanSpanish string = "Natacha"
	nameKorean               string = "프랑소와"
	nameRussian              string = "Франсин"
	nameSpanish              string = "Natacha"
	nameSimplifiedChinese    string = "法蓝琪"
	nameTraditionalChinese   string = "法藍琪"
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

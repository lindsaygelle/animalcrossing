package colton

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Colton"
	nameCanadianFrench       string = "Tony"
	nameDutch                string = "Colton"
	nameFrench               string = "Tony"
	nameGerman               string = "Carsten"
	nameItalian              string = "Furio"
	nameJapanese             string = "アンソニー"
	nameLatinAmericanSpanish string = "Furio"
	nameKorean               string = "안소니"
	nameRussian              string = "Колтон"
	nameSpanish              string = "Furio"
	nameSimplifiedChinese    string = "安索尼"
	nameTraditionalChinese   string = "安索尼"
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

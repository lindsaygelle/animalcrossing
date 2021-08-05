package billy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Billy"
	nameCanadianFrench       string = "Seguin"
	nameDutch                string = "Billy"
	nameFrench               string = "Seguin"
	nameGerman               string = "Hennes"
	nameItalian              string = "Billy"
	nameJapanese             string = "アーシンド"
	nameLatinAmericanSpanish string = "Brito"
	nameKorean               string = "힘드러"
	nameRussian              string = "Билли"
	nameSpanish              string = "Brito"
	nameSimplifiedChinese    string = "亚星"
	nameTraditionalChinese   string = "亞星"
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

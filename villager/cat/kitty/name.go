package kitty

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Kitty"
	nameCanadianFrench       string = "Kitty"
	nameDutch                string = "Kitty"
	nameFrench               string = "Kitty"
	nameGerman               string = "Karen"
	nameItalian              string = "Ester"
	nameJapanese             string = "ショコラ"
	nameLatinAmericanSpanish string = "Kasandra"
	nameKorean               string = "쇼콜라"
	nameRussian              string = "Китти"
	nameSpanish              string = "Kasandra"
	nameSimplifiedChinese    string = "朱古莉"
	nameTraditionalChinese   string = "朱古莉"
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

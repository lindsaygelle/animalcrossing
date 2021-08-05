package sandy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sandy"
	nameCanadianFrench       string = "Ottie"
	nameDutch                string = "Sandy"
	nameFrench               string = "Ottie"
	nameGerman               string = "Senta"
	nameItalian              string = "Alessia"
	nameJapanese             string = "ラン"
	nameLatinAmericanSpanish string = "Priscila"
	nameKorean               string = "샌디"
	nameRussian              string = "Сэнди"
	nameSpanish              string = "Priscila"
	nameSimplifiedChinese    string = "小岚"
	nameTraditionalChinese   string = "小嵐"
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

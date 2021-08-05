package diva

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Diva"
	nameCanadianFrench       string = "Violette"
	nameDutch                string = "Diva"
	nameFrench               string = "Violette"
	nameGerman               string = "Dörte"
	nameItalian              string = "Viola"
	nameJapanese             string = "アイーダ"
	nameLatinAmericanSpanish string = "Morania"
	nameKorean               string = "아이다"
	nameRussian              string = "Дива"
	nameSpanish              string = "Morania"
	nameSimplifiedChinese    string = "小睫"
	nameTraditionalChinese   string = "小睫"
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

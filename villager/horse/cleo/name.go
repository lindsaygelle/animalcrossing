package cleo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cleo"
	nameCanadianFrench       string = "Cléa"
	nameDutch                string = "Cleo"
	nameFrench               string = "Cléa"
	nameGerman               string = "Birgit"
	nameItalian              string = "Susanna"
	nameJapanese             string = "アイソトープ"
	nameLatinAmericanSpanish string = "Clotilde"
	nameKorean               string = "아이소토프"
	nameRussian              string = "Клео"
	nameSpanish              string = "Clotilde"
	nameSimplifiedChinese    string = "小铀"
	nameTraditionalChinese   string = "小鈾"
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

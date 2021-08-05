package chow

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chow"
	nameCanadianFrench       string = "Chulin"
	nameDutch                string = "Chow"
	nameFrench               string = "Chulin"
	nameGerman               string = "Chang"
	nameItalian              string = "Chowchow"
	nameJapanese             string = "チャウヤン"
	nameLatinAmericanSpanish string = "Pando"
	nameKorean               string = "츄양"
	nameRussian              string = "Чау"
	nameSpanish              string = "Pando"
	nameSimplifiedChinese    string = "朝阳"
	nameTraditionalChinese   string = "朝陽"
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

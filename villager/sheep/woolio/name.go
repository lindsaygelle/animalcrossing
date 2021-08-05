package woolio

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Woolio"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "bêêêê quoi"
	nameGerman               string = "bizääää"
	nameItalian              string = "sbaallo"
	nameJapanese             string = "ヨロシク"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "paaasssa"
	nameSimplifiedChinese    string = "帮忙"
	nameTraditionalChinese   string = ""
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

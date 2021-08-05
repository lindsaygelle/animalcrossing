package maggie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Maggie"
	nameCanadianFrench       string = "Marjorie"
	nameDutch                string = "Maggie"
	nameFrench               string = "Marjorie"
	nameGerman               string = "Magda"
	nameItalian              string = "Marcella"
	nameJapanese             string = "マーガレット"
	nameLatinAmericanSpanish string = "Margarí"
	nameKorean               string = "마가렛"
	nameRussian              string = "Мэгги"
	nameSpanish              string = "Margarí"
	nameSimplifiedChinese    string = "玛格"
	nameTraditionalChinese   string = "瑪格"
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

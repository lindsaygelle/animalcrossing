package iggly

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Iggly"
	nameCanadianFrench       string = "Urbain"
	nameDutch                string = "Iggly"
	nameFrench               string = "Urbain"
	nameGerman               string = "Pingi"
	nameItalian              string = "Piccio"
	nameJapanese             string = "のりまき"
	nameLatinAmericanSpanish string = "Bobo"
	nameKorean               string = "김말이"
	nameRussian              string = "Игли"
	nameSpanish              string = "Bobo"
	nameSimplifiedChinese    string = "花寿司"
	nameTraditionalChinese   string = "花壽司"
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

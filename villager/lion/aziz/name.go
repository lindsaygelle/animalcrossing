package aziz

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Aziz"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "zigoto"
	nameGerman               string = "RRAAHH"
	nameItalian              string = "ROAR"
	nameJapanese             string = "カンジー"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "aborigen"
	nameSimplifiedChinese    string = "罗尔"
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

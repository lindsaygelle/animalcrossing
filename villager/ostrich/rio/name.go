package rio

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rio"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "bon bref"
	nameGerman               string = "küken"
	nameItalian              string = "bocconcino"
	nameJapanese             string = "リン"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "sambasamba"
	nameSimplifiedChinese    string = "铃"
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

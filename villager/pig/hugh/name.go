package hugh

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hugh"
	nameCanadianFrench       string = "Bonno"
	nameDutch                string = "Hugh"
	nameFrench               string = "Bonno"
	nameGerman               string = "Hugo"
	nameItalian              string = "Jambon"
	nameJapanese             string = "クッチャネ"
	nameLatinAmericanSpanish string = "Jacobo"
	nameKorean               string = "먹고파"
	nameRussian              string = "Хью"
	nameSpanish              string = "Jacobo"
	nameSimplifiedChinese    string = "阿猪"
	nameTraditionalChinese   string = "阿豬"
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

package valise

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Valise"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "p'tit bout"
	nameGerman               string = "krabbel"
	nameItalian              string = "bonjour"
	nameJapanese             string = "ピョン"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "personilla"
	nameSimplifiedChinese    string = "噼哟"
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

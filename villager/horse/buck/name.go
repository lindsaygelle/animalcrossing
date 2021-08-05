package buck

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Buck"
	nameCanadianFrench       string = "Daniel"
	nameDutch                string = "Buck"
	nameFrench               string = "Daniel"
	nameGerman               string = "Rudi"
	nameItalian              string = "Trottolo"
	nameJapanese             string = "ヴァヤシコフ"
	nameLatinAmericanSpanish string = "Trotón"
	nameKorean               string = "바야시코프"
	nameRussian              string = "Бак"
	nameSpanish              string = "Trotón"
	nameSimplifiedChinese    string = "威亚"
	nameTraditionalChinese   string = "威亞"
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

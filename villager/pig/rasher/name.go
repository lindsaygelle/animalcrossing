package rasher

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rasher"
	nameCanadianFrench       string = "Salami"
	nameDutch                string = "Rasher"
	nameFrench               string = "Salami"
	nameGerman               string = "Ede"
	nameItalian              string = "Broncio"
	nameJapanese             string = "グレオ"
	nameLatinAmericanSpanish string = "Curtis"
	nameKorean               string = "글레이"
	nameRussian              string = "Рашер"
	nameSpanish              string = "Curtis"
	nameSimplifiedChinese    string = "古烈"
	nameTraditionalChinese   string = "古烈"
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

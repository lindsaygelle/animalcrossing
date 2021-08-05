package charlise

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Charlise"
	nameCanadianFrench       string = "Zabou"
	nameDutch                string = "Charlise"
	nameFrench               string = "Zabou"
	nameGerman               string = "Tabea"
	nameItalian              string = "Orsola"
	nameJapanese             string = "チャーミー"
	nameLatinAmericanSpanish string = "Charo"
	nameKorean               string = "챠미"
	nameRussian              string = "Шарлиз"
	nameSpanish              string = "Charo"
	nameSimplifiedChinese    string = "恰咪"
	nameTraditionalChinese   string = "恰咪"
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

package caroline

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Caroline"
	nameCanadianFrench       string = "Isabelle"
	nameDutch                string = "Caroline"
	nameFrench               string = "Isabelle"
	nameGerman               string = "Ricarda"
	nameItalian              string = "Carolina"
	nameJapanese             string = "キャロライン"
	nameLatinAmericanSpanish string = "Mariló"
	nameKorean               string = "캐롤라인"
	nameRussian              string = "Каролина"
	nameSpanish              string = "Mariló"
	nameSimplifiedChinese    string = "贾萝琳"
	nameTraditionalChinese   string = "賈蘿琳"
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

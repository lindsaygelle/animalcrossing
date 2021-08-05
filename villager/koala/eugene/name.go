package eugene

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Eugene"
	nameCanadianFrench       string = "Jamy"
	nameDutch                string = "Eugene"
	nameFrench               string = "Jamy"
	nameGerman               string = "Sunny"
	nameItalian              string = "Corrado"
	nameJapanese             string = "ロッキー"
	nameLatinAmericanSpanish string = "Eucalín"
	nameKorean               string = "코알"
	nameRussian              string = "Юджин"
	nameSpanish              string = "Eucalín"
	nameSimplifiedChinese    string = "罗奇"
	nameTraditionalChinese   string = "羅奇"
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

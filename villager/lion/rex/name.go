package rex

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rex"
	nameCanadianFrench       string = "Léo"
	nameDutch                string = "Rex"
	nameFrench               string = "Léo"
	nameGerman               string = "Rex"
	nameItalian              string = "Rex"
	nameJapanese             string = "サンデー"
	nameLatinAmericanSpanish string = "Leoncio"
	nameKorean               string = "렉스"
	nameRussian              string = "Рекс"
	nameSpanish              string = "Leoncio"
	nameSimplifiedChinese    string = "晴天"
	nameTraditionalChinese   string = "晴天"
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

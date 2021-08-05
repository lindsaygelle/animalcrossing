package ruby

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ruby"
	nameCanadianFrench       string = "Rubis"
	nameDutch                string = "Ruby"
	nameFrench               string = "Rubis"
	nameGerman               string = "Rubina"
	nameItalian              string = "Rubina"
	nameJapanese             string = "ルナ"
	nameLatinAmericanSpanish string = "Rubí"
	nameKorean               string = "루나"
	nameRussian              string = "Руби"
	nameSpanish              string = "Rubí"
	nameSimplifiedChinese    string = "月兔"
	nameTraditionalChinese   string = "月兔"
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

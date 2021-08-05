package agnes

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Agnes"
	nameCanadianFrench       string = "Pansy"
	nameDutch                string = "Agnes"
	nameFrench               string = "Pansy"
	nameGerman               string = "Nora"
	nameItalian              string = "Maia"
	nameJapanese             string = "アグネス"
	nameLatinAmericanSpanish string = "Nerea"
	nameKorean               string = "아그네스"
	nameRussian              string = "Агнес"
	nameSpanish              string = "Negrea"
	nameSimplifiedChinese    string = "古乃欣"
	nameTraditionalChinese   string = "古乃欣"
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

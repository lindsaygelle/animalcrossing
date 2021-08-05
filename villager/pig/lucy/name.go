package lucy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Lucy"
	nameCanadianFrench       string = "Lucie"
	nameDutch                string = "Lucy"
	nameFrench               string = "Lucie"
	nameGerman               string = "Larissa"
	nameItalian              string = "Lucy"
	nameJapanese             string = "ルーシー"
	nameLatinAmericanSpanish string = "Aurelia"
	nameKorean               string = "루시"
	nameRussian              string = "Люси"
	nameSpanish              string = "Aurelia"
	nameSimplifiedChinese    string = "露西"
	nameTraditionalChinese   string = "露西"
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

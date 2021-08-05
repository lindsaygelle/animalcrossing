package cube

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cube"
	nameCanadianFrench       string = "Cube"
	nameDutch                string = "Cube"
	nameFrench               string = "Cube"
	nameGerman               string = "Cube"
	nameItalian              string = "Cubetto"
	nameJapanese             string = "ビス"
	nameLatinAmericanSpanish string = "Cube"
	nameKorean               string = "빙수"
	nameRussian              string = "Кьюб"
	nameSpanish              string = "Cube"
	nameSimplifiedChinese    string = "罗斯"
	nameTraditionalChinese   string = "羅斯"
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

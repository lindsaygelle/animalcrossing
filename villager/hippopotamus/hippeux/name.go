package hippeux

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hippeux"
	nameCanadianFrench       string = "Paulito"
	nameDutch                string = "Hippeux"
	nameFrench               string = "Paulito"
	nameGerman               string = "Herbert"
	nameItalian              string = "Poppy"
	nameJapanese             string = "ディビッド"
	nameLatinAmericanSpanish string = "Hipóleo"
	nameKorean               string = "데이빗"
	nameRussian              string = "Гиппо"
	nameSpanish              string = "Hipóleo"
	nameSimplifiedChinese    string = "戴维"
	nameTraditionalChinese   string = "戴維"
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

package pinky

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pinky"
	nameCanadianFrench       string = "Rosine"
	nameDutch                string = "Pinky"
	nameFrench               string = "Rosine"
	nameGerman               string = "Pia"
	nameItalian              string = "Pinky"
	nameJapanese             string = "タンタン"
	nameLatinAmericanSpanish string = "Violeta"
	nameKorean               string = "링링"
	nameRussian              string = "Пинки"
	nameSpanish              string = "Violeta"
	nameSimplifiedChinese    string = "丹丹"
	nameTraditionalChinese   string = "丹丹"
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

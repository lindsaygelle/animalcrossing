package moose

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Moose"
	nameCanadianFrench       string = "Joachim"
	nameDutch                string = "Moose"
	nameFrench               string = "Joachim"
	nameGerman               string = "Mausbert"
	nameItalian              string = "Aldo"
	nameJapanese             string = "ピン"
	nameLatinAmericanSpanish string = "Sato"
	nameKorean               string = "핑"
	nameRussian              string = "Муз"
	nameSpanish              string = "Sato"
	nameSimplifiedChinese    string = "阿聘"
	nameTraditionalChinese   string = "阿聘"
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

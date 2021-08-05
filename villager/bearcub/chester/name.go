package chester

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chester"
	nameCanadianFrench       string = "Placide"
	nameDutch                string = "Chester"
	nameFrench               string = "Placide"
	nameGerman               string = "Eduard"
	nameItalian              string = "Clemente"
	nameJapanese             string = "パンタ"
	nameLatinAmericanSpanish string = "Osunio"
	nameKorean               string = "팬타"
	nameRussian              string = "Честер"
	nameSpanish              string = "Osunio"
	nameSimplifiedChinese    string = "胖达"
	nameTraditionalChinese   string = "胖達"
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

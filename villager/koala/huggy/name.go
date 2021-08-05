package huggy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Huggy"
	nameCanadianFrench       string = "nounours"
	nameDutch                string = ""
	nameFrench               string = "nounours"
	nameGerman               string = "bärchen"
	nameItalian              string = "coalilà"
	nameJapanese             string = "エヘ"
	nameLatinAmericanSpanish string = "osi-cosi"
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "osi-cosi"
	nameSimplifiedChinese    string = "啀"
	nameTraditionalChinese   string = ""
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

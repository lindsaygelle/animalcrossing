package leigh

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Leigh"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "poussinet"
	nameGerman               string = "goldi"
	nameItalian              string = "cocò"
	nameJapanese             string = "だもん"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "chuli"
	nameSimplifiedChinese    string = "啄啄"
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

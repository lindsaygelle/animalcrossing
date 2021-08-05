package bruce

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bruce"
	nameCanadianFrench       string = "Boubou"
	nameDutch                string = "Bruce"
	nameFrench               string = "Boubou"
	nameGerman               string = "Oswald"
	nameItalian              string = "Arturo"
	nameJapanese             string = "ブルース"
	nameLatinAmericanSpanish string = "Aristo"
	nameKorean               string = "브루스"
	nameRussian              string = "Брюс"
	nameSpanish              string = "Aristo"
	nameSimplifiedChinese    string = "布鹿斯"
	nameTraditionalChinese   string = "布鹿斯"
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

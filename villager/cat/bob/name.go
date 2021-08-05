package bob

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bob"
	nameCanadianFrench       string = "Robert"
	nameDutch                string = "Bob"
	nameFrench               string = "Robert"
	nameGerman               string = "Jens"
	nameItalian              string = "Bob"
	nameJapanese             string = "ニコバン"
	nameLatinAmericanSpanish string = "Arándano"
	nameKorean               string = "히죽"
	nameRussian              string = "Боб"
	nameSpanish              string = "Arándano"
	nameSimplifiedChinese    string = "阿判"
	nameTraditionalChinese   string = "阿判"
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

package epona

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Epona"
	nameCanadianFrench       string = "Épona"
	nameDutch                string = ""
	nameFrench               string = "Épona"
	nameGerman               string = "Epona"
	nameItalian              string = "Epona"
	nameJapanese             string = "エポナ"
	nameLatinAmericanSpanish string = "Epona"
	nameKorean               string = "에포나"
	nameRussian              string = ""
	nameSpanish              string = "Epona"
	nameSimplifiedChinese    string = ""
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

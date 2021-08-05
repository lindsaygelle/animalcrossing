package weber

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Weber"
	nameCanadianFrench       string = "Bébert"
	nameDutch                string = "Weber"
	nameFrench               string = "Bébert"
	nameGerman               string = "Volker"
	nameItalian              string = "Pasquale"
	nameJapanese             string = "アチョット"
	nameLatinAmericanSpanish string = "Paticio"
	nameKorean               string = "아잠만"
	nameRussian              string = "Вебер"
	nameSpanish              string = "Paticio"
	nameSimplifiedChinese    string = "卿德"
	nameTraditionalChinese   string = "卿德"
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

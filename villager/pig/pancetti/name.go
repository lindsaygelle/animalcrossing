package pancetti

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pancetti"
	nameCanadianFrench       string = "Lydie"
	nameDutch                string = "Pancetti"
	nameFrench               string = "Lydie"
	nameGerman               string = "Brigitte"
	nameItalian              string = "Cicciola"
	nameJapanese             string = "ブリトニー"
	nameLatinAmericanSpanish string = "Talía"
	nameKorean               string = "브리트니"
	nameRussian              string = "Панчетти"
	nameSpanish              string = "Talía"
	nameSimplifiedChinese    string = "布兰妮"
	nameTraditionalChinese   string = "布蘭妮"
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

package samson

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Samson"
	nameCanadianFrench       string = "Samson"
	nameDutch                string = "Samson"
	nameFrench               string = "Samson"
	nameGerman               string = "Samson"
	nameItalian              string = "Sansone"
	nameJapanese             string = "ピース"
	nameLatinAmericanSpanish string = "Sansón"
	nameKorean               string = "피스"
	nameRussian              string = "Самсон"
	nameSpanish              string = "Sansón"
	nameSimplifiedChinese    string = "和平"
	nameTraditionalChinese   string = "和平"
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

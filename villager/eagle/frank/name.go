package frank

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Frank"
	nameCanadianFrench       string = "Greggae"
	nameDutch                string = "Frank"
	nameFrench               string = "Greggae"
	nameGerman               string = "Arthur"
	nameItalian              string = "Frank"
	nameJapanese             string = "ハルク"
	nameLatinAmericanSpanish string = "Aquilino"
	nameKorean               string = "헐크"
	nameRussian              string = "Франк"
	nameSpanish              string = "Aquilino"
	nameSimplifiedChinese    string = "浩克"
	nameTraditionalChinese   string = "浩克"
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

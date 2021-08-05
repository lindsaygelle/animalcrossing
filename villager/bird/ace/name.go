package ace

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ace"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "ma caille"
	nameGerman               string = "astrein"
	nameItalian              string = "flapflap"
	nameJapanese             string = "でヤス"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "figura"
	nameSimplifiedChinese    string = "高手"
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

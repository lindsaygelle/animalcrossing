package cupcake

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cupcake"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "hum miam"
	nameGerman               string = "schätzchen"
	nameItalian              string = "zucchero"
	nameJapanese             string = "かしら"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "pastelito"
	nameSimplifiedChinese    string = "明白"
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

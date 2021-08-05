package peggy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Peggy"
	nameCanadianFrench       string = "Rose"
	nameDutch                string = "Peggy"
	nameFrench               string = "Rose"
	nameGerman               string = "Quiekie"
	nameItalian              string = "Sally"
	nameJapanese             string = "ちえり"
	nameLatinAmericanSpanish string = "Peggy"
	nameKorean               string = "체리"
	nameRussian              string = "Пегги"
	nameSpanish              string = "Peggy"
	nameSimplifiedChinese    string = "千惠"
	nameTraditionalChinese   string = "千惠"
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

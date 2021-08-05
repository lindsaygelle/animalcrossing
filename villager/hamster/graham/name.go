package graham

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Graham"
	nameCanadianFrench       string = "Graham"
	nameDutch                string = "Graham"
	nameFrench               string = "Graham"
	nameGerman               string = "Günther"
	nameItalian              string = "Evaristo"
	nameJapanese             string = "グラハム"
	nameLatinAmericanSpanish string = "Roelio"
	nameKorean               string = "글라햄"
	nameRussian              string = "Грэм"
	nameSpanish              string = "Roelio"
	nameSimplifiedChinese    string = "麦麦"
	nameTraditionalChinese   string = "麥麥"
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

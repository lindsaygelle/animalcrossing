package kidd

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Kidd"
	nameCanadianFrench       string = "Moktar"
	nameDutch                string = "Kidd"
	nameFrench               string = "Mokhtar"
	nameGerman               string = "Bocki"
	nameItalian              string = "Vittorio"
	nameJapanese             string = "やさお"
	nameLatinAmericanSpanish string = "Cabrálex"
	nameKorean               string = "염두리"
	nameRussian              string = "Кидд"
	nameSpanish              string = "Cabrálex"
	nameSimplifiedChinese    string = "文青"
	nameTraditionalChinese   string = "文青"
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

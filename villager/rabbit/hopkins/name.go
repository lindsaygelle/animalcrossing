package hopkins

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hopkins"
	nameCanadianFrench       string = "Grignote"
	nameDutch                string = "Hopkins"
	nameFrench               string = "Grignote"
	nameGerman               string = "Poldi"
	nameItalian              string = "Azeglio"
	nameJapanese             string = "プースケ"
	nameLatinAmericanSpanish string = "Saltiago"
	nameKorean               string = "홉킨스"
	nameRussian              string = "Хопкинс"
	nameSpanish              string = "Saltiago"
	nameSimplifiedChinese    string = "风杰"
	nameTraditionalChinese   string = "風杰"
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

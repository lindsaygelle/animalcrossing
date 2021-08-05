package merengue

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Merengue"
	nameCanadianFrench       string = "Patty"
	nameDutch                string = "Merengue"
	nameFrench               string = "Patty"
	nameGerman               string = "Maria"
	nameItalian              string = "Irina"
	nameJapanese             string = "パティ"
	nameLatinAmericanSpanish string = "Natura"
	nameKorean               string = "스트로베리"
	nameRussian              string = "Меренг"
	nameSpanish              string = "Natura"
	nameSimplifiedChinese    string = "草莓"
	nameTraditionalChinese   string = "草莓"
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

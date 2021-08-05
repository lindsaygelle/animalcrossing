package marcie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Marcie"
	nameCanadianFrench       string = "Selma"
	nameDutch                string = "Marcie"
	nameFrench               string = "Selma"
	nameGerman               string = "Marlies"
	nameItalian              string = "Adelaide"
	nameJapanese             string = "マリア"
	nameLatinAmericanSpanish string = "Brisa"
	nameKorean               string = "마리아"
	nameRussian              string = "Марси"
	nameSpanish              string = "Brisa"
	nameSimplifiedChinese    string = "玛莉亚"
	nameTraditionalChinese   string = "瑪莉亞"
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

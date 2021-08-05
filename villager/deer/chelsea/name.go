package chelsea

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chelsea"
	nameCanadianFrench       string = "Chelsea"
	nameDutch                string = "Chelsea"
	nameFrench               string = "Chelsea"
	nameGerman               string = "Chelsea"
	nameItalian              string = "Chelsea"
	nameJapanese             string = "チェルシー"
	nameLatinAmericanSpanish string = "Chelsea"
	nameKorean               string = "첼시"
	nameRussian              string = "Челси"
	nameSpanish              string = "Chelsea"
	nameSimplifiedChinese    string = "雀儿喜"
	nameTraditionalChinese   string = "雀兒喜"
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

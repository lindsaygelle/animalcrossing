package sylvia

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sylvia"
	nameCanadianFrench       string = "Madsi"
	nameDutch                string = "Sylvia"
	nameFrench               string = "Madsi"
	nameGerman               string = "Sylvia"
	nameItalian              string = "Silvia"
	nameJapanese             string = "シルビア"
	nameLatinAmericanSpanish string = "Cándida"
	nameKorean               string = "실비아"
	nameRussian              string = "Сильвия"
	nameSpanish              string = "Cándida"
	nameSimplifiedChinese    string = "希薇亚"
	nameTraditionalChinese   string = "希薇亞"
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

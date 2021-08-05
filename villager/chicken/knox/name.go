package knox

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Knox"
	nameCanadianFrench       string = "Wolfram"
	nameDutch                string = "Knox"
	nameFrench               string = "Wolfram"
	nameGerman               string = "Tschiwi"
	nameItalian              string = "Kalin"
	nameJapanese             string = "キンカク"
	nameLatinAmericanSpanish string = "Tono"
	nameKorean               string = "금끼오"
	nameRussian              string = "Нокс"
	nameSpanish              string = "Tono"
	nameSimplifiedChinese    string = "金阁"
	nameTraditionalChinese   string = "金閣"
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

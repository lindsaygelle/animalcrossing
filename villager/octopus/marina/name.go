package marina

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Marina"
	nameCanadianFrench       string = "Marina"
	nameDutch                string = "Marina"
	nameFrench               string = "Marina"
	nameGerman               string = "Marianne"
	nameItalian              string = "Marina"
	nameJapanese             string = "タコリーナ"
	nameLatinAmericanSpanish string = "Marina"
	nameKorean               string = "문리나"
	nameRussian              string = "Марина"
	nameSpanish              string = "Marina"
	nameSimplifiedChinese    string = "章立娜"
	nameTraditionalChinese   string = "章立娜"
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

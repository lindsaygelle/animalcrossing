package croque

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Croque"
	nameCanadianFrench       string = "Carlos"
	nameDutch                string = "Croque"
	nameFrench               string = "Carlos"
	nameGerman               string = "Carlo"
	nameItalian              string = "Gracido"
	nameJapanese             string = "タイシ"
	nameLatinAmericanSpanish string = "Ranolfo"
	nameKorean               string = "투투"
	nameRussian              string = "Крок"
	nameSpanish              string = "Ranolfo"
	nameSimplifiedChinese    string = "太子"
	nameTraditionalChinese   string = "太子"
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

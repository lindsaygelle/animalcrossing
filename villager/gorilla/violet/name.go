package violet

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Violet"
	nameCanadianFrench       string = "Gaëlle"
	nameDutch                string = "Violet"
	nameFrench               string = "Gaëlle"
	nameGerman               string = "Konga"
	nameItalian              string = "Konga"
	nameJapanese             string = "ウズメ"
	nameLatinAmericanSpanish string = "Magenta"
	nameKorean               string = "줌마"
	nameRussian              string = "Конга"
	nameSpanish              string = "Magenta"
	nameSimplifiedChinese    string = "吴紫眉"
	nameTraditionalChinese   string = "吳紫眉"
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

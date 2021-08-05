package leonardo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Leonardo"
	nameCanadianFrench       string = "Dolph"
	nameDutch                string = "Leonardo"
	nameFrench               string = "Dolph"
	nameGerman               string = "Tim"
	nameItalian              string = "Bengalo"
	nameJapanese             string = "ヒョウタ"
	nameLatinAmericanSpanish string = "Leocadio"
	nameKorean               string = "범호"
	nameRussian              string = "Леонардо"
	nameSpanish              string = "Leocadio"
	nameSimplifiedChinese    string = "阿彪"
	nameTraditionalChinese   string = "阿彪"
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

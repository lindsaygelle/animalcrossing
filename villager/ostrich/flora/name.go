package flora

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Flora"
	nameCanadianFrench       string = "Justine"
	nameDutch                string = "Flora"
	nameFrench               string = "Justine"
	nameGerman               string = "Flora"
	nameItalian              string = "Rosalba"
	nameJapanese             string = "フララ"
	nameLatinAmericanSpanish string = "Azucena"
	nameKorean               string = "플라라"
	nameRussian              string = "Флора"
	nameSpanish              string = "Azucena"
	nameSimplifiedChinese    string = "胡拉拉"
	nameTraditionalChinese   string = "胡拉拉"
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

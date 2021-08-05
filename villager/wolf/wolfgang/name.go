package wolfgang

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Wolfgang"
	nameCanadianFrench       string = "Wolfgang"
	nameDutch                string = "Wolfgang"
	nameFrench               string = "Wolfgang"
	nameGerman               string = "Weber"
	nameItalian              string = "Wolfgang"
	nameJapanese             string = "ロボ"
	nameLatinAmericanSpanish string = "Wolfi"
	nameKorean               string = "로보"
	nameRussian              string = "Вольфи"
	nameSpanish              string = "Wolfi"
	nameSimplifiedChinese    string = "罗博"
	nameTraditionalChinese   string = "羅博"
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

package dotty

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Dotty"
	nameCanadianFrench       string = "Dorothée"
	nameDutch                string = "Dotty"
	nameFrench               string = "Dorothée"
	nameGerman               string = "Doro"
	nameItalian              string = "Dotty"
	nameJapanese             string = "マーサ"
	nameLatinAmericanSpanish string = "Katia"
	nameKorean               string = "서머"
	nameRussian              string = "Дотти"
	nameSpanish              string = "Katia"
	nameSimplifiedChinese    string = "玛莎"
	nameTraditionalChinese   string = "瑪莎"
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

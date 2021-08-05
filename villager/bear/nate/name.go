package nate

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Nate"
	nameCanadianFrench       string = "Nathan"
	nameDutch                string = "Nate"
	nameFrench               string = "Nathan"
	nameGerman               string = "Nathan"
	nameItalian              string = "Gianni"
	nameJapanese             string = "バッカス"
	nameLatinAmericanSpanish string = "Nachete"
	nameKorean               string = "박하스"
	nameRussian              string = "Нэйт"
	nameSpanish              string = "Nachete"
	nameSimplifiedChinese    string = "巴克思"
	nameTraditionalChinese   string = "巴克思"
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

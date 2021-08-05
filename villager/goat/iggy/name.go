package iggy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Iggy"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "mon jojo"
	nameGerman               string = "kumpel"
	nameItalian              string = "salameee"
	nameJapanese             string = "ですじゃ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "coleguiita"
	nameSimplifiedChinese    string = "Unknown"
	nameTraditionalChinese   string = ""
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

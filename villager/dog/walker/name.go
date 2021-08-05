package walker

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Walker"
	nameCanadianFrench       string = "George"
	nameDutch                string = "Walker"
	nameFrench               string = "George"
	nameGerman               string = "Fido"
	nameItalian              string = "Walter"
	nameJapanese             string = "ベン"
	nameLatinAmericanSpanish string = "Miki"
	nameKorean               string = "벤"
	nameRussian              string = "Уокер"
	nameSpanish              string = "Miki"
	nameSimplifiedChinese    string = "阿笨"
	nameTraditionalChinese   string = "阿笨"
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

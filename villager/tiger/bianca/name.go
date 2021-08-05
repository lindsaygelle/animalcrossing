package bianca

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bianca"
	nameCanadianFrench       string = "Noémie"
	nameDutch                string = "Bianca"
	nameFrench               string = "Noémie"
	nameGerman               string = "Bettina"
	nameItalian              string = "Grattina"
	nameJapanese             string = "コユキ"
	nameLatinAmericanSpanish string = "Bianca"
	nameKorean               string = "백희"
	nameRussian              string = "Бьянка"
	nameSpanish              string = "Bianca"
	nameSimplifiedChinese    string = "小雪"
	nameTraditionalChinese   string = "小雪"
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

package annalisa

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Annalisa"
	nameCanadianFrench       string = "Roberta"
	nameDutch                string = "Annalisa"
	nameFrench               string = "Roberta"
	nameGerman               string = "Annalena"
	nameItalian              string = "Anita"
	nameJapanese             string = "みやび"
	nameLatinAmericanSpanish string = "Alba"
	nameKorean               string = "설백"
	nameRussian              string = "Аннализа"
	nameSpanish              string = "Alba"
	nameSimplifiedChinese    string = "小雅"
	nameTraditionalChinese   string = "小雅"
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

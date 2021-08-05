package queenie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Queenie"
	nameCanadianFrench       string = "Reine"
	nameDutch                string = "Queenie"
	nameFrench               string = "Reine"
	nameGerman               string = "Isabella"
	nameItalian              string = "Regina"
	nameJapanese             string = "タキュ"
	nameLatinAmericanSpanish string = "Marujita"
	nameKorean               string = "택주"
	nameRussian              string = "Квини"
	nameSpanish              string = "Marujita"
	nameSimplifiedChinese    string = "果果"
	nameTraditionalChinese   string = "果果"
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

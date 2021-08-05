package bam

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bam"
	nameCanadianFrench       string = "Nacer"
	nameDutch                string = "Bam"
	nameFrench               string = "Nacer"
	nameGerman               string = "Benjamin"
	nameItalian              string = "Cornelio"
	nameJapanese             string = "タケル"
	nameLatinAmericanSpanish string = "Cornelio"
	nameKorean               string = "록키"
	nameRussian              string = "Бам"
	nameSpanish              string = "Cornelio"
	nameSimplifiedChinese    string = "小健"
	nameTraditionalChinese   string = "小健"
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

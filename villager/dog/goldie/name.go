package goldie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Goldie"
	nameCanadianFrench       string = "Mirza"
	nameDutch                string = "Goldie"
	nameFrench               string = "Mirza"
	nameGerman               string = "Bienchen"
	nameItalian              string = "Dora"
	nameJapanese             string = "キャラメル"
	nameLatinAmericanSpanish string = "Tere"
	nameKorean               string = "카라멜"
	nameRussian              string = "Голди"
	nameSpanish              string = "Tere"
	nameSimplifiedChinese    string = "牛奶糖"
	nameTraditionalChinese   string = "牛奶糖"
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

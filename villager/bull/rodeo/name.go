package rodeo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rodeo"
	nameCanadianFrench       string = "Flèche"
	nameDutch                string = "Rodeo"
	nameFrench               string = "Flèche"
	nameGerman               string = "Toro"
	nameItalian              string = "Rodeo"
	nameJapanese             string = "ロデオ"
	nameLatinAmericanSpanish string = "Rodeo"
	nameKorean               string = "로데오"
	nameRussian              string = "Родео"
	nameSpanish              string = "Rodeo"
	nameSimplifiedChinese    string = "罗迪欧"
	nameTraditionalChinese   string = "羅迪歐"
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

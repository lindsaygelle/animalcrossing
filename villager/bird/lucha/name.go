package lucha

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Lucha"
	nameCanadianFrench       string = "Condor"
	nameDutch                string = "Lucha"
	nameFrench               string = "Condor"
	nameGerman               string = "Lukas"
	nameItalian              string = "Loreto"
	nameJapanese             string = "マスカラス"
	nameLatinAmericanSpanish string = "Plumerio"
	nameKorean               string = "마스카라스"
	nameRussian              string = "Луча"
	nameSpanish              string = "Plumerio"
	nameSimplifiedChinese    string = "摔角鸦"
	nameTraditionalChinese   string = "摔角鴉"
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

package flip

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Flip"
	nameCanadianFrench       string = "Rudy"
	nameDutch                string = "Flip"
	nameFrench               string = "Rudy"
	nameGerman               string = "Pippo"
	nameItalian              string = "Benny"
	nameJapanese             string = "さすけ"
	nameLatinAmericanSpanish string = "Monet"
	nameKorean               string = "원승"
	nameRussian              string = "Флип"
	nameSpanish              string = "Monet"
	nameSimplifiedChinese    string = "天佑"
	nameTraditionalChinese   string = "天佑"
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

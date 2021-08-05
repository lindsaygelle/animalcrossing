package simon

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Simon"
	nameCanadianFrench       string = "Simon"
	nameDutch                string = "Simon"
	nameFrench               string = "Simon"
	nameGerman               string = "Simon"
	nameItalian              string = "Simon"
	nameJapanese             string = "エテキチ"
	nameLatinAmericanSpanish string = "Simón"
	nameKorean               string = "시몬"
	nameRussian              string = "Саймон"
	nameSpanish              string = "Simón"
	nameSimplifiedChinese    string = "远仁"
	nameTraditionalChinese   string = "遠仁"
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

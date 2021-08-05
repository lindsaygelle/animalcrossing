package marty

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Marty"
	nameCanadianFrench       string = "Marty"
	nameDutch                string = "Marty"
	nameFrench               string = "Marty"
	nameGerman               string = "Marty"
	nameItalian              string = "Marty"
	nameJapanese             string = "マーティー"
	nameLatinAmericanSpanish string = "Marty"
	nameKorean               string = "마티"
	nameRussian              string = "Марти"
	nameSpanish              string = "Marty"
	nameSimplifiedChinese    string = "玛丁"
	nameTraditionalChinese   string = "瑪丁"
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

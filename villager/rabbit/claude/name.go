package claude

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Claude"
	nameCanadianFrench       string = "Claude"
	nameDutch                string = "Claude"
	nameFrench               string = "Claude"
	nameGerman               string = "Claude"
	nameItalian              string = "Claude"
	nameJapanese             string = "ビネガー"
	nameLatinAmericanSpanish string = "Pablo"
	nameKorean               string = "비니거"
	nameRussian              string = "Клод"
	nameSpanish              string = "Pablo"
	nameSimplifiedChinese    string = "阿醋"
	nameTraditionalChinese   string = "阿醋"
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

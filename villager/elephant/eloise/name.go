package eloise

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Eloise"
	nameCanadianFrench       string = "Éloïse"
	nameDutch                string = "Eloise"
	nameFrench               string = "Éloïse"
	nameGerman               string = "Frauke"
	nameItalian              string = "Eloisa"
	nameJapanese             string = "エレフィン"
	nameLatinAmericanSpanish string = "Eloísa"
	nameKorean               string = "엘레핀"
	nameRussian              string = "Элоиза"
	nameSpanish              string = "Eloísa"
	nameSimplifiedChinese    string = "艾勒芬"
	nameTraditionalChinese   string = "艾勒芬"
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

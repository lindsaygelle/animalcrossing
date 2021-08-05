package rodney

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rodney"
	nameCanadianFrench       string = "Chico"
	nameDutch                string = "Rodney"
	nameFrench               string = "Chico"
	nameGerman               string = "Reinhold"
	nameItalian              string = "Rosikkio"
	nameJapanese             string = "ジミー"
	nameLatinAmericanSpanish string = "Chusquis"
	nameKorean               string = "지미"
	nameRussian              string = "Родни"
	nameSpanish              string = "Chusquis"
	nameSimplifiedChinese    string = "米米"
	nameTraditionalChinese   string = "米米"
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

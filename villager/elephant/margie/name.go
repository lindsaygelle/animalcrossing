package margie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Margie"
	nameCanadianFrench       string = "Maguy"
	nameDutch                string = "Margie"
	nameFrench               string = "Maguy"
	nameGerman               string = "Elisa"
	nameItalian              string = "Marianna"
	nameJapanese             string = "サリー"
	nameLatinAmericanSpanish string = "Rita"
	nameKorean               string = "샐리"
	nameRussian              string = "Марджи"
	nameSpanish              string = "Rita"
	nameSimplifiedChinese    string = "莎莉"
	nameTraditionalChinese   string = "莎莉"
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

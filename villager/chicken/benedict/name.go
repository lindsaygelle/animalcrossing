package benedict

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Benedict"
	nameCanadianFrench       string = "Dimitri"
	nameDutch                string = "Benedict"
	nameFrench               string = "Dimitri"
	nameGerman               string = "Benedikt"
	nameItalian              string = "Biagio"
	nameJapanese             string = "ぺしみち"
	nameLatinAmericanSpanish string = "Benito"
	nameKorean               string = "페니실린"
	nameRussian              string = "Бенедикт"
	nameSpanish              string = "Benito"
	nameSimplifiedChinese    string = "沛希"
	nameTraditionalChinese   string = "沛希"
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

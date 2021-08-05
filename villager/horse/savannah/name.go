package savannah

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Savannah"
	nameCanadianFrench       string = "Savana"
	nameDutch                string = "Savannah"
	nameFrench               string = "Savana"
	nameGerman               string = "Zara"
	nameItalian              string = "Savannah"
	nameJapanese             string = "サバンナ"
	nameLatinAmericanSpanish string = "Sabana"
	nameKorean               string = "사반나"
	nameRussian              string = "Саванна"
	nameSpanish              string = "Sabana"
	nameSimplifiedChinese    string = "草斑娜"
	nameTraditionalChinese   string = "草斑娜"
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

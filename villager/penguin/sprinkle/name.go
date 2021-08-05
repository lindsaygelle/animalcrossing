package sprinkle

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sprinkle"
	nameCanadianFrench       string = "Laurie"
	nameDutch                string = "Sprinkle"
	nameFrench               string = "Laurie"
	nameGerman               string = "Svenja"
	nameItalian              string = "Brina"
	nameJapanese             string = "フラッペ"
	nameLatinAmericanSpanish string = "Rizolda"
	nameKorean               string = "크리미"
	nameRussian              string = "Спринкл"
	nameSpanish              string = "Rizolda"
	nameSimplifiedChinese    string = "冰莎"
	nameTraditionalChinese   string = "冰莎"
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

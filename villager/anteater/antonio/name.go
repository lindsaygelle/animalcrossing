package antonio

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Antonio"
	nameCanadianFrench       string = "Antonio"
	nameDutch                string = "Antonio"
	nameFrench               string = "Antonio"
	nameGerman               string = "Siggi"
	nameItalian              string = "Antonio"
	nameJapanese             string = "マコト"
	nameLatinAmericanSpanish string = "Antonio"
	nameKorean               string = "퍼머거"
	nameRussian              string = "Антонио"
	nameSpanish              string = "Antonio"
	nameSimplifiedChinese    string = "阿诚"
	nameTraditionalChinese   string = "阿誠"
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

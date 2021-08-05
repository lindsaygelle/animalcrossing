package holden

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Holden"
	nameCanadianFrench       string = "Glutin"
	nameDutch                string = ""
	nameFrench               string = "Glutin"
	nameGerman               string = "Holden"
	nameItalian              string = "Collotto"
	nameJapanese             string = "のりぼう"
	nameLatinAmericanSpanish string = "Holden"
	nameKorean               string = "노리보"
	nameRussian              string = ""
	nameSpanish              string = "Holden"
	nameSimplifiedChinese    string = ""
	nameTraditionalChinese   string = ""
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

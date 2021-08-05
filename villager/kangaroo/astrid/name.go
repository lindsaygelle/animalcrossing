package astrid

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Astrid"
	nameCanadianFrench       string = "Rhona"
	nameDutch                string = "Astrid"
	nameFrench               string = "Rhona"
	nameGerman               string = "Astrid"
	nameItalian              string = "Stella"
	nameJapanese             string = "キッズ"
	nameLatinAmericanSpanish string = "Astrid"
	nameKorean               string = "펑키맘"
	nameRussian              string = "Астрид"
	nameSpanish              string = "Astrid"
	nameSimplifiedChinese    string = "童儿"
	nameTraditionalChinese   string = "童兒"
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

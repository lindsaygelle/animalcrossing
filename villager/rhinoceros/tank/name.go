package tank

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tank"
	nameCanadianFrench       string = "Ben"
	nameDutch                string = "Tank"
	nameFrench               string = "Ben"
	nameGerman               string = "Frank"
	nameItalian              string = "Rinaldo"
	nameJapanese             string = "くるぶし"
	nameLatinAmericanSpanish string = "Aníbal"
	nameKorean               string = "탱크"
	nameRussian              string = "Танк"
	nameSpanish              string = "Aníbal"
	nameSimplifiedChinese    string = "阿足"
	nameTraditionalChinese   string = "阿足"
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

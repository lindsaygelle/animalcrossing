package snake

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Snake"
	nameCanadianFrench       string = "Jeannot"
	nameDutch                string = "Snake"
	nameFrench               string = "Jeannot"
	nameGerman               string = "Philip"
	nameItalian              string = "Razzo"
	nameJapanese             string = "モモチ"
	nameLatinAmericanSpanish string = "Rayo"
	nameKorean               string = "닌토"
	nameRussian              string = "Снейк"
	nameSpanish              string = "Rayo"
	nameSimplifiedChinese    string = "百川"
	nameTraditionalChinese   string = "百川"
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

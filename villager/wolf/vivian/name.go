package vivian

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Vivian"
	nameCanadianFrench       string = "Viviane"
	nameDutch                string = "Vivian"
	nameFrench               string = "Viviane"
	nameGerman               string = "Viviane"
	nameItalian              string = "Viviana"
	nameJapanese             string = "ヴァネッサ"
	nameLatinAmericanSpanish string = "Viviana"
	nameKorean               string = "바네사"
	nameRussian              string = "Вивьен"
	nameSpanish              string = "Viviana"
	nameSimplifiedChinese    string = "范妮沙"
	nameTraditionalChinese   string = "范妮莎"
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

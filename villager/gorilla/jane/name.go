package jane

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Jane"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "ouistiti"
	nameGerman               string = "freund"
	nameItalian              string = "scimmietta"
	nameJapanese             string = "だゴリ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "chimpa"
	nameSimplifiedChinese    string = "嘣嘣"
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

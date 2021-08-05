package filbert

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Filbert"
	nameCanadianFrench       string = "Filibert"
	nameDutch                string = "Filbert"
	nameFrench               string = "Filibert"
	nameGerman               string = "Felix"
	nameItalian              string = "Scriccio"
	nameJapanese             string = "リッキー"
	nameLatinAmericanSpanish string = "Filberto"
	nameKorean               string = "리키"
	nameRussian              string = "Филберт"
	nameSpanish              string = "Filberto"
	nameSimplifiedChinese    string = "瑞奇"
	nameTraditionalChinese   string = "瑞奇"
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

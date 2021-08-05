package felyne

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Felyne"
	nameCanadianFrench       string = "Felyne"
	nameDutch                string = ""
	nameFrench               string = "Felyne"
	nameGerman               string = "Felyne"
	nameItalian              string = "Felyne"
	nameJapanese             string = "アイルー"
	nameLatinAmericanSpanish string = "Felyne"
	nameKorean               string = "아이루"
	nameRussian              string = ""
	nameSpanish              string = "Felyne"
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

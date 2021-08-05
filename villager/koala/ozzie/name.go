package ozzie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ozzie"
	nameCanadianFrench       string = "Koko"
	nameDutch                string = "Ozzie"
	nameFrench               string = "Koko"
	nameGerman               string = "Oskar"
	nameItalian              string = "Ozzy"
	nameJapanese             string = "ドングリ"
	nameLatinAmericanSpanish string = "Koloa"
	nameKorean               string = "동동이"
	nameRussian              string = "Оззи"
	nameSpanish              string = "Koloa"
	nameSimplifiedChinese    string = "阿栗"
	nameTraditionalChinese   string = "阿栗"
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

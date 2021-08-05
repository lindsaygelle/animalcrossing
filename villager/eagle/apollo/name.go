package apollo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Apollo"
	nameCanadianFrench       string = "Apollon"
	nameDutch                string = "Apollo"
	nameFrench               string = "Apollon"
	nameGerman               string = "Apollo"
	nameItalian              string = "Apollo"
	nameJapanese             string = "アポロ"
	nameLatinAmericanSpanish string = "Apolo"
	nameKorean               string = "아폴로"
	nameRussian              string = "Аполло"
	nameSpanish              string = "Apolo"
	nameSimplifiedChinese    string = "阿波罗"
	nameTraditionalChinese   string = "阿波羅"
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

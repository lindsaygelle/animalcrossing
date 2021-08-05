package bonbon

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bonbon"
	nameCanadianFrench       string = "Sylvette"
	nameDutch                string = "Bonbon"
	nameFrench               string = "Sylvette"
	nameGerman               string = "Henrike"
	nameItalian              string = "Lolly"
	nameJapanese             string = "ミミィ"
	nameLatinAmericanSpanish string = "Chocolat"
	nameKorean               string = "미미"
	nameRussian              string = "Бонбон"
	nameSpanish              string = "Chocolat"
	nameSimplifiedChinese    string = "妙妙"
	nameTraditionalChinese   string = "妙妙"
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

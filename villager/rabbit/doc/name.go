package doc

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Doc"
	nameCanadianFrench       string = "Doc"
	nameDutch                string = "Doc"
	nameFrench               string = "Doc"
	nameGerman               string = "Gustl"
	nameItalian              string = "Doc"
	nameJapanese             string = "トビオ"
	nameLatinAmericanSpanish string = "Crispín"
	nameKorean               string = "토니"
	nameRussian              string = "Док"
	nameSpanish              string = "Crispín"
	nameSimplifiedChinese    string = "阿飞"
	nameTraditionalChinese   string = "阿飛"
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

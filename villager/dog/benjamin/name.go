package benjamin

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Benjamin"
	nameCanadianFrench       string = "Bernardo"
	nameDutch                string = "Benjamin"
	nameFrench               string = "Bernardo"
	nameGerman               string = "Wastl"
	nameItalian              string = "Bernardo"
	nameJapanese             string = "ハチ"
	nameLatinAmericanSpanish string = "Bernardo"
	nameKorean               string = "땡칠이"
	nameRussian              string = "Бернардо"
	nameSpanish              string = "Bernardo"
	nameSimplifiedChinese    string = "小八"
	nameTraditionalChinese   string = "小八"
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

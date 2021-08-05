package hans

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hans"
	nameCanadianFrench       string = "Loran"
	nameDutch                string = "Hans"
	nameFrench               string = "Loran"
	nameGerman               string = "Hans"
	nameItalian              string = "Grigilio"
	nameJapanese             string = "スナイル"
	nameLatinAmericanSpanish string = "Hans"
	nameKorean               string = "스나일"
	nameRussian              string = "Ганс"
	nameSpanish              string = "Hans"
	nameSimplifiedChinese    string = "史奈鲁"
	nameTraditionalChinese   string = "史奈魯"
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

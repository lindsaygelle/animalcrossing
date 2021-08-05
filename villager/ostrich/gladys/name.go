package gladys

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gladys"
	nameCanadianFrench       string = "Gladys"
	nameDutch                string = "Gladys"
	nameFrench               string = "Gladys"
	nameGerman               string = "Sandra"
	nameItalian              string = "Marta"
	nameJapanese             string = "ちとせ"
	nameLatinAmericanSpanish string = "Gladis"
	nameKorean               string = "빅토리아"
	nameRussian              string = "Глэдис"
	nameSpanish              string = "Gladis"
	nameSimplifiedChinese    string = "仙仙"
	nameTraditionalChinese   string = "仙仙"
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

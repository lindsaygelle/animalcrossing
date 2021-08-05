package chai

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chai"
	nameCanadianFrench       string = "Chaï"
	nameDutch                string = "Chai"
	nameFrench               string = "Chaï"
	nameGerman               string = "Chai"
	nameItalian              string = "Chai"
	nameJapanese             string = "フィーカ"
	nameLatinAmericanSpanish string = "Chai"
	nameKorean               string = "피카"
	nameRussian              string = "Чай"
	nameSpanish              string = "Chai"
	nameSimplifiedChinese    string = "啡卡"
	nameTraditionalChinese   string = "啡卡"
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

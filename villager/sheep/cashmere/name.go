package cashmere

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cashmere"
	nameCanadianFrench       string = "Cashmir"
	nameDutch                string = "Cashmere"
	nameFrench               string = "Cashmir"
	nameGerman               string = "Lana"
	nameItalian              string = "Cashmere"
	nameJapanese             string = "ラムール"
	nameLatinAmericanSpanish string = "Cachemir"
	nameKorean               string = "캐시미어"
	nameRussian              string = "Кашмир"
	nameSpanish              string = "Cachemir"
	nameSimplifiedChinese    string = "爱睦"
	nameTraditionalChinese   string = "愛睦"
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

package sally

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sally"
	nameCanadianFrench       string = "Damia"
	nameDutch                string = "Sally"
	nameFrench               string = "Damia"
	nameGerman               string = "Hanne"
	nameItalian              string = "Rossella"
	nameJapanese             string = "ララミー"
	nameLatinAmericanSpanish string = "Praliné"
	nameKorean               string = "라라미"
	nameRussian              string = "Салли"
	nameSpanish              string = "Praliné"
	nameSimplifiedChinese    string = "拉拉米"
	nameTraditionalChinese   string = "拉拉米"
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

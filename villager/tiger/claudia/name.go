package claudia

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Claudia"
	nameCanadianFrench       string = "Vanina"
	nameDutch                string = "Claudia"
	nameFrench               string = "Vanina"
	nameGerman               string = "Lilly"
	nameItalian              string = "Lilla"
	nameJapanese             string = "マリリン"
	nameLatinAmericanSpanish string = "Lilu"
	nameKorean               string = "신디"
	nameRussian              string = "Клаудия"
	nameSpanish              string = "Lilu"
	nameSimplifiedChinese    string = "马丽莲"
	nameTraditionalChinese   string = "馬麗蓮"
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

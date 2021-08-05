package jay

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Jay"
	nameCanadianFrench       string = "Gérard"
	nameDutch                string = "Jay"
	nameFrench               string = "Gérard"
	nameGerman               string = "Jan"
	nameItalian              string = "Jet"
	nameJapanese             string = "ツバクロ"
	nameLatinAmericanSpanish string = "Jairo"
	nameKorean               string = "참돌이"
	nameRussian              string = "Джей"
	nameSpanish              string = "Jota"
	nameSimplifiedChinese    string = "阿燕"
	nameTraditionalChinese   string = "阿燕"
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

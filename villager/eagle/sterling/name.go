package sterling

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sterling"
	nameCanadianFrench       string = "Manfred"
	nameDutch                string = "Sterling"
	nameFrench               string = "Manfred"
	nameGerman               string = "Horst"
	nameItalian              string = "Beccodì"
	nameJapanese             string = "ギンカク"
	nameLatinAmericanSpanish string = "Arni"
	nameKorean               string = "은수리"
	nameRussian              string = "Стерлинг"
	nameSpanish              string = "Arni"
	nameSimplifiedChinese    string = "银阁"
	nameTraditionalChinese   string = "銀閣"
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

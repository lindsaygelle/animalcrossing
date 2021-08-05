package olaf

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Olaf"
	nameCanadianFrench       string = "Blair"
	nameDutch                string = "Olaf"
	nameFrench               string = "Blair"
	nameGerman               string = "Olaf"
	nameItalian              string = "Ettore"
	nameJapanese             string = "アントニオ"
	nameLatinAmericanSpanish string = "Olaf"
	nameKorean               string = "안토니오"
	nameRussian              string = "Олаф"
	nameSpanish              string = "Olaf"
	nameSimplifiedChinese    string = "安东尼"
	nameTraditionalChinese   string = "安東尼"
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

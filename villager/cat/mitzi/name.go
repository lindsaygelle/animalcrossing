package mitzi

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Mitzi"
	nameCanadianFrench       string = "Mistigri"
	nameDutch                string = "Mitzi"
	nameFrench               string = "Mistigri"
	nameGerman               string = "Miezi"
	nameItalian              string = "Zampetta"
	nameJapanese             string = "マール"
	nameLatinAmericanSpanish string = "Misi"
	nameKorean               string = "마르"
	nameRussian              string = "Митци"
	nameSpanish              string = "Misi"
	nameSimplifiedChinese    string = "圆圆"
	nameTraditionalChinese   string = "圓圓"
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

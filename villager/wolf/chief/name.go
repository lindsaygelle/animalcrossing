package chief

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chief"
	nameCanadianFrench       string = "Chef"
	nameDutch                string = "Chief"
	nameFrench               string = "Chef"
	nameGerman               string = "Sascha"
	nameItalian              string = "Artiglio"
	nameJapanese             string = "チーフ"
	nameLatinAmericanSpanish string = "Zoilo"
	nameKorean               string = "대장"
	nameRussian              string = "Чиф"
	nameSpanish              string = "Zoilo"
	nameSimplifiedChinese    string = "施傅"
	nameTraditionalChinese   string = "施傅"
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

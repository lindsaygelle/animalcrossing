package mint

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Mint"
	nameCanadianFrench       string = "Amande"
	nameDutch                string = "Mint"
	nameFrench               string = "Amande"
	nameGerman               string = "Marika"
	nameItalian              string = "Mentulla"
	nameJapanese             string = "ミント"
	nameLatinAmericanSpanish string = "Menta"
	nameKorean               string = "민트"
	nameRussian              string = "Минт"
	nameSpanish              string = "Menta"
	nameSimplifiedChinese    string = "小敏"
	nameTraditionalChinese   string = "小敏"
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

package broccolo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Broccolo"
	nameCanadianFrench       string = "Steven"
	nameDutch                string = "Broccolo"
	nameFrench               string = "Steven"
	nameGerman               string = "Fritzi"
	nameItalian              string = "Franco"
	nameJapanese             string = "ブロッコリー"
	nameLatinAmericanSpanish string = "Brócoli"
	nameKorean               string = "브로콜리"
	nameRussian              string = "Брокколо"
	nameSpanish              string = "Brócoli"
	nameSimplifiedChinese    string = "茜蓝"
	nameTraditionalChinese   string = "茜藍"
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

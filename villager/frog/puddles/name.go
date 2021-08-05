package puddles

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Puddles"
	nameCanadianFrench       string = "Rénata"
	nameDutch                string = "Puddles"
	nameFrench               string = "Rénata"
	nameGerman               string = "Nele"
	nameItalian              string = "Grazia"
	nameJapanese             string = "チョキ"
	nameLatinAmericanSpanish string = "Nenúfar"
	nameKorean               string = "가위"
	nameRussian              string = "Рената"
	nameSpanish              string = "Nenúfar"
	nameSimplifiedChinese    string = "剪刀"
	nameTraditionalChinese   string = "剪刀"
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

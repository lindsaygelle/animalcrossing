package inkwell

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Inkwell"
	nameCanadianFrench       string = "Pigmento"
	nameDutch                string = ""
	nameFrench               string = "Pigmento"
	nameGerman               string = "Klecks"
	nameItalian              string = "Polpotto"
	nameJapanese             string = "スミダクン"
	nameLatinAmericanSpanish string = "Tintelio"
	nameKorean               string = "멍무리"
	nameRussian              string = ""
	nameSpanish              string = "Tintelio"
	nameSimplifiedChinese    string = ""
	nameTraditionalChinese   string = ""
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

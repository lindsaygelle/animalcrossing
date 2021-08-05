package cookie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cookie"
	nameCanadianFrench       string = "Cookie"
	nameDutch                string = "Cookie"
	nameFrench               string = "Cookie"
	nameGerman               string = "Rosi"
	nameItalian              string = "Briciola"
	nameJapanese             string = "ペリーヌ"
	nameLatinAmericanSpanish string = "Purita"
	nameKorean               string = "베리"
	nameRussian              string = "Куки"
	nameSpanish              string = "Purita"
	nameSimplifiedChinese    string = "珮琳"
	nameTraditionalChinese   string = "珮琳"
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

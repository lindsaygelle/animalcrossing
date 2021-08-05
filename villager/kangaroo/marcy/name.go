package marcy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Marcy"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "mon bébé"
	nameGerman               string = "kerlchen"
	nameItalian              string = "bebè"
	nameJapanese             string = "つきょ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "pimpollo"
	nameSimplifiedChinese    string = "哟"
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

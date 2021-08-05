package curlos

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Curlos"
	nameCanadianFrench       string = "Tonton"
	nameDutch                string = "Curlos"
	nameFrench               string = "Tonton"
	nameGerman               string = "Locke"
	nameItalian              string = "Boccolo"
	nameJapanese             string = "カルロス"
	nameLatinAmericanSpanish string = "Carnerio"
	nameKorean               string = "카를로스"
	nameRussian              string = "Керлос"
	nameSpanish              string = "Carnerio"
	nameSimplifiedChinese    string = "贾洛斯"
	nameTraditionalChinese   string = "賈洛斯"
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

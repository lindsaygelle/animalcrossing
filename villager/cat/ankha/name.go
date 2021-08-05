package ankha

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ankha"
	nameCanadianFrench       string = "Neferti"
	nameDutch                string = "Ankha"
	nameFrench               string = "Neferti"
	nameGerman               string = "Kleo"
	nameItalian              string = "Cleo"
	nameJapanese             string = "ナイル"
	nameLatinAmericanSpanish string = "Patri"
	nameKorean               string = "클레오"
	nameRussian              string = "Анка"
	nameSpanish              string = "Patri"
	nameSimplifiedChinese    string = "艳后"
	nameTraditionalChinese   string = "艷后"
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

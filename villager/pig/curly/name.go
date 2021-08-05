package curly

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Curly"
	nameCanadianFrench       string = "Tirbou"
	nameDutch                string = "Curly"
	nameFrench               string = "Tirbou"
	nameGerman               string = "Oink"
	nameItalian              string = "Ricciolo"
	nameJapanese             string = "ハムカツ"
	nameLatinAmericanSpanish string = "Rufueto"
	nameKorean               string = "햄까스"
	nameRussian              string = "Керли"
	nameSpanish              string = "Rufueto"
	nameSimplifiedChinese    string = "胜利"
	nameTraditionalChinese   string = "勝利"
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

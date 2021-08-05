package hamphrey

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hamphrey"
	nameCanadianFrench       string = "Charles"
	nameDutch                string = "Hamphrey"
	nameFrench               string = "Charles"
	nameGerman               string = "Heinrich"
	nameItalian              string = "Nerone"
	nameJapanese             string = "ハムジ"
	nameLatinAmericanSpanish string = "Arsenio"
	nameKorean               string = "햄쥐"
	nameRussian              string = "Хэмфри"
	nameSpanish              string = "Arsenio"
	nameSimplifiedChinese    string = "小仓"
	nameTraditionalChinese   string = "小倉"
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

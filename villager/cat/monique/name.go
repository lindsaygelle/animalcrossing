package monique

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Monique"
	nameCanadianFrench       string = "Monique"
	nameDutch                string = "Monique"
	nameFrench               string = "Monique"
	nameGerman               string = "Monique"
	nameItalian              string = "Marylin"
	nameJapanese             string = "ジェーン"
	nameLatinAmericanSpanish string = "Monique"
	nameKorean               string = "제인"
	nameRussian              string = "Моник"
	nameSpanish              string = "Monique"
	nameSimplifiedChinese    string = "珍珍"
	nameTraditionalChinese   string = "珍珍"
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

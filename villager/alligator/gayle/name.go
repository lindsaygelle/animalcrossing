package gayle

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gayle"
	nameCanadianFrench       string = "Odile"
	nameDutch                string = "Gayle"
	nameFrench               string = "Odile"
	nameGerman               string = "Rosa"
	nameItalian              string = "Codrilla"
	nameJapanese             string = "アリゲッティ"
	nameLatinAmericanSpanish string = "Boni"
	nameKorean               string = "앨리"
	nameRussian              string = "Гейл"
	nameSpanish              string = "Boni"
	nameSimplifiedChinese    string = "爱莉"
	nameTraditionalChinese   string = "愛莉"
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

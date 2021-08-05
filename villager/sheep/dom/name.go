package dom

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Dom"
	nameCanadianFrench       string = "Bouloche"
	nameDutch                string = "Donny"
	nameFrench               string = "Bouloche"
	nameGerman               string = "Dominik"
	nameItalian              string = "Ovilio"
	nameJapanese             string = "ちゃちゃまる"
	nameLatinAmericanSpanish string = "Fibrilio"
	nameKorean               string = "차둘"
	nameRussian              string = "Дом"
	nameSpanish              string = "Fibrilio"
	nameSimplifiedChinese    string = "茶茶丸"
	nameTraditionalChinese   string = "茶茶丸"
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

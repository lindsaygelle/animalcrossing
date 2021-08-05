package etoile

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Etoile"
	nameCanadianFrench       string = "Étoile"
	nameDutch                string = "Étoile"
	nameFrench               string = "Étoile"
	nameGerman               string = "Etoile"
	nameItalian              string = "Étoile"
	nameJapanese             string = "エトワール"
	nameLatinAmericanSpanish string = "Etoile"
	nameKorean               string = "에뜨와르"
	nameRussian              string = "Этуаль"
	nameSpanish              string = "Etoile"
	nameSimplifiedChinese    string = "彩星"
	nameTraditionalChinese   string = "彩星"
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

package lionel

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Lionel"
	nameCanadianFrench       string = "Léopold"
	nameDutch                string = "Lionel"
	nameFrench               string = "Léopold"
	nameGerman               string = "Lorenz"
	nameItalian              string = "Leonida"
	nameJapanese             string = "ライオネル"
	nameLatinAmericanSpanish string = "Lionel"
	nameKorean               string = "라이오넬"
	nameRussian              string = "Лайонел"
	nameSpanish              string = "Lionel"
	nameSimplifiedChinese    string = "赖恩睿"
	nameTraditionalChinese   string = "賴恩睿"
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

package broffina

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Broffina"
	nameCanadianFrench       string = "Jo"
	nameDutch                string = "Broffina"
	nameFrench               string = "Jo"
	nameGerman               string = "Elfriede"
	nameItalian              string = "Concetta"
	nameJapanese             string = "カサンドラ"
	nameLatinAmericanSpanish string = "Brunilda"
	nameKorean               string = "히킨"
	nameRussian              string = "Броффина"
	nameSpanish              string = "Brunilda"
	nameSimplifiedChinese    string = "凯西"
	nameTraditionalChinese   string = "凱西"
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

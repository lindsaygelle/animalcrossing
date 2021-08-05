package bunnie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bunnie"
	nameCanadianFrench       string = "Clara"
	nameDutch                string = "Bunnie"
	nameFrench               string = "Clara"
	nameGerman               string = "Mimmi"
	nameItalian              string = "Bonny"
	nameJapanese             string = "リリアン"
	nameLatinAmericanSpanish string = "Coni"
	nameKorean               string = "릴리안"
	nameRussian              string = "Банни"
	nameSpanish              string = "Coni"
	nameSimplifiedChinese    string = "莉莉安"
	nameTraditionalChinese   string = "莉莉安"
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

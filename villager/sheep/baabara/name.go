package baabara

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Baabara"
	nameCanadianFrench       string = "Bêêtty"
	nameDutch                string = "Baabara"
	nameFrench               string = "Bêêtty"
	nameGerman               string = "Babsi"
	nameItalian              string = "Bea"
	nameJapanese             string = "トロワ"
	nameLatinAmericanSpanish string = "Beelén"
	nameKorean               string = "트로와"
	nameRussian              string = "Бабара"
	nameSpanish              string = "Beelén"
	nameSimplifiedChinese    string = "华尔滋"
	nameTraditionalChinese   string = "華爾滋"
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

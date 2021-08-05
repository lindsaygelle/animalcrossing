package tangy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tangy"
	nameCanadianFrench       string = "Marine"
	nameDutch                string = "Tangy"
	nameFrench               string = "Marine"
	nameGerman               string = "Tanja"
	nameItalian              string = "Arancina"
	nameJapanese             string = "ヒャクパー"
	nameLatinAmericanSpanish string = "Tricia"
	nameKorean               string = "백프로"
	nameRussian              string = "Тэнджи"
	nameSpanish              string = "Tricia"
	nameSimplifiedChinese    string = "纯纯"
	nameTraditionalChinese   string = "純純"
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

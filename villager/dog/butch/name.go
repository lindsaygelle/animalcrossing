package butch

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Butch"
	nameCanadianFrench       string = "Avril"
	nameDutch                string = "Butch"
	nameFrench               string = "Avril"
	nameGerman               string = "Hasso"
	nameItalian              string = "Attila"
	nameJapanese             string = "ジョン"
	nameLatinAmericanSpanish string = "Bruno"
	nameKorean               string = "존"
	nameRussian              string = "Бутч"
	nameSpanish              string = "Bruno"
	nameSimplifiedChinese    string = "约翰"
	nameTraditionalChinese   string = "約翰"
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

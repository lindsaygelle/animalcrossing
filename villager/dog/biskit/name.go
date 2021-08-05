package biskit

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Biskit"
	nameCanadianFrench       string = "Crocket"
	nameDutch                string = "Biskit"
	nameFrench               string = "Crocket"
	nameGerman               string = "Keks"
	nameItalian              string = "Buffetto"
	nameJapanese             string = "ロビン"
	nameLatinAmericanSpanish string = "Amnesio"
	nameKorean               string = "로빈"
	nameRussian              string = "Бискит"
	nameSpanish              string = "Amnesio"
	nameSimplifiedChinese    string = "罗宾"
	nameTraditionalChinese   string = "羅賓"
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

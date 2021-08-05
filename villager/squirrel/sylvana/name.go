package sylvana

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sylvana"
	nameCanadianFrench       string = "Mounia"
	nameDutch                string = "Sylvana"
	nameFrench               string = "Mounia"
	nameGerman               string = "Maren"
	nameItalian              string = "Silvana"
	nameJapanese             string = "もんぺ"
	nameLatinAmericanSpanish string = "Silvana"
	nameKorean               string = "실바나"
	nameRussian              string = "Сильвана"
	nameSpanish              string = "Silvana"
	nameSimplifiedChinese    string = "孟珮"
	nameTraditionalChinese   string = "孟珮"
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

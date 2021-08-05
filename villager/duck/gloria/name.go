package gloria

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gloria"
	nameCanadianFrench       string = "Déborah"
	nameDutch                string = "Gloria"
	nameFrench               string = "Déborah"
	nameGerman               string = "Gustavia"
	nameItalian              string = "Pappy"
	nameJapanese             string = "スワンソン"
	nameLatinAmericanSpanish string = "Jarrea"
	nameKorean               string = "마릴린"
	nameRussian              string = "Глория"
	nameSpanish              string = "Jarrea"
	nameSimplifiedChinese    string = "素返真"
	nameTraditionalChinese   string = "素返真"
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

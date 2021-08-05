package tex

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tex"
	nameCanadianFrench       string = "Émilien"
	nameDutch                string = "Tex"
	nameFrench               string = "Émilien"
	nameGerman               string = "Matze"
	nameItalian              string = "Freddy"
	nameJapanese             string = "ボルト"
	nameLatinAmericanSpanish string = "Tomeo"
	nameKorean               string = "볼트"
	nameRussian              string = "Текс"
	nameSpanish              string = "Tomeo"
	nameSimplifiedChinese    string = "伏特"
	nameTraditionalChinese   string = "伏特"
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

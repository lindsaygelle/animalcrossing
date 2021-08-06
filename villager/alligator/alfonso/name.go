package alfonso

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Alfonso"
	nameCanadianFrench       string = "Alphonse"
	nameDutch                string = "Alfonso"
	nameFrench               string = "Alphonse"
	nameGerman               string = "Markus"
	nameItalian              string = "Alfonso"
	nameJapanese             string = "アルベルト"
	nameLatinAmericanSpanish string = "Kaimán"
	nameKorean               string = "알베르트"
	nameRussian              string = "Альфонсо"
	nameSpanish              string = "Kaimán"
	nameSimplifiedChinese    string = "阿泥"
	nameTraditionalChinese   string = "阿泥"
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

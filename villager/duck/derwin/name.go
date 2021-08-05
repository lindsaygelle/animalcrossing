package derwin

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Derwin"
	nameCanadianFrench       string = "Prof"
	nameDutch                string = "Derwin"
	nameFrench               string = "Prof"
	nameGerman               string = "Erwin"
	nameItalian              string = "Mike"
	nameJapanese             string = "ボン"
	nameLatinAmericanSpanish string = "Torcuato"
	nameKorean               string = "봉"
	nameRussian              string = "Дервин"
	nameSpanish              string = "Torcuato"
	nameSimplifiedChinese    string = "阿坊"
	nameTraditionalChinese   string = "阿坊"
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

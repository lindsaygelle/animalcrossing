package nan

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Nan"
	nameCanadianFrench       string = "Nana"
	nameDutch                string = "Nan"
	nameFrench               string = "Nana"
	nameGerman               string = "Zenobi"
	nameItalian              string = "Nan"
	nameJapanese             string = "スミ"
	nameLatinAmericanSpanish string = "Pécora"
	nameKorean               string = "순이"
	nameRussian              string = "Нань"
	nameSpanish              string = "Pécora"
	nameSimplifiedChinese    string = "墨墨"
	nameTraditionalChinese   string = "墨墨"
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

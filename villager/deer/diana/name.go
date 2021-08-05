package diana

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Diana"
	nameCanadianFrench       string = "Didi"
	nameDutch                string = "Diana"
	nameFrench               string = "Didi"
	nameGerman               string = "Vroni"
	nameItalian              string = "Diana"
	nameJapanese             string = "ナタリー"
	nameLatinAmericanSpanish string = "Bambina"
	nameKorean               string = "나탈리"
	nameRussian              string = "Диана"
	nameSpanish              string = "Bambina"
	nameSimplifiedChinese    string = "倪家莉"
	nameTraditionalChinese   string = "倪家莉"
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

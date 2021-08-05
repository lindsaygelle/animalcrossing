package bella

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bella"
	nameCanadianFrench       string = "Belle"
	nameDutch                string = "Bella"
	nameFrench               string = "Belle"
	nameGerman               string = "Susi"
	nameItalian              string = "Bella"
	nameJapanese             string = "イザベラ"
	nameLatinAmericanSpanish string = "Prity"
	nameKorean               string = "이자벨"
	nameRussian              string = "Белла"
	nameSpanish              string = "Prity"
	nameSimplifiedChinese    string = "贝拉"
	nameTraditionalChinese   string = "貝拉"
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

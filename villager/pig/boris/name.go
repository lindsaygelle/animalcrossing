package boris

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Boris"
	nameCanadianFrench       string = "Boris"
	nameDutch                string = "Boris"
	nameFrench               string = "Boris"
	nameGerman               string = "Bolle"
	nameItalian              string = "Boris"
	nameJapanese             string = "ダリー"
	nameLatinAmericanSpanish string = "Boris"
	nameKorean               string = "보리"
	nameRussian              string = "Борис"
	nameSpanish              string = "Boris"
	nameSimplifiedChinese    string = "达利"
	nameTraditionalChinese   string = "達利"
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

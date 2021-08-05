package velma

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Velma"
	nameCanadianFrench       string = "Véra"
	nameDutch                string = "Velma"
	nameFrench               string = "Véra"
	nameGerman               string = "Wilma"
	nameItalian              string = "Belarda"
	nameJapanese             string = "ピティエ"
	nameLatinAmericanSpanish string = "Erika"
	nameKorean               string = "벨마"
	nameRussian              string = "Велма"
	nameSpanish              string = "Erika"
	nameSimplifiedChinese    string = "班长"
	nameTraditionalChinese   string = "班長"
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

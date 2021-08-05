package cole

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cole"
	nameCanadianFrench       string = "Épicure"
	nameDutch                string = "Cole"
	nameFrench               string = "Épicure"
	nameGerman               string = "Karl"
	nameItalian              string = "Lello"
	nameJapanese             string = "アマミン"
	nameLatinAmericanSpanish string = "Nicolás"
	nameKorean               string = "아마민"
	nameRussian              string = "Коул"
	nameSpanish              string = "Nicolás"
	nameSimplifiedChinese    string = "阿甘"
	nameTraditionalChinese   string = "阿甘"
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

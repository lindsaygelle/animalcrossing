package axel

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Axel"
	nameCanadianFrench       string = "Axel"
	nameDutch                string = "Axel"
	nameFrench               string = "Axel"
	nameGerman               string = "Axel"
	nameItalian              string = "Sandro"
	nameJapanese             string = "エックスエル"
	nameLatinAmericanSpanish string = "Eustakio"
	nameKorean               string = "엑스엘리"
	nameRussian              string = "Аксель"
	nameSpanish              string = "Eustakio"
	nameSimplifiedChinese    string = "大大"
	nameTraditionalChinese   string = "大大"
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

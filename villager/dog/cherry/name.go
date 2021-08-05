package cherry

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cherry"
	nameCanadianFrench       string = "Anna"
	nameDutch                string = "Cherry"
	nameFrench               string = "Anna"
	nameGerman               string = "Bella"
	nameItalian              string = "Amarena"
	nameJapanese             string = "ハンナ"
	nameLatinAmericanSpanish string = "Luna"
	nameKorean               string = "한나"
	nameRussian              string = "Черри"
	nameSpanish              string = "Luna"
	nameSimplifiedChinese    string = "花娜"
	nameTraditionalChinese   string = "花娜"
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

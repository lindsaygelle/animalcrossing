package bluebear

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bluebear"
	nameCanadianFrench       string = "Myrtille"
	nameDutch                string = "Bluebear"
	nameFrench               string = "Myrtille"
	nameGerman               string = "Birte"
	nameItalian              string = "Azzurra"
	nameJapanese             string = "グルミン"
	nameLatinAmericanSpanish string = "Celeste"
	nameKorean               string = "글루민"
	nameRussian              string = "Блюбеар"
	nameSpanish              string = "Celeste"
	nameSimplifiedChinese    string = "娃娃"
	nameTraditionalChinese   string = "娃娃"
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

package tabby

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tabby"
	nameCanadianFrench       string = "Tigri"
	nameDutch                string = "Tabby"
	nameFrench               string = "Tigri"
	nameGerman               string = "Zita"
	nameItalian              string = "Lisca"
	nameJapanese             string = "トラこ"
	nameLatinAmericanSpanish string = "Liana"
	nameKorean               string = "호냥이"
	nameRussian              string = "Тэбби"
	nameSpanish              string = "Liana"
	nameSimplifiedChinese    string = "小虎"
	nameTraditionalChinese   string = "小虎"
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

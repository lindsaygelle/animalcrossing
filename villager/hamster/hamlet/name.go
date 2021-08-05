package hamlet

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hamlet"
	nameCanadianFrench       string = "Jojo"
	nameDutch                string = "Hamlet"
	nameFrench               string = "Jojo"
	nameGerman               string = "Hamid"
	nameItalian              string = "Amleto"
	nameJapanese             string = "ハムスケ"
	nameLatinAmericanSpanish string = "Bombo"
	nameKorean               string = "햄스틴"
	nameRussian              string = "Гамлет"
	nameSpanish              string = "Bombo"
	nameSimplifiedChinese    string = "哈姆"
	nameTraditionalChinese   string = "哈姆"
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

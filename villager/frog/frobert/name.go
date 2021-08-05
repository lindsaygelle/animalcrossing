package frobert

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Frobert"
	nameCanadianFrench       string = "Verbert"
	nameDutch                string = "Frobert"
	nameFrench               string = "Verbert"
	nameGerman               string = "Fritz"
	nameItalian              string = "Froberto"
	nameJapanese             string = "コージィ"
	nameLatinAmericanSpanish string = "Croaldo"
	nameKorean               string = "구리구리"
	nameRussian              string = "Фроберт"
	nameSpanish              string = "Croaldo"
	nameSimplifiedChinese    string = "江适"
	nameTraditionalChinese   string = "江適"
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

package quillson

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Quillson"
	nameCanadianFrench       string = "Narcisse"
	nameDutch                string = "Quillson"
	nameFrench               string = "Narcisse"
	nameGerman               string = "Quentin"
	nameItalian              string = "Verdonio"
	nameJapanese             string = "タックン"
	nameLatinAmericanSpanish string = "Cuálter"
	nameKorean               string = "덕근"
	nameRussian              string = "Квилсон"
	nameSpanish              string = "Cuálter"
	nameSimplifiedChinese    string = "何童"
	nameTraditionalChinese   string = "何童"
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

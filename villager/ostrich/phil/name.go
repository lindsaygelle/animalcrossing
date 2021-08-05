package phil

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Phil"
	nameCanadianFrench       string = "Phil"
	nameDutch                string = "Phil"
	nameFrench               string = "Phil"
	nameGerman               string = "Ingo"
	nameItalian              string = "Enzo"
	nameJapanese             string = "ケイン"
	nameLatinAmericanSpanish string = "Amalio"
	nameKorean               string = "케인"
	nameRussian              string = "Фил"
	nameSpanish              string = "Amalio"
	nameSimplifiedChinese    string = "凯恩"
	nameTraditionalChinese   string = "凱恩"
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

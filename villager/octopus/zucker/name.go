package zucker

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Zucker"
	nameCanadianFrench       string = "Marvin"
	nameDutch                string = "Zucker"
	nameFrench               string = "Marvin"
	nameGerman               string = "Ottokar"
	nameItalian              string = "Poldo"
	nameJapanese             string = "タコヤ"
	nameLatinAmericanSpanish string = "Paulino"
	nameKorean               string = "탁호"
	nameRussian              string = "Цукер"
	nameSpanish              string = "Paulino"
	nameSimplifiedChinese    string = "章丸丸"
	nameTraditionalChinese   string = "章丸丸"
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

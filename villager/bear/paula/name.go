package paula

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Paula"
	nameCanadianFrench       string = "Wendy"
	nameDutch                string = "Paula"
	nameFrench               string = "Wendy"
	nameGerman               string = "Paula"
	nameItalian              string = "Bruna"
	nameJapanese             string = "レイチェル"
	nameLatinAmericanSpanish string = "Lizi"
	nameKorean               string = "레이첼"
	nameRussian              string = "Паула"
	nameSpanish              string = "Lizi"
	nameSimplifiedChinese    string = "瑞秋"
	nameTraditionalChinese   string = "瑞秋"
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

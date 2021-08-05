package bones

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bones"
	nameCanadianFrench       string = "Nonos"
	nameDutch                string = "Bones"
	nameFrench               string = "Nonos"
	nameGerman               string = "Strolch"
	nameItalian              string = "Tobia"
	nameJapanese             string = "トミ"
	nameLatinAmericanSpanish string = "Cándido"
	nameKorean               string = "토미"
	nameRussian              string = "Боунс"
	nameSpanish              string = "Cándido"
	nameSimplifiedChinese    string = "阿富"
	nameTraditionalChinese   string = "阿富"
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

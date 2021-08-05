package tammi

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tammi"
	nameCanadianFrench       string = "Lili"
	nameDutch                string = "Tammi"
	nameFrench               string = "Lili"
	nameGerman               string = "Bonni"
	nameItalian              string = "Tammi"
	nameJapanese             string = "エイプリル"
	nameLatinAmericanSpanish string = "Tami"
	nameKorean               string = "에이프릴"
	nameRussian              string = "Тэмми"
	nameSpanish              string = "Tami"
	nameSimplifiedChinese    string = "四月"
	nameTraditionalChinese   string = "四月"
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

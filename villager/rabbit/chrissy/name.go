package chrissy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chrissy"
	nameCanadianFrench       string = "Kristine"
	nameDutch                string = "Chrissy"
	nameFrench               string = "Kristine"
	nameGerman               string = "Anne"
	nameItalian              string = "Natascha"
	nameJapanese             string = "クリスチーヌ"
	nameLatinAmericanSpanish string = "Lali"
	nameKorean               string = "크리스틴"
	nameRussian              string = "Крисси"
	nameSpanish              string = "Lali"
	nameSimplifiedChinese    string = "克莉琪"
	nameTraditionalChinese   string = "克莉琪"
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

package spike

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Spike"
	nameCanadianFrench       string = "Rhino"
	nameDutch                string = "Spike"
	nameFrench               string = "Rhino"
	nameGerman               string = "Atze"
	nameItalian              string = "Bruto"
	nameJapanese             string = "スクワット"
	nameLatinAmericanSpanish string = "Cornio"
	nameKorean               string = "스쿼트"
	nameRussian              string = "Спайк"
	nameSpanish              string = "Cornio"
	nameSimplifiedChinese    string = "阿蹲"
	nameTraditionalChinese   string = "阿蹲"
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

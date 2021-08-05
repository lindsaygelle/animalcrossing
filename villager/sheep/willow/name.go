package willow

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Willow"
	nameCanadianFrench       string = "Maï"
	nameDutch                string = "Willow"
	nameFrench               string = "Maï"
	nameGerman               string = "Natascha"
	nameItalian              string = "Belinda"
	nameJapanese             string = "マリー"
	nameLatinAmericanSpanish string = "Cuqui"
	nameKorean               string = "마리"
	nameRussian              string = "Уиллоу"
	nameSpanish              string = "Cuqui"
	nameSimplifiedChinese    string = "梅丽诺"
	nameTraditionalChinese   string = "梅麗諾"
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

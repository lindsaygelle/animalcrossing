package maple

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Maple"
	nameCanadianFrench       string = "Léa"
	nameDutch                string = "Maple"
	nameFrench               string = "Léa"
	nameGerman               string = "Mona"
	nameItalian              string = "Dulcinea"
	nameJapanese             string = "メープル"
	nameLatinAmericanSpanish string = "Dulce"
	nameKorean               string = "메이첼"
	nameRussian              string = "Мейпл"
	nameSpanish              string = "Dulce"
	nameSimplifiedChinese    string = "小枫"
	nameTraditionalChinese   string = "小楓"
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

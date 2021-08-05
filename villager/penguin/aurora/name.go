package aurora

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Aurora"
	nameCanadianFrench       string = "Aurore"
	nameDutch                string = "Aurora"
	nameFrench               string = "Aurore"
	nameGerman               string = "Sonja"
	nameItalian              string = "Aurora"
	nameJapanese             string = "オーロラ"
	nameLatinAmericanSpanish string = "Aurora"
	nameKorean               string = "오로라"
	nameRussian              string = "Аврора"
	nameSpanish              string = "Aurora"
	nameSimplifiedChinese    string = "欧若拉"
	nameTraditionalChinese   string = "歐若拉"
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

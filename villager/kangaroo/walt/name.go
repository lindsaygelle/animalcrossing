package walt

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Walt"
	nameCanadianFrench       string = "Walt"
	nameDutch                string = "Walt"
	nameFrench               string = "Walt"
	nameGerman               string = "Sinan"
	nameItalian              string = "Zompo"
	nameJapanese             string = "カンロク"
	nameLatinAmericanSpanish string = "Brinco"
	nameKorean               string = "관록"
	nameRussian              string = "Уолт"
	nameSpanish              string = "Brinco"
	nameSimplifiedChinese    string = "袋钢"
	nameTraditionalChinese   string = "袋鋼"
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

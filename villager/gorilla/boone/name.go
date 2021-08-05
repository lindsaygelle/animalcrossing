package boone

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Boone"
	nameCanadianFrench       string = "Babouin"
	nameDutch                string = "Boone"
	nameFrench               string = "Babouin"
	nameGerman               string = "Kong"
	nameItalian              string = "Babu"
	nameJapanese             string = "まんたろう"
	nameLatinAmericanSpanish string = "Babú"
	nameKorean               string = "만복이"
	nameRussian              string = "Буин"
	nameSpanish              string = "Babú"
	nameSimplifiedChinese    string = "万泰"
	nameTraditionalChinese   string = "萬泰"
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

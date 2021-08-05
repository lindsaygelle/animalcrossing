package whitney

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Whitney"
	nameCanadianFrench       string = "Blanche"
	nameDutch                string = "Whitney"
	nameFrench               string = "Blanche"
	nameGerman               string = "Lupa"
	nameItalian              string = "Bianca"
	nameJapanese             string = "ビアンカ"
	nameLatinAmericanSpanish string = "Lupe"
	nameKorean               string = "비앙카"
	nameRussian              string = "Уитни"
	nameSpanish              string = "Lupe"
	nameSimplifiedChinese    string = "毕安卡"
	nameTraditionalChinese   string = "畢安卡"
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

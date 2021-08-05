package julian

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Julian"
	nameCanadianFrench       string = "Lico"
	nameDutch                string = "Julian"
	nameFrench               string = "Lico"
	nameGerman               string = "Jimmy"
	nameItalian              string = "Giuliano"
	nameJapanese             string = "ジュリー"
	nameLatinAmericanSpanish string = "Azulino"
	nameKorean               string = "유니오"
	nameRussian              string = "Джулиан"
	nameSpanish              string = "Azulino"
	nameSimplifiedChinese    string = "朱黎"
	nameTraditionalChinese   string = "朱黎"
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

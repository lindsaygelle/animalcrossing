package rocco

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rocco"
	nameCanadianFrench       string = "Bebel"
	nameDutch                string = "Rocco"
	nameFrench               string = "Bebel"
	nameGerman               string = "Richi"
	nameItalian              string = "Rocco"
	nameJapanese             string = "ゴンザレス"
	nameLatinAmericanSpanish string = "Roco"
	nameKorean               string = "곤잘레스"
	nameRussian              string = "Рокко"
	nameSpanish              string = "Roco"
	nameSimplifiedChinese    string = "孔在来"
	nameTraditionalChinese   string = "孔在來"
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

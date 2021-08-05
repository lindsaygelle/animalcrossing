package lily

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Lily"
	nameCanadianFrench       string = "Raina"
	nameDutch                string = "Lily"
	nameFrench               string = "Raina"
	nameGerman               string = "Liliane"
	nameItalian              string = "Gigliola"
	nameJapanese             string = "レイニー"
	nameLatinAmericanSpanish string = "Lili"
	nameKorean               string = "레이니"
	nameRussian              string = "Лили"
	nameSpanish              string = "Lili"
	nameSimplifiedChinese    string = "雷妮"
	nameTraditionalChinese   string = "雷妮"
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

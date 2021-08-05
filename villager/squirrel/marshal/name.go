package marshal

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Marshal"
	nameCanadianFrench       string = "Mathéo"
	nameDutch                string = "Marshal"
	nameFrench               string = "Mathéo"
	nameGerman               string = "Huschke"
	nameItalian              string = "Scott"
	nameJapanese             string = "ジュン"
	nameLatinAmericanSpanish string = "Munchi"
	nameKorean               string = "쭈니"
	nameRussian              string = "Маршал"
	nameSpanish              string = "Munchi"
	nameSimplifiedChinese    string = "小润"
	nameTraditionalChinese   string = "小潤"
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

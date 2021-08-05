package klaus

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Klaus"
	nameCanadianFrench       string = "Klaus"
	nameDutch                string = "Klaus"
	nameFrench               string = "Klaus"
	nameGerman               string = "Grischa"
	nameItalian              string = "Maciste"
	nameJapanese             string = "クマロス"
	nameLatinAmericanSpanish string = "Gruñerto"
	nameKorean               string = "곰도로스"
	nameRussian              string = "Клаус"
	nameSpanish              string = "Gruñerto"
	nameSimplifiedChinese    string = "熊战士"
	nameTraditionalChinese   string = "熊戰士"
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

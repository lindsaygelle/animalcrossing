package buzz

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Buzz"
	nameCanadianFrench       string = "Phébus"
	nameDutch                string = "Buzz"
	nameFrench               string = "Phébus"
	nameGerman               string = "Balduin"
	nameItalian              string = "Golia"
	nameJapanese             string = "ひでよし"
	nameLatinAmericanSpanish string = "Nabar"
	nameKorean               string = "근엄"
	nameRussian              string = "Базз"
	nameSpanish              string = "Nabar"
	nameSimplifiedChinese    string = "继光"
	nameTraditionalChinese   string = "繼光"
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

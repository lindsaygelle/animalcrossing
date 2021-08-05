package sprocket

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sprocket"
	nameCanadianFrench       string = "Laflèche"
	nameDutch                string = "Sprocket"
	nameFrench               string = "Laflèche"
	nameGerman               string = "Lutz"
	nameItalian              string = "Frankie"
	nameJapanese             string = "ヘルツ"
	nameLatinAmericanSpanish string = "Ráfaga"
	nameKorean               string = "헤르츠"
	nameRussian              string = "Спрокет"
	nameSpanish              string = "Ráfaga"
	nameSimplifiedChinese    string = "鹤兹"
	nameTraditionalChinese   string = "鶴茲"
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

package alice

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Alice"
	nameCanadianFrench       string = "Alice"
	nameDutch                string = "Alice"
	nameFrench               string = "Alice"
	nameGerman               string = "Konny"
	nameItalian              string = "Alice"
	nameJapanese             string = "メルボルン"
	nameLatinAmericanSpanish string = "Zelanda"
	nameKorean               string = "멜버른"
	nameRussian              string = "Алиса"
	nameSpanish              string = "Zelanda"
	nameSimplifiedChinese    string = "莫儿"
	nameTraditionalChinese   string = "莫兒"
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

package henry

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Henry"
	nameCanadianFrench       string = "Henri"
	nameDutch                string = "Henry"
	nameFrench               string = "Henri"
	nameGerman               string = "Freddy"
	nameItalian              string = "Renato"
	nameJapanese             string = "ヘンリー"
	nameLatinAmericanSpanish string = "Saperto"
	nameKorean               string = "헨리"
	nameRussian              string = "Генри"
	nameSpanish              string = "Saperto"
	nameSimplifiedChinese    string = "亨利"
	nameTraditionalChinese   string = "亨利"
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

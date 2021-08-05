package felicity

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Felicity"
	nameCanadianFrench       string = "Maud"
	nameDutch                string = "Felicity"
	nameFrench               string = "Maud"
	nameGerman               string = "Minka"
	nameItalian              string = "Milly"
	nameJapanese             string = "みかっち"
	nameLatinAmericanSpanish string = "Micha"
	nameKorean               string = "예링"
	nameRussian              string = "Фелисити"
	nameSpanish              string = "Micha"
	nameSimplifiedChinese    string = "美佳"
	nameTraditionalChinese   string = "美佳"
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

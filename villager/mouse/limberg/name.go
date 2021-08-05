package limberg

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Limberg"
	nameCanadianFrench       string = "Gruyère"
	nameDutch                string = "Limberg"
	nameFrench               string = "Gruyère"
	nameGerman               string = "Rafael"
	nameItalian              string = "Rosico"
	nameJapanese             string = "らっきょ"
	nameLatinAmericanSpanish string = "Camember"
	nameKorean               string = "단무지"
	nameRussian              string = "Лимберг"
	nameSpanish              string = "Camember"
	nameSimplifiedChinese    string = "韭菜"
	nameTraditionalChinese   string = "韭菜"
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

package tom

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tom"
	nameCanadianFrench       string = "Tom"
	nameDutch                string = "Tom"
	nameFrench               string = "Tom"
	nameGerman               string = "Timo"
	nameItalian              string = "Isidoro"
	nameJapanese             string = "バンタム"
	nameLatinAmericanSpanish string = "Zapirón"
	nameKorean               string = "밴덤"
	nameRussian              string = "Том"
	nameSpanish              string = "Zapirón"
	nameSimplifiedChinese    string = "阿邦"
	nameTraditionalChinese   string = "阿邦"
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

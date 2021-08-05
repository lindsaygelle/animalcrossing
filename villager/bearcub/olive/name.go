package olive

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Olive"
	nameCanadianFrench       string = "Grisa"
	nameDutch                string = "Olive"
	nameFrench               string = "Grisa"
	nameGerman               string = "Linda"
	nameItalian              string = "Olivia"
	nameJapanese             string = "ピッコロ"
	nameLatinAmericanSpanish string = "Osalina"
	nameKorean               string = "올리브"
	nameRussian              string = "Олив"
	nameSpanish              string = "Osalina"
	nameSimplifiedChinese    string = "毕克萝"
	nameTraditionalChinese   string = "畢克蘿"
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

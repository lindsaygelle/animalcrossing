package pompom

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pompom"
	nameCanadianFrench       string = "Pompon"
	nameDutch                string = "Pompom"
	nameFrench               string = "Pompon"
	nameGerman               string = "Erika"
	nameItalian              string = "Erica"
	nameJapanese             string = "のりっぺ"
	nameLatinAmericanSpanish string = "Flopi"
	nameKorean               string = "주디"
	nameRussian              string = "Помпом"
	nameSpanish              string = "Flopi"
	nameSimplifiedChinese    string = "海苔裴"
	nameTraditionalChinese   string = "海苔裴"
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

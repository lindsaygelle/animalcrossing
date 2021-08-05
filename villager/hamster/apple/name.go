package apple

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Apple"
	nameCanadianFrench       string = "Esther"
	nameDutch                string = "Apple"
	nameFrench               string = "Esther"
	nameGerman               string = "Jessi"
	nameItalian              string = "Cicci"
	nameJapanese             string = "アップル"
	nameLatinAmericanSpanish string = "Rosi"
	nameKorean               string = "애플"
	nameRussian              string = "Эпл"
	nameSpanish              string = "Rosi"
	nameSimplifiedChinese    string = "苹果"
	nameTraditionalChinese   string = "蘋果"
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

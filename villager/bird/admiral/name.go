package admiral

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Admiral"
	nameCanadianFrench       string = "Maréchal"
	nameDutch                string = "Admiral"
	nameFrench               string = "Maréchal"
	nameGerman               string = "Ansgar"
	nameItalian              string = "Adolfo"
	nameJapanese             string = "イッテツ"
	nameLatinAmericanSpanish string = "Avutardo"
	nameKorean               string = "일섭"
	nameRussian              string = "Адмирал"
	nameSpanish              string = "Avutardo"
	nameSimplifiedChinese    string = "李彻"
	nameTraditionalChinese   string = "李徹"
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

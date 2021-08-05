package ellie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ellie"
	nameCanadianFrench       string = "Ella"
	nameDutch                string = "Ellie"
	nameFrench               string = "Ella"
	nameGerman               string = "Elfi"
	nameItalian              string = "Elly"
	nameJapanese             string = "エクレア"
	nameLatinAmericanSpanish string = "Eli"
	nameKorean               string = "에끌레르"
	nameRussian              string = "Элла"
	nameSpanish              string = "Eli"
	nameSimplifiedChinese    string = "泡芙"
	nameTraditionalChinese   string = "泡芙"
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

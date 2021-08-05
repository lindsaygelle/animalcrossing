package boots

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Boots"
	nameCanadianFrench       string = "Croko"
	nameDutch                string = "Boots"
	nameFrench               string = "Croko"
	nameGerman               string = "Tilmann"
	nameItalian              string = "Crocco"
	nameJapanese             string = "ホウサク"
	nameLatinAmericanSpanish string = "Braulio"
	nameKorean               string = "풍작"
	nameRussian              string = "Бутс"
	nameSpanish              string = "Braulio"
	nameSimplifiedChinese    string = "丰年"
	nameTraditionalChinese   string = "豐年"
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

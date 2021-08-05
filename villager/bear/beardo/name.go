package beardo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Beardo"
	nameCanadianFrench       string = "Eustache"
	nameDutch                string = "Beardo"
	nameFrench               string = "Eustache"
	nameGerman               string = "Berthold"
	nameItalian              string = "Barnaba"
	nameJapanese             string = "ベアード"
	nameLatinAmericanSpanish string = "Bernabé"
	nameKorean               string = "베어드"
	nameRussian              string = "Беардо"
	nameSpanish              string = "Bernabé"
	nameSimplifiedChinese    string = "熊大叔"
	nameTraditionalChinese   string = "熊大叔"
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

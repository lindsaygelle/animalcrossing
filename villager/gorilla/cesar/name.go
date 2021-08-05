package cesar

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cesar"
	nameCanadianFrench       string = "César"
	nameDutch                string = "Cesar"
	nameFrench               string = "César"
	nameGerman               string = "Alfredo"
	nameItalian              string = "Cesare"
	nameJapanese             string = "アラン"
	nameLatinAmericanSpanish string = "César"
	nameKorean               string = "앨런"
	nameRussian              string = "Цезарь"
	nameSpanish              string = "César"
	nameSimplifiedChinese    string = "阿朗"
	nameTraditionalChinese   string = "阿朗"
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

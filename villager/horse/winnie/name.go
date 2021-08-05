package winnie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Winnie"
	nameCanadianFrench       string = "Anne"
	nameDutch                string = "Winnie"
	nameFrench               string = "Anne"
	nameGerman               string = "Walli"
	nameItalian              string = "Clara"
	nameJapanese             string = "マキバスター"
	nameLatinAmericanSpanish string = "Soonia"
	nameKorean               string = "카로틴"
	nameRussian              string = "Винни"
	nameSpanish              string = "Soonia"
	nameSimplifiedChinese    string = "马星星"
	nameTraditionalChinese   string = "馬星星"
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

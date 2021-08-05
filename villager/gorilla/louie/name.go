package louie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Louie"
	nameCanadianFrench       string = "Louis"
	nameDutch                string = "Louie"
	nameFrench               string = "Louis"
	nameGerman               string = "Ludwig"
	nameItalian              string = "Lou"
	nameJapanese             string = "マッスル"
	nameLatinAmericanSpanish string = "Lou"
	nameKorean               string = "머슬"
	nameRussian              string = "Луи"
	nameSpanish              string = "Lou"
	nameSimplifiedChinese    string = "壮壮"
	nameTraditionalChinese   string = "壯壯"
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

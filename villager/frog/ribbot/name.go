package ribbot

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ribbot"
	nameCanadianFrench       string = "Crabot"
	nameDutch                string = "Ribbot"
	nameFrench               string = "Crabot"
	nameGerman               string = "Robbi"
	nameItalian              string = "Ranabot"
	nameJapanese             string = "ガチャ"
	nameLatinAmericanSpanish string = "Ranobot"
	nameKorean               string = "철컥"
	nameRussian              string = "Риббот"
	nameSpanish              string = "Ranobot"
	nameSimplifiedChinese    string = "锵锵"
	nameTraditionalChinese   string = "鏘鏘"
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

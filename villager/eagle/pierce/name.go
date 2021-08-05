package pierce

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pierce"
	nameCanadianFrench       string = "Napoléon"
	nameDutch                string = "Pierce"
	nameFrench               string = "Napoléon"
	nameGerman               string = "Adrian"
	nameItalian              string = "Attilio"
	nameJapanese             string = "セバスチャン"
	nameLatinAmericanSpanish string = "Hugo"
	nameKorean               string = "세바스찬"
	nameRussian              string = "Пирс"
	nameSpanish              string = "Hugo"
	nameSimplifiedChinese    string = "谢博强"
	nameTraditionalChinese   string = "謝博強"
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

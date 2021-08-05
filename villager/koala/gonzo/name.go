package gonzo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gonzo"
	nameCanadianFrench       string = "Gonzo"
	nameDutch                string = "Gonzo"
	nameFrench               string = "Gonzo"
	nameGerman               string = "Heribert"
	nameItalian              string = "Gonzo"
	nameJapanese             string = "ゴンゾー"
	nameLatinAmericanSpanish string = "Bronco"
	nameKorean               string = "근성"
	nameRussian              string = "Гонзо"
	nameSpanish              string = "Bronco"
	nameSimplifiedChinese    string = "钢锁"
	nameTraditionalChinese   string = "鋼鎖"
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

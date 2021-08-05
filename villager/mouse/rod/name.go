package rod

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rod"
	nameCanadianFrench       string = "Marcel"
	nameDutch                string = "Rod"
	nameFrench               string = "Marcel"
	nameGerman               string = "Manni"
	nameItalian              string = "Riccardo"
	nameJapanese             string = "ジャン"
	nameLatinAmericanSpanish string = "Rodi"
	nameKorean               string = "쟝"
	nameRussian              string = "Род"
	nameSpanish              string = "Rodi"
	nameSimplifiedChinese    string = "阿将"
	nameTraditionalChinese   string = "阿將"
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

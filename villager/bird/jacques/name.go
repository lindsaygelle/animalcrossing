package jacques

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Jacques"
	nameCanadianFrench       string = "Jacky"
	nameDutch                string = "Jacques"
	nameFrench               string = "Jacky"
	nameGerman               string = "Pierre"
	nameItalian              string = "Zampiero"
	nameJapanese             string = "ジョッキー"
	nameLatinAmericanSpanish string = "Gorrelmo"
	nameKorean               string = "쪼끼"
	nameRussian              string = "Жак"
	nameSpanish              string = "Gorrelmo"
	nameSimplifiedChinese    string = "赵奇"
	nameTraditionalChinese   string = "趙奇"
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

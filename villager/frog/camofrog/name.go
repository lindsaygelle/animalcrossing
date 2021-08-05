package camofrog

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Camofrog"
	nameCanadianFrench       string = "Milos"
	nameDutch                string = "Camofrog"
	nameFrench               string = "Milos"
	nameGerman               string = "Tarno"
	nameItalian              string = "Crambo"
	nameJapanese             string = "フルメタル"
	nameLatinAmericanSpanish string = "Comando"
	nameKorean               string = "충성"
	nameRussian              string = "Камофрог"
	nameSpanish              string = "Comando"
	nameSimplifiedChinese    string = "迷仔"
	nameTraditionalChinese   string = "迷仔"
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

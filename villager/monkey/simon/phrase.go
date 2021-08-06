package simon

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "zzzook"
	phraseCanadianFrench       string = "zzziik"
	phraseDutch                string = "zzzoeh oeh"
	phraseFrench               string = "zzziik"
	phraseGerman               string = "bannabanna"
	phraseItalian              string = "zzzook"
	phraseJapanese             string = "でござる"
	phraseLatinAmericanSpanish string = "chimpa"
	phraseKorean               string = "하옵니다"
	phraseRussian              string = "дружок"
	phraseSpanish              string = "chimpa"
	phraseSimplifiedChinese    string = "猿猿"
	phraseTraditionalChinese   string = "猿猿"
)

var (
	phrase = map[language.Tag]string{
		language.AmericanEnglish:      phraseAmericanEnglish,
		language.CanadianFrench:       phraseCanadianFrench,
		language.Dutch:                phraseDutch,
		language.French:               phraseFrench,
		language.German:               phraseGerman,
		language.Italian:              phraseItalian,
		language.Japanese:             phraseJapanese,
		language.Korean:               phraseKorean,
		language.LatinAmericanSpanish: phraseLatinAmericanSpanish,
		language.Russian:              phraseRussian,
		language.Spanish:              phraseSpanish,
		language.SimplifiedChinese:    phraseSimplifiedChinese,
		language.TraditionalChinese:   phraseTraditionalChinese}
)

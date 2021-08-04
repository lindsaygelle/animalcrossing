package alfonso

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish    string = "it⁠'⁠sa me"
	phraseFrench             string = "moi moi"
	phraseGerman             string = "ahhhrg"
	phraseItalian            string = "proprio"
	phraseJapanese           string = "だワニ"
	phraseKorean             string = "나아거"
	phraseRussian            string = "это я"
	phraseSpanish            string = "ñam-ñam"
	phraseSimplifiedChinese  string = "鳄泥"
	phraseTraditionalChinese string = "鱷泥"
)

var (
	phrase = map[language.Tag]string{
		language.AmericanEnglish:    phraseAmericanEnglish,
		language.French:             phraseFrench,
		language.German:             phraseGerman,
		language.Italian:            phraseItalian,
		language.Japanese:           phraseJapanese,
		language.Korean:             phraseKorean,
		language.Russian:            phraseRussian,
		language.Spanish:            phraseSpanish,
		language.SimplifiedChinese:  phraseSimplifiedChinese,
		language.TraditionalChinese: phraseTraditionalChinese}
)

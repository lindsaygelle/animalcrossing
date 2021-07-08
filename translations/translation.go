package translations

type Translation interface {
	ChineseSimplified() Language
	ChineseTraditional() Language
	Dutch() Language
	German() Language
	Italian() Language
	Japanese() Language
	Korean() Language
	Russian() Language
	Spanish() Language
}

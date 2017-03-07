package talib4g

type Rule interface {
	And(Rule) Rule
	Or(Rule) Rule
	Xor(Rule) Rule
	Negation() Rule
	IsSatisfied(index int, record TradingRecord) bool
}

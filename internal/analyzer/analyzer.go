package analyzer

type Analyzer interface {
	analyze(code []string)
}

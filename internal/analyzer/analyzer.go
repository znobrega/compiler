package analyzer

type Analyzer interface {
	Analyze(code []string) error
}

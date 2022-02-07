package gex

var (
	_                 = MaxLines
	_ option          = (*optionMax)(nil)
	_ collectorOption = (*optionMax)(nil)
)

type optionMax struct {
	max int
}

// MaxLines 配置最大收集行数
func MaxLines(max int) *optionMax {
	return &optionMax{
		max: max,
	}
}

func (ml *optionMax) apply(options *options) {
	options.max = ml.max
}

func (ml *optionMax) applyCollector(options *collectorOptions) {
	options.max = ml.max
}
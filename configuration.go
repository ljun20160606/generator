package generator

type Configurator func(gen *Generator)

type ConfiguratorComponent []Configurator

func (g ConfiguratorComponent) Run(gen *Generator) {
	for i := range g {
		g[i](gen)
	}
}

func WithSpecifiedTables(tables []string) Configurator {
	return func(gen *Generator) {
		gen.specifiedTables = tables
	}
}

package main

type Foo struct {
	verbosity int
}

// function returning its inverse
type option func(f *Foo) option

func Verbosity(v int) option {
	return func(f *Foo) option {
		previous := f.verbosity
		f.verbosity = v
		return Verbosity(previous)
	}
}

// Option sets the options, returning an option to restore the last args previous value.
func (f *Foo) Option(opts ...option) (previous option) {
	for _, opt := range opts {
		previous = opt(f)
	}
	return previous
}

func main() {
	f := Foo{}
	// set verbosity higher
	prevVerbosity := f.Option(Verbosity(3))
	// return to previous value after using it
	defer f.Option(prevVerbosity)
	// use at higher verbosity
}

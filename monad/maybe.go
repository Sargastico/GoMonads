package monad

type Maybe struct {
	value interface{}
	ok    bool
}

/*

The monad laws holds:

	Left identity: Return(v).Bind(f) is equivalent to f(v)
	Right identity: m.Bind(Return) is equivalent to m
	Associativity: m.Bind(f).Bind(g) is equivalent to m.Bind(func(x) { return f(x).Bind(g) })

*/

func Return(v interface{}) Maybe {
	return Maybe{value: v, ok: true}
}

func (m Maybe) Bind(f func(interface{}) Maybe) Maybe {
	if !m.ok {
		return Maybe{}
	}
	return f(m.value)
}

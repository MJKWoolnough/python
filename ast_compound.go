package python

var compounds = [...]string{"if", "while", "for", "try", "with", "func", "class", "async", "def"}

type CompoundStatement struct{}

func (c *CompoundStatement) parser(_ *pyParser) error {
	return nil
}

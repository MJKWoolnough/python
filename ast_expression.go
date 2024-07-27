package python

type Primary struct{}

func (p *Primary) parse(_ *pyParser) error {
	return nil
}

func (p *Primary) IsIdentifier() bool {
	return false
}

type ExpressionList struct{}

func (e *ExpressionList) parse(_ *pyParser) error {
	return nil
}

type SliceList struct{}

func (s *SliceList) parse(_ *pyParser) error {
	return nil
}

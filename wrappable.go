package python

// ConditionalWrappable represents the types that can be wrapped with
// WrapConditional and unwrapped with UnwrapConditional.
type ConditionalWrappable interface {
	Type
	conditionalWrapabble()
}

func (ConditionalExpression) conditionalWrapabble() {}

func (OrTest) conditionalWrapabble() {}

func (AndTest) conditionalWrapabble() {}

func (NotTest) conditionalWrapabble() {}

func (Comparison) conditionalWrapabble() {}

func (OrExpression) conditionalWrapabble() {}

func (XorExpression) conditionalWrapabble() {}

func (AndExpression) conditionalWrapabble() {}

func (ShiftExpression) conditionalWrapabble() {}

func (AddExpression) conditionalWrapabble() {}

func (MultiplyExpression) conditionalWrapabble() {}

func (UnaryExpression) conditionalWrapabble() {}

func (PowerExpression) conditionalWrapabble() {}

func (PrimaryExpression) conditionalWrapabble() {}

func (Atom) conditionalWrapabble() {}

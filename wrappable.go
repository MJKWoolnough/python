package python

type ConditionalWrappable interface {
	conditionalWrapabble()
}

func (ConditionalExpression) conditionalWrapabble()

func (OrTest) conditionalWrapabble()

func (AndTest) conditionalWrapabble()

func (NotTest) conditionalWrapabble()

func (Comparison) conditionalWrapabble()

func (OrExpression) conditionalWrapabble()

func (XorExpression) conditionalWrapabble()

func (AndExpression) conditionalWrapabble()

func (ShiftExpression) conditionalWrapabble()

func (AddExpression) conditionalWrapabble()

func (MultiplyExpression) conditionalWrapabble()

func (UnaryExpression) conditionalWrapabble()

func (PowerExpression) conditionalWrapabble()

func (PrimaryExpression) conditionalWrapabble()

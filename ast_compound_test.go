package python

import "testing"

func TestTypeParam(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = TypeParam{
				Identifier: &tk[0],
				Tokens:     tk[:1],
			}
		}},
		{`a:b`, func(t *test, tk Tokens) { // 2
			t.Output = TypeParam{
				Identifier: &tk[0],
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a : b`, func(t *test, tk Tokens) { // 3
			t.Output = TypeParam{
				Identifier: &tk[0],
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`*a`, func(t *test, tk Tokens) { // 4
			t.Output = TypeParam{
				Type:       TypeParamVar,
				Identifier: &tk[1],
				Tokens:     tk[:2],
			}
		}},
		{`* a`, func(t *test, tk Tokens) { // 5
			t.Output = TypeParam{
				Type:       TypeParamVar,
				Identifier: &tk[2],
				Tokens:     tk[:3],
			}
		}},
		{`**a`, func(t *test, tk Tokens) { // 6
			t.Output = TypeParam{
				Type:       TypeParamVarTuple,
				Identifier: &tk[1],
				Tokens:     tk[:2],
			}
		}},
		{`** a`, func(t *test, tk Tokens) { // 7
			t.Output = TypeParam{
				Type:       TypeParamVarTuple,
				Identifier: &tk[2],
				Tokens:     tk[:3],
			}
		}},
	}, func(t *test) (Type, error) {
		var tp TypeParam

		err := tp.parse(t.Tokens)

		return tp, err
	})
}
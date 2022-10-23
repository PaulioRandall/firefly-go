package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func newAud(given ...token.Token) auditor {
	lr := inout.NewListReader[token.Token](given)
	br := inout.NewBufReader[token.Token](lr)
	pr := inout.NewPosReader[token.Token](br)
	return auditor{r: pr}
}

func tok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func lit(tk token.Token) ast.Literal {
	return asttest.Literal(tk)
}

func lits(tks ...token.Token) []ast.Expr {
	var nodes []ast.Expr

	for _, tk := range tks {
		nodes = append(nodes, asttest.Literal(tk))
	}

	return nodes
}

func vars(tks ...token.Token) []ast.Variable {
	return asttest.Variables(tks...)
}

func binOp(left ast.Expr, op token.Token, right ast.Expr) ast.Expr {
	return asttest.BinaryOperation(left, op, right)
}

func exprs(tks ...token.Token) []ast.Expr {
	return asttest.Expressions(tks...)
}

func assStmt(left []ast.Variable, op token.Token, right []ast.Expr) ast.Assign {
	return asttest.Assign(
		left,
		op,
		asttest.ExprSet(right...),
	)
}

func whenStmt(
	keyword token.Token,
	subject ast.Expr,
	cases []ast.WhenCase,
	end token.Token,
) ast.When {
	return asttest.When(keyword, subject, cases, end)
}

func whenCase(
	condition ast.Expr,
	body ast.Stmt,
) ast.WhenCase {
	return asttest.WhenCase(condition, body)
}

func whenCases(cases ...ast.WhenCase) []ast.WhenCase {
	return cases
}

func stmts(stmts ...ast.Stmt) []ast.Stmt {
	return stmts
}

func ifStmt(
	keyword token.Token,
	condition ast.Expr,
	body []ast.Stmt,
	end token.Token,
) ast.If {
	return asttest.If(keyword, condition, body, end)
}

func doParseTest(t *testing.T, given []token.Token, exp ...ast.Node) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.Nil(t, e, "%s", debug.String(e))
	require.Equal(t, exp, w.List())
}

func doErrorTest(t *testing.T, given []token.Token, exp error) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.True(t, err.Is(e, exp), "Want error %q but got %s", exp, debug.String(e))
}

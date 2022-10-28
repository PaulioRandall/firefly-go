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

func listExpr(opener token.Token, values []ast.Expr, closer token.Token) ast.List {
	return asttest.List(opener, values, closer)
}

func mapExpr(opener token.Token, entries []ast.MapEntry, closer token.Token) ast.Map {
	return asttest.Map(opener, entries, closer)
}

func mapEntries(entries ...ast.MapEntry) []ast.MapEntry {
	return entries
}

func mapEntry(key token.Token, value token.Token) ast.MapEntry {
	return asttest.MapEntry(asttest.ExprFor(key), asttest.ExprFor(value))
}

func varExpr(tk token.Token) ast.Variable {
	return asttest.Variable(tk)
}

func vars(tks ...token.Token) []ast.Variable {
	return asttest.Variables(tks...)
}

func binOp(left ast.Expr, op token.Token, right ast.Expr) ast.Expr {
	return asttest.BinaryOperation(left, op, right)
}

func is(keyword token.Token, expr ast.Expr) ast.Expr {
	return asttest.Is(keyword, expr)
}

func exprs(tks ...token.Token) []ast.Expr {
	return asttest.Expressions(tks...)
}

func assStmt(left []ast.Variable, op token.Token, right []ast.Expr) ast.Assign {
	return asttest.Assign(
		asttest.SeriesOfVar(left...),
		op,
		asttest.SeriesOfExpr(right...),
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

func watchStmt(
	keyword token.Token,
	variable ast.Variable,
	body []ast.Stmt,
	end token.Token,
) ast.Watch {
	return asttest.Watch(keyword, variable, body, end)
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

func forStmt(
	keyword token.Token,
	initialiser ast.Stmt,
	condition ast.Expr,
	advancement ast.Stmt,
	body []ast.Stmt,
	end token.Token,
) ast.For {
	return asttest.For(keyword, initialiser, condition, advancement, body, end)
}

func ForEach(
	keyword token.Token,
	index ast.Variable,
	item ast.Variable,
	list ast.Expr,
	body []ast.Stmt,
	end token.Token,
) ast.ForEach {
	return asttest.ForEach(keyword, index, item, list, body, end)
}

func doParseTest(t *testing.T, given []token.Token, exp ...ast.Node) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.Nil(t, e, "%s", debug.String(e))
	require.Equal(t, exp, w.List(), debug.String(w.List()))
}

func doErrorTest(t *testing.T, given []token.Token, exps ...error) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	act := Parse(r, w)

	for _, exp := range exps {
		requireIsError(t, exp, act)
	}
}

func requireIsError(t *testing.T, exp, act error) {
	exists := err.Is(act, exp)
	require.True(t, exists,
		"Want error %v but got %s", exp, debug.String(act))
}

package asttest

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func Literal(tk token.Token) ast.Literal {
	return ast.Literal{
		Token: tk,
	}
}

func Variable(id token.Token) ast.Variable {
	return ast.Variable{
		Identifier: id,
	}
}

func List(opener token.Token, values []ast.Expr, closer token.Token) ast.List {
	return ast.List{
		Opener: opener,
		Values: values,
		Closer: closer,
	}
}

func Map(opener token.Token, entries []ast.MapEntry, closer token.Token) ast.Map {
	return ast.Map{
		Opener:  opener,
		Entries: entries,
		Closer:  closer,
	}
}

func MapEntry(key, value ast.Expr) ast.MapEntry {
	return ast.MapEntry{
		Key:   key,
		Value: value,
	}
}

func BinaryOperation(
	left ast.Expr,
	op token.Token,
	right ast.Expr,
) ast.BinaryOperation {
	return ast.BinaryOperation{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}

func Assign(left ast.SeriesOfVar, op token.Token, right ast.Proc) ast.Assign {
	return ast.Assign{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}

func SeriesOfVar(nodes ...ast.Variable) ast.SeriesOfVar {
	return ast.SeriesOfVar{
		Nodes: nodes,
	}
}

func SeriesOfExpr(nodes ...ast.Expr) ast.SeriesOfExpr {
	return ast.SeriesOfExpr{
		Nodes: nodes,
	}
}

func If(
	keyword token.Token,
	condition ast.Expr,
	body []ast.Stmt,
	end token.Token,
) ast.If {
	return ast.If{
		Keyword:   keyword,
		Condition: condition,
		Body:      body,
		End:       end,
	}
}

func For(
	keyword token.Token,
	initialiser ast.Stmt,
	condition ast.Expr,
	advancement ast.Stmt,
	body []ast.Stmt,
	end token.Token,
) ast.For {
	return ast.For{
		Keyword:     keyword,
		Initialiser: initialiser,
		Condition:   condition,
		Advancement: advancement,
		Body:        body,
		End:         end,
	}
}

func ForEach(
	keyword token.Token,
	vars ast.SeriesOfVar,
	iterable ast.Expr,
	body []ast.Stmt,
	end token.Token,
) ast.ForEach {
	return ast.ForEach{
		Keyword:  keyword,
		Vars:     vars,
		Iterable: iterable,
		Body:     body,
		End:      end,
	}
}

func When(
	keyword token.Token,
	subject ast.Expr,
	cases []ast.WhenCase,
	end token.Token,
) ast.When {
	return ast.When{
		Keyword: keyword,
		Subject: subject,
		Cases:   cases,
		End:     end,
	}
}

func WhenCase(
	condition ast.Expr,
	stmt ast.Stmt,
) ast.WhenCase {
	return ast.WhenCase{
		Condition: condition,
		Statement: stmt,
	}
}

func Is(
	keyword token.Token,
	expr ast.Expr,
) ast.Is {
	return ast.Is{
		Keyword: keyword,
		Expr:    expr,
	}
}

func Watch(
	keyword token.Token,
	variable ast.Variable,
	body []ast.Stmt,
	end token.Token,
) ast.Watch {
	return ast.Watch{
		Keyword:  keyword,
		Variable: variable,
		Body:     body,
		End:      end,
	}
}

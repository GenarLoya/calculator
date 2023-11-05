package token_manager

import (
	"errors"
	"genarold/calculator/utils"
	"strconv"
)

type UnitType int

const (
	Number UnitType = iota
	Operator
)

type Ps struct {
	Pe int
	Pp int
}

type Unit struct {
	Unit string
	Ps   Ps
	Is   UnitType
}

var pppeTable = []struct {
	token string
	pe    int
	pp    int
}{
	{
		token: "+",
		pe:    1,
		pp:    1,
	},
	{
		token: "-",
		pe:    1,
		pp:    1,
	},
	{
		token: "*",
		pe:    2,
		pp:    2,
	},
	{
		token: "/",
		pe:    2,
		pp:    2,
	},
	{
		token: "^",
		pe:    3,
		pp:    4,
	},
	{
		token: "(",
		pe:    5,
		pp:    0,
	},
	{
		token: ")",
		pe:    -1,
		pp:    -1,
	},
}

func unitTypeResolver(token string) UnitType {
	_, err := strconv.Atoi(token)

	if err == nil {
		return Number
	}

	return Operator
}

func pppeResolver(token string) Ps {
	for _, t := range pppeTable {
		if t.token == token {
			return Ps{
				Pe: t.pe,
				Pp: t.pp,
			}
		}
	}

	return Ps{
		Pe: 0,
		Pp: 0,
	}
}

func NewUnit(token string) *Unit {
	return &Unit{
		Unit: token,
		Ps:   pppeResolver(token),
		Is:   unitTypeResolver(token),
	}
}

func TokenManager(text string) ([]*Unit, error) {
	tokens := utils.Tokenizer(text)

	if !utils.TokenizerValidator(text) {
		return nil, errors.New("invalid expression")
	}

	units := make([]*Unit, len(tokens))

	for i, token := range tokens {
		units[i] = NewUnit(token)
	}

	return units, nil
}

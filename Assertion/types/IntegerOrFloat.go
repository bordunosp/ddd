package types

import "golang.org/x/exp/constraints"

type IntegerOrFloat interface {
	constraints.Float | constraints.Integer
}

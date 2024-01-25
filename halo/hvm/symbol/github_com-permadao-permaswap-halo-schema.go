// Code generated by 'yaegi extract github.com/permadao/permaswap/halo/schema'. DO NOT EDIT.

package symbol

import (
	"github.com/permadao/permaswap/halo/schema"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["github.com/permadao/permaswap/halo/schema/schema"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrInvalidGenesisBalance":     reflect.ValueOf(&schema.ErrInvalidGenesisBalance).Elem(),
		"ErrInvalidGenesisTotalSupply": reflect.ValueOf(&schema.ErrInvalidGenesisTotalSupply).Elem(),
		"ErrInvalidGenesisTx":          reflect.ValueOf(&schema.ErrInvalidGenesisTx).Elem(),
		"ErrInvalidSubmitTxNonce":      reflect.ValueOf(&schema.ErrInvalidSubmitTxNonce).Elem(),
		"EverTxActionBundle":           reflect.ValueOf(constant.MakeFromLiteral("\"bundle\"", token.STRING, 0)),
		"EverTxActionTransfer":         reflect.ValueOf(constant.MakeFromLiteral("\"transfer\"", token.STRING, 0)),

		// type definitions
		"GenesisTxData":   reflect.ValueOf((*schema.GenesisTxData)(nil)),
		"HaloTransaction": reflect.ValueOf((*schema.HaloTransaction)(nil)),
		"InfoRes":         reflect.ValueOf((*schema.InfoRes)(nil)),
		"TxApply":         reflect.ValueOf((*schema.TxApply)(nil)),
	}
}
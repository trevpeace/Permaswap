// Code generated by 'yaegi extract github.com/permadao/permaswap/halo/hvm/schema'. DO NOT EDIT.

package symbol

import (
	"github.com/permadao/permaswap/halo/hvm/schema"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["github.com/permadao/permaswap/halo/hvm/schema/schema"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CopyRouterState":         reflect.ValueOf(schema.CopyRouterState),
		"CopyToken":               reflect.ValueOf(schema.CopyToken),
		"ErrInsufficientStake":    reflect.ValueOf(&schema.ErrInsufficientStake).Elem(),
		"ErrInvalidAmount":        reflect.ValueOf(&schema.ErrInvalidAmount).Elem(),
		"ErrInvalidFee":           reflect.ValueOf(&schema.ErrInvalidFee).Elem(),
		"ErrInvalidFeeRecipient":  reflect.ValueOf(&schema.ErrInvalidFeeRecipient).Elem(),
		"ErrInvalidFromRouter":    reflect.ValueOf(&schema.ErrInvalidFromRouter).Elem(),
		"ErrInvalidNonce":         reflect.ValueOf(&schema.ErrInvalidNonce).Elem(),
		"ErrInvalidProposal":      reflect.ValueOf(&schema.ErrInvalidProposal).Elem(),
		"ErrInvalidProposer":      reflect.ValueOf(&schema.ErrInvalidProposer).Elem(),
		"ErrInvalidRouterAddress": reflect.ValueOf(&schema.ErrInvalidRouterAddress).Elem(),
		"ErrInvalidRouterName":    reflect.ValueOf(&schema.ErrInvalidRouterName).Elem(),
		"ErrInvalidStakePool":     reflect.ValueOf(&schema.ErrInvalidStakePool).Elem(),
		"ErrInvalidTx":            reflect.ValueOf(&schema.ErrInvalidTx).Elem(),
		"ErrInvalidTxAction":      reflect.ValueOf(&schema.ErrInvalidTxAction).Elem(),
		"ErrInvalidTxField":       reflect.ValueOf(&schema.ErrInvalidTxField).Elem(),
		"ErrInvalidTxParams":      reflect.ValueOf(&schema.ErrInvalidTxParams).Elem(),
		"ErrNoProposalFound":      reflect.ValueOf(&schema.ErrNoProposalFound).Elem(),
		"ErrNotARouter":           reflect.ValueOf(&schema.ErrNotARouter).Elem(),
		"ErrRouterAlreadyJoined":  reflect.ValueOf(&schema.ErrRouterAlreadyJoined).Elem(),
		"ErrTxExecuted":           reflect.ValueOf(&schema.ErrTxExecuted).Elem(),
		"Fee0005":                 reflect.ValueOf(&schema.Fee0005).Elem(),
		"Fee001":                  reflect.ValueOf(&schema.Fee001).Elem(),
		"Fee003":                  reflect.ValueOf(&schema.Fee003).Elem(),
		"Fee01":                   reflect.ValueOf(&schema.Fee01).Elem(),
		"PoolEco":                 reflect.ValueOf(constant.MakeFromLiteral("\"ecosystem\"", token.STRING, 0)),
		"PoolInc":                 reflect.ValueOf(constant.MakeFromLiteral("\"incentive\"", token.STRING, 0)),
		"PoolInv":                 reflect.ValueOf(constant.MakeFromLiteral("\"investor\"", token.STRING, 0)),
		"PoolTeam":                reflect.ValueOf(constant.MakeFromLiteral("\"team\"", token.STRING, 0)),
		"TxActionCall":            reflect.ValueOf(constant.MakeFromLiteral("\"call\"", token.STRING, 0)),
		"TxActionJoin":            reflect.ValueOf(constant.MakeFromLiteral("\"join\"", token.STRING, 0)),
		"TxActionLeave":           reflect.ValueOf(constant.MakeFromLiteral("\"leave\"", token.STRING, 0)),
		"TxActionPropose":         reflect.ValueOf(constant.MakeFromLiteral("\"propose\"", token.STRING, 0)),
		"TxActionStake":           reflect.ValueOf(constant.MakeFromLiteral("\"stake\"", token.STRING, 0)),
		"TxActionSwap":            reflect.ValueOf(constant.MakeFromLiteral("\"swap\"", token.STRING, 0)),
		"TxActionTransfer":        reflect.ValueOf(constant.MakeFromLiteral("\"transfer\"", token.STRING, 0)),
		"TxActionUnstake":         reflect.ValueOf(constant.MakeFromLiteral("\"unstake\"", token.STRING, 0)),
		"TxActionsSupported":      reflect.ValueOf(&schema.TxActionsSupported).Elem(),
		"TxVersionV1":             reflect.ValueOf(constant.MakeFromLiteral("\"v1\"", token.STRING, 0)),

		// type definitions
		"EverToken":        reflect.ValueOf((*schema.EverToken)(nil)),
		"Executor":         reflect.ValueOf((*schema.Executor)(nil)),
		"Pool":             reflect.ValueOf((*schema.Pool)(nil)),
		"Proposal":         reflect.ValueOf((*schema.Proposal)(nil)),
		"RouterState":      reflect.ValueOf((*schema.RouterState)(nil)),
		"State":            reflect.ValueOf((*schema.State)(nil)),
		"StateForProposal": reflect.ValueOf((*schema.StateForProposal)(nil)),
		"Transaction":      reflect.ValueOf((*schema.Transaction)(nil)),
		"TxApply":          reflect.ValueOf((*schema.TxApply)(nil)),
		"TxCallParams":     reflect.ValueOf((*schema.TxCallParams)(nil)),
		"TxProposeParams":  reflect.ValueOf((*schema.TxProposeParams)(nil)),
		"TxStakeParams":    reflect.ValueOf((*schema.TxStakeParams)(nil)),
		"TxSwapParams":     reflect.ValueOf((*schema.TxSwapParams)(nil)),
		"TxTransferParams": reflect.ValueOf((*schema.TxTransferParams)(nil)),
		"TxUnstakeParams":  reflect.ValueOf((*schema.TxUnstakeParams)(nil)),
	}
}

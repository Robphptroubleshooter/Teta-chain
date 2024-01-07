import (
	"fmt"

	"github.com/Pylons-tech/pylons/x/pylons/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

type HandlerOptions struct {
	ante.HandlerOptions

	AccountKeeper types.AccountKeeper
	PylonsKeeper  pylonskeeper.Keeper
	IBCKeeper     *keeper.Keeper
}

func NewAccountCreationDecorator(ak types.AccountKeeper) AccountCreationDecorator {
	return AccountCreationDecorator{
		ak: ak,
	}
}

func NewCustomSigVerificationDecorator(ak types.AccountKeeper, signModeHandler authsigning.SignModeHandler) CustomSigVerificationDecorator {
	return CustomSigVerificationDecorator{
		ak:              ak,
		signModeHandler: signModeHandler,
	}
}

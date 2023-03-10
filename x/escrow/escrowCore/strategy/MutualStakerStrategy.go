package strategy

type MutualStakerStrategy struct {
	Strategy
}

func (MutualStakerStrategy) AreCreateEscrowParamsValid(instigatorWager Coin, riderWager Coin) error {
	areFundsCorrect := instigatorWager > 0 && riderWager > 0

	if !areFundsCorrect {
		return ErrFooInvalidCreateAgreementParams
	}

	return nil
}

func newMutualStakerStrategy(escrowData Strategy) IStrategy {
	return &MutualStakerStrategy{
		Strategy: escrowData,
	}
}

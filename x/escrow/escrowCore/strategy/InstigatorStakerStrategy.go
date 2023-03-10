package strategy

type InstigatorStakerStrategy struct {
	Strategy
}

func (InstigatorStakerStrategy) AreCreateEscrowParamsValid(instigatorWager Coin, riderWager Coin) error {
	areFundsCorrect := instigatorWager > 0 && riderWager == 0

	if !areFundsCorrect {
		return ErrFooInvalidCreateAgreementParams
	}

	return nil
}

func newInstigatorStakerStrategy(escrowData Strategy) IStrategy {
	return &InstigatorStakerStrategy{
		Strategy: escrowData,
	}
}

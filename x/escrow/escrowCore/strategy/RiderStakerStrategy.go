package strategy

type RiderStakerStrategy struct {
	Strategy
}

func (RiderStakerStrategy) AreCreateEscrowParamsValid(instigatorWager Coin, riderWager Coin) error {
	areFundsCorrect := instigatorWager > 0 && riderWager == 0

	if !areFundsCorrect {
		return ErrFooInvalidCreateAgreementParams
	}

	return nil
}

func newRiderStakerStrategy(escrowData Strategy) IStrategy {
	return &RiderStakerStrategy{
		Strategy: escrowData,
	}
}

package strategy

import (
	"errors"
)

type StrategyName string
type EscrowStatus string
type ReleaseStatus string
type Address string
type Coin uint64
type Index uint64

const (
	STATUS_OPEN                  EscrowStatus = "STATUS_OPEN"
	STATUS_INSTIGATOR_APPROVED   EscrowStatus = "INSTIGATOR_APPROVED"
	STATUS_RIDER_APPROVED        EscrowStatus = "RIDER_APPROVED"
	STATUS_APPROVED              EscrowStatus = "APPROVED"
	STATUS_INSTIGATOR_DEPOSITED  EscrowStatus = "STATUS_INSTIGATOR_DEPOSITED"
	STATUS_RIDER_DEPOSITED       EscrowStatus = "STATUS_RIDER_DEPOSITED"
	STATUS_PENDING               EscrowStatus = "STATUS_PENDING"
	STATUS_RIDER_CONFIRMED       EscrowStatus = "RIDER_CONFIRMED"
	STATUS_CONFIRMED             EscrowStatus = "STATUS_CONFIRMED"
	STATUS_RIDER_WITHDRAWED      EscrowStatus = "STATUS_RIDER_WITHDRAWED"
	STATUS_INSIGNATOR_WITHDRAWED EscrowStatus = "STATUS_INSIGNATOR_WITHDRAWED"
	STATUS_WITHDRAWED            EscrowStatus = "STATUS_WITHDRAWED"
)

const (
	RELEASE_STATUS_NOT_RELEASED          ReleaseStatus = "NOT_RELEASED"
	RELEASE_STATUS_PENDING               ReleaseStatus = "PENDING"
	RELEASE_STATUS_INSTIGATOR_WITHDRAWED ReleaseStatus = "INSTIGATOR_WITHDRAWED"
	RELEASE_STATUS_RIDER_WITHDRAWED      ReleaseStatus = "RIDER_WITHDRAWED"
	RELEASE_STATUS_WITHDRAWED            ReleaseStatus = "WITHDRAWED"
)

const (
	INSTIGATOR_STAKER_STRATEGY StrategyName = "InstigatorStakerStrategy"
	MUTUAL_STAKER_STRATEGY     StrategyName = "MutualStakerStrategy"
	RIDER_STAKER_STRATEGY      StrategyName = "RiderStakerStrategy"
)

var Strategies = map[string]StrategyName{
	"1": INSTIGATOR_STAKER_STRATEGY,
	"2": MUTUAL_STAKER_STRATEGY,
	"3": RIDER_STAKER_STRATEGY,
}

type IStrategy interface {
	AreCreateEscrowParamsValid(instigatorWager Coin, riderWager Coin) error
	AreReleaseEscrowParamsValid(instigatorRelease Coin, riderRelease Coin) error
	IsStatusOfBeingReleasedAgreementValid() error
	Approve(from Address) error
	Deposit(from Address) (depositAmount Coin, err error)
	Confirm(from Address, winner Address) error
	Withdraw(from Address, releaseEscrowData Strategy) (
		withdrawAmount Coin,
		err error,
	)
}

type Strategy struct {
	Status            EscrowStatus
	Instigator        Address
	InstigatorWinner  Address
	InstigatorWager   Coin
	InstigatorBalance Coin
	Rider             Address
	RiderWinner       Address
	RiderWager        Coin
	RiderBalance      Coin
	InstigatorRelease Coin
	RiderRelease      Coin
	ReleaseStatus     ReleaseStatus
}

var ErrFooInvalidCreateAgreementParams = errors.New("ErrFoo: Invalid create agreement params")
var ErrFooInvalidStatusOfBeingReleasedAgreement = errors.New("ErrFoo: Invalid status of agreement which is being released")
var ErrFooInvalidCreateReleaseAgreementParams = errors.New("ErrFoo: Invalid create release agreement params")
var ErrFooAgreementReleased = errors.New("ErrFoo: Agreement already released")
var ErrFooInvalidEscrowStatus = errors.New("ErrFoo: Invalid status of agreement")
var ErrFooInvalidApproval = errors.New("ErrFoo: Invalid approval")
var ErrFooInvalidDeposit = errors.New("ErrFoo: Invalid deposit")
var ErrFooInvalidConfirm = errors.New("ErrFoo: Invalid confirm")
var ErrFooInvalidStrategy = errors.New("ErrFoo: Invalid strategy")
var ErrFooInvalidWithdraw = errors.New("ErrFoo: Invalid withdraw")

func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func (escrow Strategy) AreReleaseEscrowParamsValid(instigatorRelease Coin, riderRelease Coin) error {
	if escrow.ReleaseStatus != RELEASE_STATUS_NOT_RELEASED {
		return ErrFooAgreementReleased
	}

	isCorrectStatus := contains([]EscrowStatus{
		STATUS_APPROVED,
		STATUS_RIDER_DEPOSITED,
		STATUS_INSTIGATOR_DEPOSITED,
		STATUS_PENDING,
		STATUS_RIDER_APPROVED,
	}, escrow.Status)

	if !isCorrectStatus {
		return ErrFooInvalidStatusOfBeingReleasedAgreement
	}

	areFundsCorrect := escrow.InstigatorBalance+escrow.RiderBalance == instigatorRelease+riderRelease

	if !areFundsCorrect {
		return ErrFooInvalidCreateReleaseAgreementParams
	}

	return nil
}

func (escrow Strategy) IsStatusOfBeingReleasedAgreementValid() error {
	isCorrectStatus := contains([]EscrowStatus{
		STATUS_APPROVED,
		STATUS_RIDER_DEPOSITED,
		STATUS_INSTIGATOR_DEPOSITED,
		STATUS_PENDING,
		STATUS_RIDER_APPROVED,
	}, escrow.Status)

	if !isCorrectStatus {
		return ErrFooInvalidStatusOfBeingReleasedAgreement
	}

	return nil
}

func (escrow *Strategy) Approve(from Address) error {
	if escrow.ReleaseStatus != RELEASE_STATUS_NOT_RELEASED {
		return ErrFooAgreementReleased
	}

	isCorrectStatus := contains([]EscrowStatus{
		STATUS_OPEN,
		STATUS_RIDER_APPROVED,
		STATUS_INSTIGATOR_APPROVED,
	}, escrow.Status)

	if !isCorrectStatus {
		return ErrFooInvalidEscrowStatus
	}

	var status EscrowStatus
	if from == escrow.Instigator && escrow.Status != STATUS_INSTIGATOR_APPROVED {
		status = STATUS_INSTIGATOR_APPROVED
	} else if from == escrow.Rider && escrow.Status != STATUS_RIDER_APPROVED {
		status = STATUS_RIDER_APPROVED
	} else {
		return ErrFooInvalidApproval
	}

	var isReadyForNextStatus bool
	if from == escrow.Instigator {
		isReadyForNextStatus = escrow.Status == STATUS_RIDER_APPROVED
	} else if from == escrow.Rider {
		isReadyForNextStatus = escrow.Status == STATUS_INSTIGATOR_APPROVED
	} else {
		return ErrFooInvalidApproval
	}

	if isReadyForNextStatus {
		escrow.Status = STATUS_APPROVED
	} else {
		escrow.Status = status
	}

	return nil
}

func (escrow *Strategy) Deposit(from Address) (depositAmount Coin, err error) {
	if escrow.ReleaseStatus != RELEASE_STATUS_NOT_RELEASED {
		return 0, ErrFooAgreementReleased
	}

	isCorrectStatus := contains([]EscrowStatus{
		STATUS_APPROVED,
		STATUS_RIDER_DEPOSITED,
		STATUS_INSTIGATOR_DEPOSITED,
	}, escrow.Status)

	if !isCorrectStatus {
		return 0, ErrFooInvalidEscrowStatus
	}

	var status EscrowStatus
	if from == escrow.Instigator && escrow.Status != STATUS_INSTIGATOR_DEPOSITED {
		status = STATUS_INSTIGATOR_DEPOSITED
		depositAmount = escrow.InstigatorWager
		escrow.InstigatorBalance = escrow.InstigatorWager
	} else if from == escrow.Rider && escrow.Status != STATUS_RIDER_DEPOSITED {
		status = STATUS_RIDER_DEPOSITED
		depositAmount = escrow.RiderWager
		escrow.RiderBalance = escrow.RiderWager
	} else {
		return 0, ErrFooInvalidDeposit
	}

	var isReadyForNextStatus bool
	if from == escrow.Instigator {
		isReadyForNextStatus = escrow.Status == STATUS_RIDER_DEPOSITED
	} else if from == escrow.Rider {
		isReadyForNextStatus = escrow.Status == STATUS_INSTIGATOR_DEPOSITED
	} else {
		return 0, ErrFooInvalidApproval
	}

	if isReadyForNextStatus {
		escrow.Status = STATUS_PENDING
	} else {
		escrow.Status = status
	}

	return depositAmount, nil
}

func (escrow *Strategy) Confirm(from Address, winner Address) (err error) {
	if escrow.ReleaseStatus != RELEASE_STATUS_NOT_RELEASED {
		return ErrFooAgreementReleased
	}

	isCorrectStatus := contains([]EscrowStatus{
		STATUS_PENDING,
		STATUS_RIDER_CONFIRMED,
	}, escrow.Status)

	if !isCorrectStatus {
		return ErrFooInvalidEscrowStatus
	}

	if from == escrow.Instigator {
		escrow.InstigatorWinner = winner
	} else if from == escrow.Rider {
		escrow.RiderWinner = winner
	} else {
		return ErrFooInvalidConfirm
	}

	if escrow.InstigatorWinner == escrow.RiderWinner {
		escrow.Status = STATUS_CONFIRMED
	} else {
		escrow.Status = STATUS_PENDING
	}

	return nil
}

func ConfirmedAgreementWithdraw(from Address, escrow *Strategy) (
	withdrawAmount Coin,
	err error,
) {
	if escrow.Status != STATUS_CONFIRMED {
		return 0, ErrFooInvalidEscrowStatus
	}

	isWinnerSelected := escrow.InstigatorWinner == escrow.RiderWinner

	if !isWinnerSelected {
		return 0, ErrFooInvalidWithdraw
	}

	winner := escrow.InstigatorWinner

	if from != winner {
		return 0, ErrFooInvalidWithdraw
	}

	withdrawAmount = escrow.InstigatorBalance + escrow.RiderBalance

	escrow.Status = STATUS_WITHDRAWED
	escrow.InstigatorBalance = 0
	escrow.RiderBalance = 0

	return withdrawAmount, nil
}

func ReleasedAgreementWithdraw(from Address, escrow *Strategy, releaseEscrow Strategy) (
	withdrawAmount Coin,
	err error,
) {
	if escrow.ReleaseStatus != RELEASE_STATUS_WITHDRAWED ||
		(from == escrow.Instigator && escrow.ReleaseStatus == RELEASE_STATUS_INSTIGATOR_WITHDRAWED) ||
		(from == escrow.Rider && escrow.ReleaseStatus == RELEASE_STATUS_RIDER_WITHDRAWED) {
		return 0, ErrFooInvalidEscrowStatus
	}

	if from == escrow.Instigator {
		withdrawAmount = releaseEscrow.InstigatorRelease
	} else if from == escrow.Rider {
		withdrawAmount = releaseEscrow.RiderRelease
	} else {
		return 0, ErrFooInvalidWithdraw
	}

	if escrow.InstigatorBalance >= withdrawAmount {
		escrow.InstigatorBalance = escrow.InstigatorBalance - withdrawAmount
	} else {
		escrow.InstigatorBalance = 0
		escrow.RiderBalance = withdrawAmount - escrow.InstigatorBalance
	}

	return withdrawAmount, nil
}

func (escrow *Strategy) Withdraw(from Address, releaseEscrow Strategy) (
	withdrawAmount Coin,
	err error,
) {
	if escrow.ReleaseStatus == RELEASE_STATUS_NOT_RELEASED {
		return ConfirmedAgreementWithdraw(from, escrow)
	} else {
		return ReleasedAgreementWithdraw(from, escrow, releaseEscrow)
	}
}

func NewEscrowData(instigator Address, instigatorWager Coin, rider Address, riderWager Coin) Strategy {
	return Strategy{
		Status:            STATUS_OPEN,
		Instigator:        instigator,
		InstigatorWinner:  Address(""),
		InstigatorWager:   instigatorWager,
		InstigatorBalance: Coin(0),
		Rider:             rider,
		RiderWinner:       Address(""),
		RiderWager:        riderWager,
		RiderBalance:      Coin(0),
		InstigatorRelease: Coin(0),
		RiderRelease:      Coin(0),
		ReleaseStatus:     RELEASE_STATUS_NOT_RELEASED,
	}
}

func NewStrategy(strategy StrategyName, escrowData Strategy) (IStrategy, error) {
	switch strategy {
	case INSTIGATOR_STAKER_STRATEGY:
		return newInstigatorStakerStrategy(escrowData), nil
	case MUTUAL_STAKER_STRATEGY:
		return newMutualStakerStrategy(escrowData), nil
	case RIDER_STAKER_STRATEGY:
		return newRiderStakerStrategy(escrowData), nil
	default:
		return nil, ErrFooInvalidStrategy
	}
}

func AreCreateEscrowParamsValid(strategy StrategyName, instigatorWager Coin, riderWager Coin) error {
	switch strategy {
	case INSTIGATOR_STAKER_STRATEGY:
		return InstigatorStakerStrategy.AreCreateEscrowParamsValid(InstigatorStakerStrategy{}, instigatorWager, riderWager)
	case MUTUAL_STAKER_STRATEGY:
		return MutualStakerStrategy.AreCreateEscrowParamsValid(MutualStakerStrategy{}, instigatorWager, riderWager)
	case RIDER_STAKER_STRATEGY:
		return RiderStakerStrategy.AreCreateEscrowParamsValid(RiderStakerStrategy{}, instigatorWager, riderWager)
	default:
		return ErrFooInvalidStrategy
	}
}

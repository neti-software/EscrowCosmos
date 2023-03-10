package escrowCore

import (
	"encoding/json"
	"errors"
	s "escrow/x/escrow/escrowCore/strategy"
)

type NullableIndex int64 // -1 is null

const NullIndex NullableIndex = -1

var ErrFooInvalidEscrow = errors.New("ErrFoo: Invalid escrow")

type EscrowCore interface {
	Approve(from s.Address, releasedEscrow Escrow) error
	Deposit(from s.Address) (s.Coin, error)
	Confirm(from s.Address, winner s.Address) error
	Withdraw(from s.Address, releaseEscrowData s.Strategy) (
		withdrawAmount s.Coin,
		err error,
	)
	Serialize() (bool, string, error)
}

type Escrow struct {
	Index                    s.Index
	Strategy                 s.Strategy
	StrategyName             s.StrategyName
	ReleaseAgreementIndex    NullableIndex
	ReleasedByAgreementIndex NullableIndex
}

func (e *Escrow) Approve(from s.Address, releaseEscrow *Escrow) error {
	err := e.Strategy.Approve(from)

	if err != nil {
		return err
	}

	if e.Strategy.Status == s.STATUS_APPROVED && e.ReleaseAgreementIndex != NullIndex {
		err := releaseEscrow.Strategy.IsStatusOfBeingReleasedAgreementValid()

		if err != nil {
			return err
		}

		releaseEscrow.ReleasedByAgreementIndex = NullableIndex(e.Index)
	}

	return nil
}

func (e *Escrow) Deposit(from s.Address) (s.Coin, error) {
	return e.Strategy.Deposit(from)
}

func (e *Escrow) Confirm(from s.Address, winner s.Address) error {
	return e.Strategy.Confirm(from, winner)
}

func (e *Escrow) Withdraw(from s.Address, releaseEscrow *Escrow) (
	withdrawAmount s.Coin,
	err error,
) {
	return e.Strategy.Withdraw(from, releaseEscrow.Strategy)
}

func (e *Escrow) Serialize() (string, error) {
	serializedEscrow, err := json.Marshal(e)
	if err != nil {
		return "", ErrFooInvalidEscrow
	}

	return string(serializedEscrow), nil
}

func NewEscrow(
	newIndex s.Index,
	strategy s.StrategyName,
	instigator s.Address,
	instigatorWager s.Coin,
	rider s.Address,
	riderWager s.Coin) (string, error) {

	err := s.AreCreateEscrowParamsValid(strategy, instigatorWager, riderWager)

	if err != nil {
		return "", err
	}

	escrowData := s.NewEscrowData(instigator, instigatorWager, rider, riderWager)

	escrow := Escrow{
		Index:                    newIndex,
		Strategy:                 escrowData,
		StrategyName:             strategy,
		ReleaseAgreementIndex:    NullIndex,
		ReleasedByAgreementIndex: NullIndex,
	}

	serializedEscrow, err := escrow.Serialize()

	if err != nil {
		return "", err
	}

	return serializedEscrow, nil
}

func NewReleaseEscrow(
	newIndex s.Index,
	releaseEscrow Escrow,
	instigatorRelease s.Coin,
	riderRelease s.Coin,
	strategy s.StrategyName,
	instigatorWager s.Coin,
	riderWager s.Coin) (Escrow, error) {

	err := releaseEscrow.Strategy.AreReleaseEscrowParamsValid(instigatorRelease, riderRelease)

	if err != nil {
		return Escrow{}, err
	}

	statusErr := releaseEscrow.Strategy.IsStatusOfBeingReleasedAgreementValid()

	if statusErr != nil {
		return Escrow{}, statusErr
	}

	newEscrowData := s.NewEscrowData(releaseEscrow.Strategy.Instigator, instigatorWager, releaseEscrow.Strategy.Rider, riderWager)

	newEscrow := Escrow{
		Index:                    newIndex,
		Strategy:                 newEscrowData,
		StrategyName:             strategy,
		ReleaseAgreementIndex:    NullableIndex(releaseEscrow.Index),
		ReleasedByAgreementIndex: NullIndex,
	}

	return newEscrow, nil
}

func InitEscrow(serializedEscrow string) (Escrow, error) {
	var escrow Escrow
	err := json.Unmarshal([]byte(serializedEscrow), &escrow)

	return escrow, err
}

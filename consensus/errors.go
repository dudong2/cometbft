package consensus

import (
	"errors"
)

var (
	// Consensus sentinel errors
	ErrInvalidProposalSignature   = errors.New("error invalid proposal signature")
	ErrInvalidProposalPOLRound    = errors.New("error invalid proposal POL round")
	ErrAddingVote                 = errors.New("error adding vote")
	ErrSignatureFoundInPastBlocks = errors.New("found signature from the same key")
	ErrPubKeyIsNotSet             = errors.New("pubkey is not set. Look for \"Can't get private validator pubkey\" errors")
)

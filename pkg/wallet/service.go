package wallet

import (
	"errors"

	"github.com/google/uuid"

	"github.com/Kamolov-Daler/wallet/pkg/types"
)

var ErrPhoneRegistered = errors.New("phone already registered")
var ErrAmountMustBePosititve = errors.New("amount must be greater than zero")
var ErrAccountNotFound = errors.New("account not found")
var ErrNotEnoughBalance = errors.New("not enough balance")
var ErrPaymentNotFound = errors.New("payment not found")
var ErrFavoriteNotFound = errors.New("favorite payment not found")

type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
	favorites     []*types.Favorite
}

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePosititve
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}

	if account == nil {
		return ErrAccountNotFound
	}

	account.Balance += amount
	return nil
}

func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {
	if amount <= 0 {
		return nil, ErrAmountMustBePosititve
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}

	if account.Balance < amount {
		return nil, ErrNotEnoughBalance
	}

	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID:        paymentID,
		AccountID: accountID,
		Amount:    amount,
		Category:  category,
		Status:    types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}

func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	findAccount := &types.Account{}

	for _, acc := range s.accounts {
		if acc.ID == accountID {
			findAccount = acc
			return findAccount, nil
		}
	}
	return nil, ErrAccountNotFound
}

func (s *Service) FindPaymentByID(paymentID string) (*types.Payment, error) {
	for _, payment := range s.payments {
		if paymentID == payment.ID {
			return payment, nil
		}
	}

	return nil, ErrPaymentNotFound
}

func (s *Service) Reject(paymentID string) error {
	var targetPayment *types.Payment
	for _, payment := range s.payments {
		if payment.ID == paymentID {
			targetPayment = payment
			break
		}
	}
	if targetPayment == nil {
		return ErrPaymentNotFound
	}

	account, err := s.FindAccountByID(targetPayment.AccountID)
	if err != nil {
		return err
	}

	targetPayment.Status = types.PaymentStatusFail
	account.Balance += targetPayment.Amount
	return nil
}

func (s *Service) Repeat(paymentID string) (*types.Payment, error) {
	payment, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil, ErrPaymentNotFound
	}

	account, err := s.FindAccountByID(payment.AccountID)
	account.Balance -= payment.Amount
	newPaymentID := uuid.New().String()
	newPayment := &types.Payment{
		ID:        newPaymentID,
		AccountID: payment.AccountID,
		Amount:    payment.Amount,
		Category:  payment.Category,
		Status:    types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, newPayment)
	return newPayment, nil
}

func (s *Service) FavoritePayment(paymentID string, name string) (*types.Favorite, error) {
	targetPayment := &types.Payment{}
	for _, payment := range s.payments {
		if payment.ID == paymentID {
			targetPayment = payment
		}
	}

	if targetPayment == nil {
		return nil, ErrPaymentNotFound
	}

	newFavoriteID := uuid.New().String()
	favorite := &types.Favorite{
		ID:        newFavoriteID,
		AccountID: targetPayment.AccountID,
		Name:      name,
		Amount:    targetPayment.Amount,
		Category:  targetPayment.Category,
	}
	s.favorites = append(s.favorites, favorite)
	return favorite, nil
}

func (s *Service) PayFromFavorite(favoriteID string) (*types.Payment, error) {
	targetFavorite := &types.Favorite{}
	for _, favorite := range s.favorites {
		if favorite.ID == favoriteID {
			targetFavorite = favorite
		}
	}

	if targetFavorite == nil {
		return nil, ErrFavoriteNotFound
	}

	account, err := s.FindAccountByID(targetFavorite.AccountID)
	if err != nil {
		return nil, ErrAccountNotFound
	}
	account.Balance -= targetFavorite.Amount
	newPaymentID := uuid.New().String()
	payment := &types.Payment{
		ID:        newPaymentID,
		AccountID: targetFavorite.AccountID,
		Amount:    targetFavorite.Amount,
		Category:  targetFavorite.Category,
		Status:    types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}

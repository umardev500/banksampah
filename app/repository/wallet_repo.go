package repository

import (
	"context"
	"fmt"

	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type walletRepo struct {
	pgxConfig *config.PgxConfig
}

func NewWalletRepository(pgxConfig *config.PgxConfig) domain.WalletRepository {
	return &walletRepo{
		pgxConfig: pgxConfig,
	}
}

func (repo *walletRepo) SetBalance(ctx context.Context, payload model.WalletSetBalanceRequest) (balance *float64, err error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := fmt.Sprintf(`--sql
		UPDATE wallets SET balance = balance %s $1 WHERE id = $2
		RETURNING balance
	`, payload.SetType)

	var tempBalance float64
	row := queries.QueryRow(ctx, sql, payload.Amount, payload.ID)
	err = row.Scan(&tempBalance)
	if err != nil {
		return nil, err
	}

	balance = &tempBalance

	return
}

func (repo *walletRepo) UpdateByID(ctx context.Context, payload model.WalletCreateOrUpdateRequest) (returning *model.Wallet, err error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		UPDATE wallets SET
	`
	rawSql, args := util.BuildUpdateQuery(sql, payload, []types.Filter{
		{
			Field:    "id",
			Operator: "=",
			Value:    payload.ID,
		},
	})
	if args == nil {
		return
	}

	_, err = queries.Exec(ctx, rawSql, args...)

	return
}

func (repo *walletRepo) MoveBalanceToWallet(ctx context.Context, payload model.WalletMoveBalanceRequest) ([]model.Wallet, error) {
	decreaseWalletFrom := `--sql
		UPDATE wallets SET
		balance = balance - $1
		WHERE id = $2
		AND balance >= $1
		RETURNING *
	`
	increaseWalletTo := `--sql
		UPDATE wallets SET
		balance = balance + $1
		WHERE id = $2
		RETURNING *
	`

	var toAndFromWallet []model.Wallet

	// Transaction
	err := repo.pgxConfig.WithTransaction(ctx, func(ctx context.Context) error {
		// Decrease wallet from
		row := repo.pgxConfig.TrOrDB(ctx).QueryRow(ctx, decreaseWalletFrom, payload.Amount, payload.FromWalletID)
		var result model.Wallet
		err := row.Scan(
			&result.ID,
			&result.UserD,
			&result.Name,
			&result.Balance,
			&result.Description,
			&result.Type,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt,
		)
		if err != nil {
			return err
		}

		toAndFromWallet = append(toAndFromWallet, result)

		// Increase wallet to
		row = repo.pgxConfig.TrOrDB(ctx).QueryRow(ctx, increaseWalletTo, payload.Amount, payload.ToWalletID)
		err = row.Scan(
			&result.ID,
			&result.UserD,
			&result.Name,
			&result.Balance,
			&result.Description,
			&result.Type,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt,
		)
		if err != nil {
			return err
		}

		toAndFromWallet = append(toAndFromWallet, result)

		return err
	})

	if err != nil {
		return nil, err
	}

	return toAndFromWallet, nil
}

func (repo *walletRepo) FindByID(ctx context.Context, id string) (model.Wallet, error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		SELECT * FROM wallets WHERE id=$1
	`
	row := queries.QueryRow(ctx, sql, id)
	var result model.Wallet
	err := row.Scan(
		&result.ID,
		&result.UserD,
		&result.Name,
		&result.Balance,
		&result.Description,
		&result.Type,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)
	return result, err
}

func (repo *walletRepo) DeleteByID(ctx context.Context, id string) error {
	sql := `--sql
		DELETE FROM wallets WHERE id=$1 AND "type" != 'master'
	`

	// Find wallet by id to get wallet info
	wallet, err := repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	userID := wallet.UserD
	prevBalance := wallet.Balance

	restoreBalanceToMaster := `--sql
		UPDATE wallets SET balance = balance + $1 WHERE "type" = 'master' AND user_id = $2
	`

	// Transaction
	err = repo.pgxConfig.WithTransaction(ctx, func(ctx context.Context) error {
		queries := repo.pgxConfig.TrOrDB(ctx)
		result, err := queries.Exec(ctx, sql, id)
		if err != nil {
			return err
		}

		rowsAffected := result.RowsAffected()

		// Update balance to master
		if rowsAffected > 0 {
			_, err = queries.Exec(ctx, restoreBalanceToMaster, prevBalance, userID)
			if err != nil {
				return err
			}
		}

		return err
	})

	return err
}

func (repo *walletRepo) Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) (*model.Wallet, error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		INSERT INTO wallets (id, user_id, "name", "description", "type") VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`
	// result, err := queries.Exec(ctx, sql, payload.ID, payload.UserID, payload.Name, payload.Description)
	// fmt.Println(result)

	row := queries.QueryRow(ctx, sql, payload.ID, payload.UserID, payload.Name, payload.Description, payload.Type)
	var result model.Wallet
	err := row.Scan(
		&result.ID,
		&result.UserD,
		&result.Name,
		&result.Balance,
		&result.Description,
		&result.Type,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	return &result, err
}

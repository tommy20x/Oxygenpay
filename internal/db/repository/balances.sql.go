// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: balances.sql

package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const createBalance = `-- name: CreateBalance :one
insert into balances (
  created_at, updated_at,
  entity_id, entity_type, uuid,
  network, network_id, currency_type, currency,
  decimals, amount
) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
returning id, created_at, updated_at, entity_id, entity_type, network, network_id, currency_type, currency, decimals, amount, uuid
`

type CreateBalanceParams struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	EntityID     int64
	EntityType   string
	Uuid         uuid.NullUUID
	Network      string
	NetworkID    string
	CurrencyType string
	Currency     string
	Decimals     int32
	Amount       pgtype.Numeric
}

func (q *Queries) CreateBalance(ctx context.Context, arg CreateBalanceParams) (Balance, error) {
	row := q.db.QueryRow(ctx, createBalance,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.EntityID,
		arg.EntityType,
		arg.Uuid,
		arg.Network,
		arg.NetworkID,
		arg.CurrencyType,
		arg.Currency,
		arg.Decimals,
		arg.Amount,
	)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EntityID,
		&i.EntityType,
		&i.Network,
		&i.NetworkID,
		&i.CurrencyType,
		&i.Currency,
		&i.Decimals,
		&i.Amount,
		&i.Uuid,
	)
	return i, err
}

const getBalanceByFilter = `-- name: GetBalanceByFilter :one
select id, created_at, updated_at, entity_id, entity_type, network, network_id, currency_type, currency, decimals, amount, uuid from balances where entity_id = $1 and entity_type = $2 and network_id = $3 and currency = $4 limit 1
`

type GetBalanceByFilterParams struct {
	EntityID   int64
	EntityType string
	NetworkID  string
	Currency   string
}

func (q *Queries) GetBalanceByFilter(ctx context.Context, arg GetBalanceByFilterParams) (Balance, error) {
	row := q.db.QueryRow(ctx, getBalanceByFilter,
		arg.EntityID,
		arg.EntityType,
		arg.NetworkID,
		arg.Currency,
	)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EntityID,
		&i.EntityType,
		&i.Network,
		&i.NetworkID,
		&i.CurrencyType,
		&i.Currency,
		&i.Decimals,
		&i.Amount,
		&i.Uuid,
	)
	return i, err
}

const getBalanceByFilterWithLock = `-- name: GetBalanceByFilterWithLock :one
select id, created_at, updated_at, entity_id, entity_type, network, network_id, currency_type, currency, decimals, amount, uuid from balances
where entity_id = $1 and entity_type = $2 and network_id = $3 and currency = $4
FOR NO KEY UPDATE limit 1
`

type GetBalanceByFilterWithLockParams struct {
	EntityID   int64
	EntityType string
	NetworkID  string
	Currency   string
}

func (q *Queries) GetBalanceByFilterWithLock(ctx context.Context, arg GetBalanceByFilterWithLockParams) (Balance, error) {
	row := q.db.QueryRow(ctx, getBalanceByFilterWithLock,
		arg.EntityID,
		arg.EntityType,
		arg.NetworkID,
		arg.Currency,
	)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EntityID,
		&i.EntityType,
		&i.Network,
		&i.NetworkID,
		&i.CurrencyType,
		&i.Currency,
		&i.Decimals,
		&i.Amount,
		&i.Uuid,
	)
	return i, err
}

const getBalanceByID = `-- name: GetBalanceByID :one
select id, created_at, updated_at, entity_id, entity_type, network, network_id, currency_type, currency, decimals, amount, uuid from balances where entity_id = $1 and entity_type = $2 and id = $3
`

type GetBalanceByIDParams struct {
	EntityID   int64
	EntityType string
	ID         int64
}

func (q *Queries) GetBalanceByID(ctx context.Context, arg GetBalanceByIDParams) (Balance, error) {
	row := q.db.QueryRow(ctx, getBalanceByID, arg.EntityID, arg.EntityType, arg.ID)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EntityID,
		&i.EntityType,
		&i.Network,
		&i.NetworkID,
		&i.CurrencyType,
		&i.Currency,
		&i.Decimals,
		&i.Amount,
		&i.Uuid,
	)
	return i, err
}

const getBalanceByIDWithLock = `-- name: GetBalanceByIDWithLock :one
select id, created_at, updated_at, entity_id, entity_type, network, network_id, currency_type, currency, decimals, amount, uuid from balances where id = $1 FOR NO KEY UPDATE
`

func (q *Queries) GetBalanceByIDWithLock(ctx context.Context, id int64) (Balance, error) {
	row := q.db.QueryRow(ctx, getBalanceByIDWithLock, id)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EntityID,
		&i.EntityType,
		&i.Network,
		&i.NetworkID,
		&i.CurrencyType,
		&i.Currency,
		&i.Decimals,
		&i.Amount,
		&i.Uuid,
	)
	return i, err
}

const getBalanceByUUID = `-- name: GetBalanceByUUID :one
select id, created_at, updated_at, entity_id, entity_type, network, network_id, currency_type, currency, decimals, amount, uuid from balances where entity_id = $1 and entity_type = $2 and uuid = $3
`

type GetBalanceByUUIDParams struct {
	EntityID   int64
	EntityType string
	Uuid       uuid.NullUUID
}

func (q *Queries) GetBalanceByUUID(ctx context.Context, arg GetBalanceByUUIDParams) (Balance, error) {
	row := q.db.QueryRow(ctx, getBalanceByUUID, arg.EntityID, arg.EntityType, arg.Uuid)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EntityID,
		&i.EntityType,
		&i.Network,
		&i.NetworkID,
		&i.CurrencyType,
		&i.Currency,
		&i.Decimals,
		&i.Amount,
		&i.Uuid,
	)
	return i, err
}

const insertBalanceAuditLog = `-- name: InsertBalanceAuditLog :exec
insert into balance_audit_log(created_at, balance_id, comment, metadata) values ($1,$2, $3, $4) returning id, created_at, balance_id, comment, metadata
`

type InsertBalanceAuditLogParams struct {
	CreatedAt time.Time
	BalanceID int64
	Comment   string
	Metadata  pgtype.JSONB
}

func (q *Queries) InsertBalanceAuditLog(ctx context.Context, arg InsertBalanceAuditLogParams) error {
	_, err := q.db.Exec(ctx, insertBalanceAuditLog,
		arg.CreatedAt,
		arg.BalanceID,
		arg.Comment,
		arg.Metadata,
	)
	return err
}

const listAllBalancesByType = `-- name: ListAllBalancesByType :many
select id, created_at, updated_at, entity_id, entity_type, network, network_id, currency_type, currency, decimals, amount, uuid from balances
where entity_type = $1
and (CASE WHEN $2::boolean THEN amount > 0 ELSE true END)
order by id desc
`

type ListAllBalancesByTypeParams struct {
	EntityType string
	HideEmpty  bool
}

func (q *Queries) ListAllBalancesByType(ctx context.Context, arg ListAllBalancesByTypeParams) ([]Balance, error) {
	rows, err := q.db.Query(ctx, listAllBalancesByType, arg.EntityType, arg.HideEmpty)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Balance
	for rows.Next() {
		var i Balance
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.EntityID,
			&i.EntityType,
			&i.Network,
			&i.NetworkID,
			&i.CurrencyType,
			&i.Currency,
			&i.Decimals,
			&i.Amount,
			&i.Uuid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBalances = `-- name: ListBalances :many
select id, created_at, updated_at, entity_id, entity_type, network, network_id, currency_type, currency, decimals, amount, uuid from balances
where entity_type = $1 and entity_id = $2
order by id desc
`

type ListBalancesParams struct {
	EntityType string
	EntityID   int64
}

func (q *Queries) ListBalances(ctx context.Context, arg ListBalancesParams) ([]Balance, error) {
	rows, err := q.db.Query(ctx, listBalances, arg.EntityType, arg.EntityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Balance
	for rows.Next() {
		var i Balance
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.EntityID,
			&i.EntityType,
			&i.Network,
			&i.NetworkID,
			&i.CurrencyType,
			&i.Currency,
			&i.Decimals,
			&i.Amount,
			&i.Uuid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBalanceByID = `-- name: UpdateBalanceByID :one
update balances
set updated_at= $2, amount = balances.amount + $3
where id = $1
returning id, created_at, updated_at, entity_id, entity_type, network, network_id, currency_type, currency, decimals, amount, uuid
`

type UpdateBalanceByIDParams struct {
	ID        int64
	UpdatedAt time.Time
	Amount    pgtype.Numeric
}

func (q *Queries) UpdateBalanceByID(ctx context.Context, arg UpdateBalanceByIDParams) (Balance, error) {
	row := q.db.QueryRow(ctx, updateBalanceByID, arg.ID, arg.UpdatedAt, arg.Amount)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EntityID,
		&i.EntityType,
		&i.Network,
		&i.NetworkID,
		&i.CurrencyType,
		&i.Currency,
		&i.Decimals,
		&i.Amount,
		&i.Uuid,
	)
	return i, err
}
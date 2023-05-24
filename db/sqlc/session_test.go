package db

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomSession(t *testing.T) Session {
	storeRandom := CreateRandomStore(t)
	arg := CreateSessionParams{
		ID:           uuid.New(),
		IDStore:      storeRandom.ID,
		RefreshToken: "test",
		UserAgent:    "test",
		ClientIp:     "test",
		IsBlocked:    false,
		ExpiresAt:    storeRandom.CreatedAt,
	}

	session, err := testQueries.CreateSession(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, session)

	require.Equal(t, arg.ID, session.ID)
	return session
}

func TestCreateSession(t *testing.T) {
	CreateRandomSession(t)
}

func TestGetSession(t *testing.T) {
	mockTimer := time.Now()
	generateUUID := uuid.New()

	db, momockTimer, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "id_store", "refresh_token", "user_agent", "client_ip", "is_blocked", "expires_at", "created_at"}).
		AddRow(generateUUID, 1, "test", "test", "test", false, mockTimer, mockTimer)

	momockTimer.ExpectQuery(regexp.QuoteMeta(`SELECT id, id_store, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at FROM sessions WHERE id = $1 LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(rows)

	session, err := testQueries.GetSession(context.Background(), generateUUID)

	require.NoError(t, err)

	require.NotEmpty(t, session)
}

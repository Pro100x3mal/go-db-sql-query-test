package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "./demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)
	require.NoError(t, err)
	assert.Equal(t, clientID, cl.ID)
	assert.NotEqual(t, "", cl.FIO)
	assert.NotEqual(t, "", cl.Login)
	assert.NotEqual(t, "", cl.Email)
	assert.NotEqual(t, "", cl.Birthday)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "./demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	clientID := -1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)
	require.ErrorIs(t, err, sql.ErrNoRows)
	assert.Equal(t, 0, cl.ID)
	assert.Equal(t, "", cl.FIO)
	assert.Equal(t, "", cl.Login)
	assert.Equal(t, "", cl.Email)
	assert.Equal(t, "", cl.Birthday)

}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "./demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEqual(t, 0, cl.ID)

	getCl, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	assert.Equal(t, cl, getCl)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "./demo.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEqual(t, 0, cl.ID)

	_, err = selectClient(db, cl.ID)
	require.NoError(t, err)

	err = deleteClient(db, cl.ID)
	require.NoError(t, err)

	_, err = selectClient(db, cl.ID)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

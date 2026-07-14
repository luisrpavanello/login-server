package database

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentibiabr/login-server/src/serviceerrors"
)

func loadPlayersOnline(db *sql.DB) (uint32, error) {
	if db == nil {
		return 0, errors.New("database connection is nil")
	}

	var playersOnline uint32
	err := db.QueryRow("SELECT COUNT(*) FROM players_online").Scan(&playersOnline)
	if err != nil {
		return 0, err
	}

	return playersOnline, nil
}

func HandleCacheInfo(c *gin.Context, db *sql.DB) {
	playersOnline, err := loadPlayersOnline(db)
	if err != nil {
		writeServiceError(c, serviceerrors.GameData(
			serviceerrors.CodeDatabaseUnavailable,
			"DATABASE_UNAVAILABLE",
			err,
		))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"playersonline":        playersOnline,
		"discord_online":       0,
		"discord_link":         "",
		"gamingyoutubestreams": 0,
		"gamingyoutubeviewer":  0,
		"youtube_link":         "",
	})
}

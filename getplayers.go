package jpmcpvpgo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Players struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	UUID            string           `json:"uuid"`
	LastLogin       time.Time        `json:"last_login"`
	LastLoginServer string           `json:"last_login_server"`
	LastLogout      time.Time        `json:"last_logout"`
	Banned          bool             `json:"banned"`
	Bowspleef       Bowspleef        `json:"bowspleef"`
	CP              CP               `json:"cp"`
	CTW             CTW              `json:"ctw"`
	Factions        Factions         `json:"factions"`
	Kills           []Stats          `json:"kills"`
	Deaths          []Stats          `json:"deaths"`
	Objective       Objective        `json:"objective"`
	PlayersMatches  []PlayersMatches `json:"matches"`
	Paintball       Paintball        `json:"paintball"`
	Splatt          Splatt           `json:"splatt"`
	Teampvp         Teampvp          `json:"teampvp"`
	Tntrun          Tntrun           `json:"tntrun"`
	Total           Total            `json:"total"`
	PermalinkURL    string           `json:"permalink_url"`
}

type Bowspleef struct {
	LoseCount  int `json:"lose_count"`
	PlayCount  int `json:"play_count"`
	TimePlayed int `json:"time_played"`
	WinCount   int `json:"win_count"`
}

type CP struct {
	DeathCount    int `json:"death_count"`
	EnvdeathCount int `json:"envdeath_count"`
	KillCount     int `json:"kill_count"`
	LoseCount     int `json:"lose_count"`
	PlayCount     int `json:"play_count"`
	TimePlayed    int `json:"time_played"`
	WinCount      int `json:"win_count"`
}

type CTW struct {
	DeathCount    int          `json:"death_count"`
	EnvdeathCount int          `json:"envdeath_count"`
	KillCount     int          `json:"kill_count"`
	LoseCount     int          `json:"lose_count"`
	PlayCount     int          `json:"play_count"`
	TimePlayed    int          `json:"time_played"`
	WinCount      int          `json:"win_count"`
	WoolPlaced    int          `json:"wool_placed"`
	WoolPlaces    []WoolPlaces `json:"wool_places"`
}

type WoolPlaces struct {
	Color     string    `json:"color"`
	Map       string    `json:"map"`
	MatchID   string    `json:"match_id"`
	Team      string    `json:"team"`
	TeamColor string    `json:"team_color"`
	Time      time.Time `json:"time"`
}

type Factions struct {
	Blocks                Blocks          `json:"blocks"`
	DeathCount            int             `json:"death_count"`
	Deaths                []FactionsStats `json:"deaths"`
	FactionID             string          `json:"faction_id"`
	FactionRole           string          `json:"faction_role"`
	KillCount             int             `json:"kill_count"`
	Kills                 []FactionsStats `json:"kills"`
	LevelAccumulated      int             `json:"level_accumulated"`
	Power                 int             `json:"power"`
	PvpProtectionDisabled bool            `json:"pvp_protection_disabled"`
	TimePlayed            int             `json:"time_played"`
}

type Blocks struct {
	COALORE     COALORE     `json:"COAL_ORE"`
	DIAMONDORE  DIAMONDORE  `json:"DIAMOND_ORE"`
	GOLDORE     GOLDORE     `json:"GOLD_ORE"`
	IRONORE     IRONORE     `json:"IRON_ORE"`
	LAPISORE    LAPISORE    `json:"LAPIS_ORE"`
	QUARTZORE   QUARTZORE   `json:"QUARTZ_ORE"`
	REDSTONEORE REDSTONEORE `json:"REDSTONE_ORE"`
	STONE       STONE       `json:"STONE"`
}

type COALORE struct {
	Break int `json:"break"`
}

type DIAMONDORE struct {
	Break int `json:"break"`
}

type GOLDORE struct {
	Break int `json:"break"`
}

type IRONORE struct {
	Break int `json:"break"`
}

type LAPISORE struct {
	Break int `json:"break"`
}

type QUARTZORE struct {
	Break int `json:"break"`
}

type REDSTONEORE struct {
	Break int `json:"break"`
}

type STONE struct {
	Break int `json:"break"`
}

type Location struct {
	World string `json:"world"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Z     int    `json:"z"`
}

type FactionsStats struct {
	Killer   string    `json:"killer"`
	Location Location  `json:"location"`
	Time     time.Time `json:"time"`
	Victim   string    `json:"victim"`
}

type Stats struct {
	Killer   string    `json:"killer"`
	Location Location  `json:"location"`
	Map      string    `json:"map"`
	MatchID  string    `json:"match_id"`
	Time     time.Time `json:"time"`
	Victim   string    `json:"victim"`
}

type Objective struct {
	CoreLeaked           int         `json:"core_leaked"`
	CoreLeaks            []CoreLeaks `json:"core_leaks"`
	DeathCount           int         `json:"death_count"`
	DestroyableDestroyed int         `json:"destroyable_destroyed"`
	Destroys             []Destroys  `json:"destroys"`
	EnvdeathCount        int         `json:"envdeath_count"`
	KillCount            int         `json:"kill_count"`
	LoseCount            int         `json:"lose_count"`
	PlayCount            int         `json:"play_count"`
	TimePlayed           int         `json:"time_played"`
	WinCount             int         `json:"win_count"`
}

type CoreLeaks struct {
	Map       string    `json:"map"`
	MatchID   string    `json:"match_id"`
	Name      string    `json:"name"`
	Team      string    `json:"team"`
	TeamColor string    `json:"team_color"`
	Time      time.Time `json:"time"`
}

type Destroys struct {
	Map       string    `json:"map"`
	MatchID   string    `json:"match_id"`
	Name      string    `json:"name"`
	Percent   int       `json:"percent"`
	Team      string    `json:"team"`
	TeamColor string    `json:"team_color"`
	Time      time.Time `json:"time"`
}

type PlayersMatches struct {
	DeathCount              int       `json:"death_count"`
	EnvdeathCount           int       `json:"envdeath_count"`
	Finished                time.Time `json:"finished"`
	Gamemode                string    `json:"gamemode"`
	GlickoKillDeviationDiff float64   `json:"glicko_kill_deviation_diff,omitempty"`
	GlickoKillRatingDiff    float64   `json:"glicko_kill_rating_diff,omitempty"`
	ID                      string    `json:"id"`
	KillCount               int       `json:"kill_count"`
	Map                     string    `json:"map"`
	Result                  string    `json:"result"`
	Score                   int       `json:"score"`
	Started                 time.Time `json:"started"`
	Time                    time.Time `json:"time"`
}

type Paintball struct {
	DeathCount int `json:"death_count"`
	HitCount   int `json:"hit_count"`
	KillCount  int `json:"kill_count"`
	LoseCount  int `json:"lose_count"`
	PlayCount  int `json:"play_count"`
	ShotCount  int `json:"shot_count"`
	TimePlayed int `json:"time_played"`
	WinCount   int `json:"win_count"`
}

type Splatt struct {
	DeathCount int `json:"death_count"`
	HitCount   int `json:"hit_count"`
	KillCount  int `json:"kill_count"`
	LoseCount  int `json:"lose_count"`
	PaintCount int `json:"paint_count"`
	PlayCount  int `json:"play_count"`
	ShotCount  int `json:"shot_count"`
	TimePlayed int `json:"time_played"`
	WinCount   int `json:"win_count"`
}

type Teampvp struct {
	DeathCount             int     `json:"death_count"`
	EnvdeathCount          int     `json:"envdeath_count"`
	KdRatio                float64 `json:"kd_ratio"`
	KdRatioRank            int     `json:"kd_ratio_rank"`
	KillCount              int     `json:"kill_count"`
	KillCountRank          int     `json:"kill_count_rank"`
	KkRatio                float64 `json:"kk_ratio"`
	KkRatioRank            int     `json:"kk_ratio_rank"`
	LoseCount              int     `json:"lose_count"`
	ObjectiveCompleted     int     `json:"objective_completed"`
	ObjectiveCompletedRank int     `json:"objective_completed_rank"`
	PlayCount              int     `json:"play_count"`
	Score                  int     `json:"score"`
	TimePlayed             int     `json:"time_played"`
	TimePlayedRank         int     `json:"time_played_rank"`
	WinCount               int     `json:"win_count"`
	WinCountRank           int     `json:"win_count_rank"`
}

type Total struct {
	DeathCount      int `json:"death_count"`
	EnvdeathCount   int `json:"envdeath_count"`
	HitCount        int `json:"hit_count"`
	KillCount       int `json:"kill_count"`
	LoseCount       int `json:"lose_count"`
	NyanAccumulated int `json:"nyan_accumulated"`
	PlayCount       int `json:"play_count"`
	ShotCount       int `json:"shot_count"`
	TimePlayed      int `json:"time_played"`
	WinCount        int `json:"win_count"`
}

type Tntrun struct {
	LoseCount  int `json:"lose_count"`
	PlayCount  int `json:"play_count"`
	TimePlayed int `json:"time_played"`
	WinCount   int `json:"win_count"`
}

type PlayersOptions struct {
	Id     string
	Fields string
}

func (session Session) GetPlayers(options PlayersOptions) (*Players, error) {
	if options.Id == "" {
		return nil, fmt.Errorf("options.Id is empty")
	}
	values := url.Values{}
	if options.Fields != "" {
		values.Add("fields", options.Fields)
	}

	body, err := session.request("players/"+options.Id, values)

	if err != nil {
		return nil, err
	}

	players := Players{}

	err = json.Unmarshal(body, &players)

	return &players, err
}

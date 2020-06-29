package jpmcpvpgo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Players struct {
	ID              *string           `json:"id,omitempty"`
	Name            *string           `json:"name,omitempty"`
	UUID            *string           `json:"uuid,omitempty"`
	LastLogin       *time.Time        `json:"last_login,omitempty"`
	LastLoginServer *string           `json:"last_login_server,omitempty"`
	LastLogout      *time.Time        `json:"last_logout,omitempty"`
	Banned          *bool             `json:"banned,omitempty"`
	Bowspleef       *Bowspleef        `json:"bowspleef,omitempty"`
	CP              *CP               `json:"cp,omitempty"`
	CTW             *CTW              `json:"ctw,omitempty"`
	Factions        *Factions         `json:"factions,omitempty"`
	Kills           *[]Stats          `json:"kills,omitempty"`
	Deaths          *[]Stats          `json:"deaths,omitempty"`
	Objective       *Objective        `json:"objective,omitempty"`
	PlayersMatches  *[]PlayersMatches `json:"matches,omitempty"`
	Paintball       *Paintball        `json:"paintball,omitempty"`
	Splatt          *Splatt           `json:"splatt,omitempty"`
	Teampvp         *Teampvp          `json:"teampvp,omitempty"`
	Tntrun          *Tntrun           `json:"tntrun,omitempty"`
	Total           *Total            `json:"total,omitempty"`
	PermalinkURL    *string           `json:"permalink_url,omitempty"`
}

type Bowspleef struct {
	LoseCount  *int `json:"lose_count,omitempty"`
	PlayCount  *int `json:"play_count,omitempty"`
	TimePlayed *int `json:"time_played,omitempty"`
	WinCount   *int `json:"win_count,omitempty"`
}

type CP struct {
	DeathCount    *int `json:"death_count,omitempty"`
	EnvdeathCount *int `json:"envdeath_count,omitempty"`
	KillCount     *int `json:"kill_count,omitempty"`
	LoseCount     *int `json:"lose_count,omitempty"`
	PlayCount     *int `json:"play_count,omitempty"`
	TimePlayed    *int `json:"time_played,omitempty"`
	WinCount      *int `json:"win_count,omitempty"`
}

type CTW struct {
	DeathCount    *int          `json:"death_count,omitempty"`
	EnvdeathCount *int          `json:"envdeath_count,omitempty"`
	KillCount     *int          `json:"kill_count,omitempty"`
	LoseCount     *int          `json:"lose_count,omitempty"`
	PlayCount     *int          `json:"play_count,omitempty"`
	TimePlayed    *int          `json:"time_played,omitempty"`
	WinCount      *int          `json:"win_count,omitempty"`
	WoolPlaced    *int          `json:"wool_placed,omitempty"`
	WoolPlaces    *[]WoolPlaces `json:"wool_places,omitempty"`
}

type WoolPlaces struct {
	Color     *string    `json:"color,omitempty"`
	Map       *string    `json:"map,omitempty"`
	MatchID   *string    `json:"match_id,omitempty"`
	Team      *string    `json:"team,omitempty"`
	TeamColor *string    `json:"team_color,omitempty"`
	Time      *time.Time `json:"time,omitempty"`
}

type Factions struct {
	Blocks                *Blocks          `json:"blocks,omitempty"`
	DeathCount            *int             `json:"death_count,omitempty"`
	Deaths                *[]FactionsStats `json:"deaths,omitempty"`
	FactionID             *string          `json:"faction_id,omitempty"`
	FactionRole           *string          `json:"faction_role,omitempty"`
	KillCount             *int             `json:"kill_count,omitempty"`
	Kills                 *[]FactionsStats `json:"kills,omitempty"`
	LevelAccumulated      *int             `json:"level_accumulated,omitempty"`
	Power                 *int             `json:"power,omitempty"`
	PvpProtectionDisabled *bool            `json:"pvp_protection_disabled,omitempty"`
	TimePlayed            *int             `json:"time_played,omitempty"`
}

type Blocks struct {
	COALORE     *Block `json:"COAL_ORE,omitempty"`
	DIAMONDORE  *Block `json:"DIAMOND_ORE,omitempty"`
	GOLDORE     *Block `json:"GOLD_ORE,omitempty"`
	IRONORE     *Block `json:"IRON_ORE,omitempty"`
	LAPISORE    *Block `json:"LAPIS_ORE,omitempty"`
	QUARTZORE   *Block `json:"QUARTZ_ORE,omitempty"`
	REDSTONEORE *Block `json:"REDSTONE_ORE,omitempty"`
	STONE       *Block `json:"STONE,omitempty"`
}

type Block struct {
	Break *int `json:"break,omitempty"`
}

type Location struct {
	World *string `json:"world,omitempty"`
	X     *int    `json:"x,omitempty"`
	Y     *int    `json:"y,omitempty"`
	Z     *int    `json:"z,omitempty"`
}

type FactionsStats struct {
	Killer   *string    `json:"killer,omitempty"`
	Location *Location  `json:"location,omitempty"`
	Time     *time.Time `json:"time,omitempty"`
	Victim   *string    `json:"victim,omitempty"`
}

type Stats struct {
	Killer   *string    `json:"killer,omitempty"`
	Location *Location  `json:"location,omitempty"`
	Map      *string    `json:"map,omitempty"`
	MatchID  *string    `json:"match_id,omitempty"`
	Time     *time.Time `json:"time,omitempty"`
	Victim   *string    `json:"victim,omitempty"`
}

type Objective struct {
	CoreLeaked           *int         `json:"core_leaked,omitempty"`
	CoreLeaks            *[]CoreLeaks `json:"core_leaks,omitempty"`
	DeathCount           *int         `json:"death_count,omitempty"`
	DestroyableDestroyed *int         `json:"destroyable_destroyed,omitempty"`
	Destroys             *[]Destroys  `json:"destroys,omitempty"`
	EnvdeathCount        *int         `json:"envdeath_count,omitempty"`
	KillCount            *int         `json:"kill_count,omitempty"`
	LoseCount            *int         `json:"lose_count,omitempty"`
	PlayCount            *int         `json:"play_count,omitempty"`
	TimePlayed           *int         `json:"time_played,omitempty"`
	WinCount             *int         `json:"win_count,omitempty"`
}

type CoreLeaks struct {
	Map       *string    `json:"map,omitempty"`
	MatchID   *string    `json:"match_id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Team      *string    `json:"team,omitempty"`
	TeamColor *string    `json:"team_color,omitempty"`
	Time      *time.Time `json:"time,omitempty"`
}

type Destroys struct {
	Map       *string    `json:"map,omitempty"`
	MatchID   *string    `json:"match_id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Percent   *int       `json:"percent,omitempty"`
	Team      *string    `json:"team,omitempty"`
	TeamColor *string    `json:"team_color,omitempty"`
	Time      *time.Time `json:"time,omitempty"`
}

type PlayersMatches struct {
	DeathCount              *int       `json:"death_count,omitempty"`
	EnvdeathCount           *int       `json:"envdeath_count,omitempty"`
	Finished                *time.Time `json:"finished,omitempty"`
	Gamemode                *string    `json:"gamemode,omitempty"`
	GlickoKillDeviationDiff *float64   `json:"glicko_kill_deviation_diff,omitempty,omitempty"`
	GlickoKillRatingDiff    *float64   `json:"glicko_kill_rating_diff,omitempty,omitempty"`
	ID                      *string    `json:"id,omitempty"`
	KillCount               *int       `json:"kill_count,omitempty"`
	Map                     *string    `json:"map,omitempty"`
	Result                  *string    `json:"result,omitempty"`
	Score                   *int       `json:"score,omitempty"`
	Started                 *time.Time `json:"started,omitempty"`
	Time                    *time.Time `json:"time,omitempty"`
}

type Paintball struct {
	DeathCount *int `json:"death_count,omitempty"`
	HitCount   *int `json:"hit_count,omitempty"`
	KillCount  *int `json:"kill_count,omitempty"`
	LoseCount  *int `json:"lose_count,omitempty"`
	PlayCount  *int `json:"play_count,omitempty"`
	ShotCount  *int `json:"shot_count,omitempty"`
	TimePlayed *int `json:"time_played,omitempty"`
	WinCount   *int `json:"win_count,omitempty"`
}

type Splatt struct {
	DeathCount *int `json:"death_count,omitempty"`
	HitCount   *int `json:"hit_count,omitempty"`
	KillCount  *int `json:"kill_count,omitempty"`
	LoseCount  *int `json:"lose_count,omitempty"`
	PaintCount *int `json:"paint_count,omitempty"`
	PlayCount  *int `json:"play_count,omitempty"`
	ShotCount  *int `json:"shot_count,omitempty"`
	TimePlayed *int `json:"time_played,omitempty"`
	WinCount   *int `json:"win_count,omitempty"`
}

type Teampvp struct {
	DeathCount             *int     `json:"death_count,omitempty"`
	EnvdeathCount          *int     `json:"envdeath_count,omitempty"`
	KdRatio                *float64 `json:"kd_ratio,omitempty"`
	KdRatioRank            *int     `json:"kd_ratio_rank,omitempty"`
	KillCount              *int     `json:"kill_count,omitempty"`
	KillCountRank          *int     `json:"kill_count_rank,omitempty"`
	KkRatio                *float64 `json:"kk_ratio,omitempty"`
	KkRatioRank            *int     `json:"kk_ratio_rank,omitempty"`
	LoseCount              *int     `json:"lose_count,omitempty"`
	ObjectiveCompleted     *int     `json:"objective_completed,omitempty"`
	ObjectiveCompletedRank *int     `json:"objective_completed_rank,omitempty"`
	PlayCount              *int     `json:"play_count,omitempty"`
	Score                  *int     `json:"score,omitempty"`
	TimePlayed             *int     `json:"time_played,omitempty"`
	TimePlayedRank         *int     `json:"time_played_rank,omitempty"`
	WinCount               *int     `json:"win_count,omitempty"`
	WinCountRank           *int     `json:"win_count_rank,omitempty"`
}

type Total struct {
	DeathCount      *int `json:"death_count,omitempty"`
	EnvdeathCount   *int `json:"envdeath_count,omitempty"`
	HitCount        *int `json:"hit_count,omitempty"`
	KillCount       *int `json:"kill_count,omitempty"`
	LoseCount       *int `json:"lose_count,omitempty"`
	NyanAccumulated *int `json:"nyan_accumulated,omitempty"`
	PlayCount       *int `json:"play_count,omitempty"`
	ShotCount       *int `json:"shot_count,omitempty"`
	TimePlayed      *int `json:"time_played,omitempty"`
	WinCount        *int `json:"win_count,omitempty"`
}

type Tntrun struct {
	LoseCount  *int `json:"lose_count,omitempty"`
	PlayCount  *int `json:"play_count,omitempty"`
	TimePlayed *int `json:"time_played,omitempty"`
	WinCount   *int `json:"win_count,omitempty"`
}

type PlayersOptions struct {
	Id     *string
	Fields *string
}

func (session Session) GetPlayers(options PlayersOptions) (*Players, error) {
	if options.Id == nil {
		return nil, fmt.Errorf("options.Id is empty")
	}
	values := url.Values{}
	if options.Fields != nil {
		values.Add("fields", *options.Fields)
	}

	body, err := session.request("players/"+*options.Id, values)

	if err != nil {
		return nil, err
	}

	players := Players{}

	err = json.Unmarshal(body, &players)

	return &players, err
}

package model

import (
	"fmt"
	"time"

	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/protobuf/ptypes"
)

type MatchFacts struct {
	Id               string      `json:"id,omitempty"`
	CompId           string      `json:"comp_id,omitempty"`
	FormattedDate    string      `json:"formatted_date,omitempty"`
	Season           string      `json:"season,omitempty"`
	Week             string      `json:"week,omitempty"`
	Venue            string      `json:"venue,omitempty"`
	VenueId          string      `json:"venue_id,omitempty"`
	VenueCity        string      `json:"venue_city,omitempty"`
	VenueLatitude    string      `json:"venue_latitude,omitempty"`
	VenueLongitude   string      `json:"venue_longitude,omitempty"`
	VenueCountry     string      `json:"venue_country,omitempty"`
	Status           string      `json:"status,omitempty"`
	Timer            string      `json:"timer,omitempty"`
	Time             string      `json:"time,omitempty"`
	LocalTeamId      string      `json:"localteam_id,omitempty"`
	LocalTeamName    string      `json:"localteam_name,omitempty"`
	LocalTeamScore   string      `json:"localteam_score,omitempty"`
	VisitorTeamId    string      `json:"visitorteam_id,omitempty"`
	VisitorTeamName  string      `json:"visitorteam_name,omitempty"`
	VisitorTeamScore string      `json:"visitorteam_score,omitempty"`
	HtScore          string      `json:"ht_score,omitempty"`
	FtScore          string      `json:"ft_score,omitempty"`
	EtScore          string      `json:"et_score,omitempty"`
	PenaltyLocal     string      `json:"penalty_local,omitempty"`
	PenaltyVisitor   string      `json:"penalty_visitor,omitempty"`
	Events           []*Event    `json:"events,omitempty"`
	Commentary       *Commentary `json:"commentary,omitempty"`
	MatchDate        time.Time   `json:"matchDate, omitempty"`
}

func (m *MatchFacts) GetDateTime() time.Time {
	layout := "02.01.2006T15:04"
	str := fmt.Sprintf("%sT%s", m.FormattedDate, m.Time)
	t, err := time.Parse(layout, str)

	if err != nil {
		t = time.Date(2020, 6, 10, 12, 0, 0, 0, time.Now().Location())
	}

	return t
}

func MatchFactsFromGrpc(matchFacts *model.MatchFacts) *MatchFacts {
	t, _ := ptypes.Timestamp(matchFacts.MatchDate)

	var events []*Event
	for _, e := range matchFacts.Events {
		events = append(events, EventsFromGrpc(e))
	}

	return &MatchFacts{
		Id:               matchFacts.Id,
		CompId:           matchFacts.CompId,
		FormattedDate:    matchFacts.FormattedDate,
		Season:           matchFacts.Season,
		Week:             matchFacts.Week,
		Venue:            matchFacts.Venue,
		VenueId:          matchFacts.VenueId,
		VenueCity:        matchFacts.VenueCity,
		VenueLatitude:    matchFacts.VenueLatitude,
		VenueLongitude:   matchFacts.VenueLongitude,
		VenueCountry:     matchFacts.VenueCountry,
		Status:           matchFacts.Status,
		Timer:            matchFacts.Timer,
		Time:             matchFacts.Time,
		LocalTeamId:      matchFacts.LocalTeamId,
		LocalTeamName:    matchFacts.LocalTeamName,
		LocalTeamScore:   matchFacts.LocalTeamScore,
		VisitorTeamId:    matchFacts.VisitorTeamId,
		VisitorTeamName:  matchFacts.VisitorTeamName,
		VisitorTeamScore: matchFacts.VisitorTeamScore,
		HtScore:          matchFacts.HtScore,
		FtScore:          matchFacts.FtScore,
		EtScore:          matchFacts.EtScore,
		PenaltyLocal:     matchFacts.PenaltyLocal,
		PenaltyVisitor:   matchFacts.PenaltyVisitor,
		Events:           events,
		Commentary:       CommentaryFromGrpc(matchFacts.Commentary),
		MatchDate:        t,
	}
}

func MatchFactsToGrpc(matchFacts *MatchFacts) *model.MatchFacts {
	t, _ := ptypes.TimestampProto(matchFacts.MatchDate)

	var events []*model.Event
	for _, e := range matchFacts.Events {
		events = append(events, EventsToGrpc(e))
	}

	return &model.MatchFacts{
		Id:               matchFacts.Id,
		CompId:           matchFacts.CompId,
		FormattedDate:    matchFacts.FormattedDate,
		Season:           matchFacts.Season,
		Week:             matchFacts.Week,
		Venue:            matchFacts.Venue,
		VenueId:          matchFacts.VenueId,
		VenueCity:        matchFacts.VenueCity,
		VenueLatitude:    matchFacts.VenueLatitude,
		VenueLongitude:   matchFacts.VenueLongitude,
		VenueCountry:     matchFacts.VenueCountry,
		Status:           matchFacts.Status,
		Timer:            matchFacts.Timer,
		Time:             matchFacts.Time,
		LocalTeamId:      matchFacts.LocalTeamId,
		LocalTeamName:    matchFacts.LocalTeamName,
		LocalTeamScore:   matchFacts.LocalTeamScore,
		VisitorTeamId:    matchFacts.VisitorTeamId,
		VisitorTeamName:  matchFacts.VisitorTeamName,
		VisitorTeamScore: matchFacts.VisitorTeamScore,
		HtScore:          matchFacts.HtScore,
		FtScore:          matchFacts.FtScore,
		EtScore:          matchFacts.EtScore,
		PenaltyLocal:     matchFacts.PenaltyLocal,
		PenaltyVisitor:   matchFacts.PenaltyVisitor,
		Events:           events,
		Commentary:       CommentaryToGrpc(matchFacts.Commentary),
		MatchDate:        t,
	}
}

type Event struct {
	Id       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Result   string `json:"result,omitempty"`
	Minute   string `json:"minute,omitempty"`
	ExtraMin string `json:"extra_min,omitempty"`
	Team     string `json:"team,omitempty"`
	Player   string `json:"player,omitempty"`
	PlayerId string `json:"player_id,omitempty"`
	Assist   string `json:"assist,omitempty"`
	AssistId string `json:"assist_id,omitempty"`
}

func EventsFromGrpc(event *model.Event) *Event {
	if event == nil {
		return nil
	}

	return &Event{
		Id:       event.Id,
		Type:     event.Type,
		Result:   event.Result,
		Minute:   event.Minute,
		ExtraMin: event.ExtraMin,
		Team:     event.Team,
		Player:   event.Player,
		PlayerId: event.PlayerId,
		Assist:   event.Assist,
		AssistId: event.AssistId,
	}
}

func EventsToGrpc(event *Event) *model.Event {
	if event == nil {
		return nil
	}

	return &model.Event{
		Id:       event.Id,
		Type:     event.Type,
		Result:   event.Result,
		Minute:   event.Minute,
		ExtraMin: event.ExtraMin,
		Team:     event.Team,
		Player:   event.Player,
		PlayerId: event.PlayerId,
		Assist:   event.Assist,
		AssistId: event.AssistId,
	}
}

type Commentary struct {
	MatchId       string         `json:"match_id,omitempty"`
	MatchInfo     *MatchInfo     `json:"match_info,omitempty"`
	Lineup        *Lineup        `json:"lineup,omitempty"`
	Subs          *Lineup        `json:"subs,omitempty"`
	Substitutions *Substitutions `json:"substitutions,omitempty"`
	Comments      []*Comment     `json:"comments,omitempty"`
	MatchStats    *MatchStats    `json:"match_stats,omitempty"`
	PlayerStats   *PlayerStats   `json:"player_stats,omitempty"`
}

func CommentaryFromGrpc(commentary *model.Commentary) *Commentary {
	if commentary == nil {
		return nil
	}

	var comments []*Comment
	for _, c := range commentary.Comments {
		comments = append(comments, CommentFromGrpc(c))
	}

	return &Commentary{
		MatchId:       commentary.MatchId,
		MatchInfo:     MatchInfoFromGrpc(commentary.MatchInfo),
		Lineup:        LineupFromGrpc(commentary.Lineup),
		Subs:          LineupFromGrpc(commentary.Subs),
		Substitutions: SubstitutionsFromGrpc(commentary.Substitutions),
		Comments:      comments,
		MatchStats:    MatchStatsFromGrpc(commentary.MatchStats),
		PlayerStats:   PlayerStatsFromGrpc(commentary.PlayerStats),
	}
}

func CommentaryToGrpc(commentary *Commentary) *model.Commentary {
	if commentary == nil {
		return nil
	}

	var comments []*model.Comment
	for _, c := range commentary.Comments {
		comments = append(comments, CommentToGrpc(c))
	}

	return &model.Commentary{
		MatchId:       commentary.MatchId,
		MatchInfo:     MatchInfoToGrpc(commentary.MatchInfo),
		Lineup:        LineupToGrpc(commentary.Lineup),
		Subs:          LineupToGrpc(commentary.Subs),
		Substitutions: SubstitutionsToGrpc(commentary.Substitutions),
		Comments:      comments,
		MatchStats:    MatchStatsToGrpc(commentary.MatchStats),
		PlayerStats:   PlayerStatsToGrpc(commentary.PlayerStats),
	}
}

type MatchInfo struct {
	Stadium    string `json:"stadium,omitempty"`
	Attendance string `json:"attendance,omitempty"`
	Referee    string `json:"referee,omitempty"`
}

func MatchInfoFromGrpc(matchInfo *model.MatchInfo) *MatchInfo {
	return &MatchInfo{
		Stadium:    matchInfo.GetStadium(),
		Attendance: matchInfo.GetAttendance(),
		Referee:    matchInfo.GetReferee(),
	}
}

func MatchInfoToGrpc(matchInfo *MatchInfo) *model.MatchInfo {
	if matchInfo == nil {
		return nil
	}

	return &model.MatchInfo{
		Stadium:    matchInfo.Stadium,
		Attendance: matchInfo.Attendance,
		Referee:    matchInfo.Referee,
	}
}

type Lineup struct {
	LocalTeam   []*Position `json:"localteam,omitempty"`
	VisitorTeam []*Position `json:"visitorteam,omitempty"`
}

func LineupFromGrpc(lineup *model.Lineup) *Lineup {
	if lineup == nil {
		return nil
	}

	var local []*Position
	for _, t := range lineup.LocalTeam {
		local = append(local, PositionFromGrpc(t))
	}

	var visitor []*Position
	for _, t := range lineup.VisitorTeam {
		visitor = append(visitor, PositionFromGrpc(t))
	}

	return &Lineup{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

func LineupToGrpc(lineup *Lineup) *model.Lineup {
	if lineup == nil {
		return nil
	}

	var local []*model.Position
	for _, t := range lineup.LocalTeam {
		local = append(local, PositionToGrpc(t))
	}

	var visitor []*model.Position
	for _, t := range lineup.VisitorTeam {
		visitor = append(visitor, PositionToGrpc(t))
	}

	return &model.Lineup{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

type Position struct {
	Id     string `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
	Pos    string `json:"pos,omitempty"`
}

func PositionFromGrpc(position *model.Position) *Position {
	if position == nil {
		return nil
	}

	return &Position{
		Id:     position.Id,
		Number: position.Number,
		Name:   position.Name,
		Pos:    position.Pos,
	}
}

func PositionToGrpc(position *Position) *model.Position {
	if position == nil {
		return nil
	}

	return &model.Position{
		Id:     position.Id,
		Number: position.Number,
		Name:   position.Name,
		Pos:    position.Pos,
	}
}

type Substitutions struct {
	LocalTeam   []*Substitution `json:"localteam,omitempty"`
	VisitorTeam []*Substitution `json:"visitorteam,omitempty"`
}

func SubstitutionsFromGrpc(substitutions *model.Substitutions) *Substitutions {
	if substitutions == nil {
		return nil
	}

	var local []*Substitution
	for _, t := range substitutions.LocalTeam {
		local = append(local, SubstitutionFromGrpc(t))
	}

	var visitor []*Substitution
	for _, t := range substitutions.VisitorTeam {
		visitor = append(visitor, SubstitutionFromGrpc(t))
	}

	return &Substitutions{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

func SubstitutionsToGrpc(substitutions *Substitutions) *model.Substitutions {
	if substitutions == nil {
		return nil
	}

	var local []*model.Substitution
	for _, t := range substitutions.LocalTeam {
		local = append(local, SubstitutionToGrpc(t))
	}

	var visitor []*model.Substitution
	for _, t := range substitutions.VisitorTeam {
		visitor = append(visitor, SubstitutionToGrpc(t))
	}

	return &model.Substitutions{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

type Substitution struct {
	OffName string `json:"off_name,omitempty"`
	OnName  string `json:"on_name,omitempty"`
	OffId   string `json:"off_id,omitempty"`
	OnId    string `json:"on_id,omitempty"`
	Minute  string `json:"minute,omitempty"`
	TableId string `json:"table_id,omitempty"`
}

func SubstitutionFromGrpc(substitution *model.Substitution) *Substitution {
	if substitution == nil {
		return nil
	}

	return &Substitution{
		OffName: substitution.OffName,
		OnName:  substitution.OnName,
		OffId:   substitution.OffId,
		OnId:    substitution.OnId,
		Minute:  substitution.Minute,
		TableId: substitution.TableId,
	}
}

func SubstitutionToGrpc(substitution *Substitution) *model.Substitution {
	if substitution == nil {
		return nil
	}

	return &model.Substitution{
		OffName: substitution.OffName,
		OnName:  substitution.OnName,
		OffId:   substitution.OffId,
		OnId:    substitution.OnId,
		Minute:  substitution.Minute,
		TableId: substitution.TableId,
	}
}

type Comment struct {
	Id        string `json:"id,omitempty"`
	Important string `json:"important,omitempty"`
	Goal      string `json:"isgoal,omitempty"`
	Minute    string `json:"minute,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

func CommentFromGrpc(comment *model.Comment) *Comment {
	if comment == nil {
		return nil
	}

	return &Comment{
		Id:        comment.Id,
		Important: comment.Important,
		Goal:      comment.Goal,
		Minute:    comment.Minute,
		Comment:   comment.Comment,
	}
}

func CommentToGrpc(comment *Comment) *model.Comment {
	if comment == nil {
		return nil
	}

	return &model.Comment{
		Id:        comment.Id,
		Important: comment.Important,
		Goal:      comment.Goal,
		Minute:    comment.Minute,
		Comment:   comment.Comment,
	}
}

type MatchStats struct {
	LocalTeam   []*TeamStats `json:"localteam,omitempty"`
	VisitorTeam []*TeamStats `json:"visitorteam,omitempty"`
}

func MatchStatsFromGrpc(matchStats *model.MatchStats) *MatchStats {
	if matchStats == nil {
		return nil
	}

	var local []*TeamStats
	for _, t := range matchStats.LocalTeam {
		local = append(local, TeamStatsFromGrpc(t))
	}

	var visitor []*TeamStats
	for _, t := range matchStats.VisitorTeam {
		visitor = append(visitor, TeamStatsFromGrpc(t))
	}

	return &MatchStats{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

func MatchStatsToGrpc(matchStats *MatchStats) *model.MatchStats {
	if matchStats == nil {
		return nil
	}

	var local []*model.TeamStats
	for _, t := range matchStats.LocalTeam {
		local = append(local, TeamStatsToGrpc(t))
	}

	var visitor []*model.TeamStats
	for _, t := range matchStats.VisitorTeam {
		visitor = append(visitor, TeamStatsToGrpc(t))
	}

	return &model.MatchStats{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

type TeamStats struct {
	ShotsTotal     string `json:"shots_total,omitempty"`
	ShotsOnGoal    string `json:"shots_ongoal,omitempty"`
	Fouls          string `json:"fouls,omitempty"`
	Corners        string `json:"corners,omitempty"`
	Offsides       string `json:"offsides,omitempty"`
	PossessionTime string `json:"possesiontime,omitempty"`
	YellowCards    string `json:"yellowcards,omitempty"`
	RedCards       string `json:"redcards,omitempty"`
	Saves          string `json:"saves,omitempty"`
	TableId        string `json:"table_id,omitempty"`
}

func TeamStatsFromGrpc(teamStats *model.TeamStats) *TeamStats {
	if teamStats == nil {
		return nil
	}

	return &TeamStats{
		ShotsTotal:     teamStats.ShotsTotal,
		ShotsOnGoal:    teamStats.ShotsOnGoal,
		Fouls:          teamStats.Fouls,
		Corners:        teamStats.Corners,
		Offsides:       teamStats.Offsides,
		PossessionTime: teamStats.PossessionTime,
		YellowCards:    teamStats.YellowCards,
		RedCards:       teamStats.RedCards,
		Saves:          teamStats.Saves,
		TableId:        teamStats.TableId,
	}
}

func TeamStatsToGrpc(teamStats *TeamStats) *model.TeamStats {
	if teamStats == nil {
		return nil
	}

	return &model.TeamStats{
		ShotsTotal:     teamStats.ShotsTotal,
		ShotsOnGoal:    teamStats.ShotsOnGoal,
		Fouls:          teamStats.Fouls,
		Corners:        teamStats.Corners,
		Offsides:       teamStats.Offsides,
		PossessionTime: teamStats.PossessionTime,
		YellowCards:    teamStats.YellowCards,
		RedCards:       teamStats.RedCards,
		Saves:          teamStats.Saves,
		TableId:        teamStats.TableId,
	}
}

type PlayerStats struct {
	LocalTeam   *Players `json:"localteam,omitempty"`
	VisitorTeam *Players `json:"visitorteam,omitempty"`
}

func PlayerStatsFromGrpc(playerStats *model.PlayerStats) *PlayerStats {
	if playerStats == nil {
		return nil
	}

	return &PlayerStats{
		LocalTeam:   PlayersFromGrpc(playerStats.LocalTeam),
		VisitorTeam: PlayersFromGrpc(playerStats.VisitorTeam),
	}
}

func PlayerStatsToGrpc(playerStats *PlayerStats) *model.PlayerStats {
	if playerStats == nil {
		return nil
	}

	return &model.PlayerStats{
		LocalTeam:   PlayersToGrpc(playerStats.LocalTeam),
		VisitorTeam: PlayersToGrpc(playerStats.VisitorTeam),
	}
}

type Players struct {
	Player []*Player `json:"player,omitempty"`
}

func PlayersFromGrpc(players *model.Players) *Players {
	if players == nil {
		return nil
	}

	var plyrs []*Player
	for _, p := range players.Player {
		plyrs = append(plyrs, PlayerFromGrpc(p))
	}

	return &Players{
		Player: plyrs,
	}
}

func PlayersToGrpc(players *Players) *model.Players {
	if players == nil {
		return nil
	}

	var plyrs []*model.Player
	for _, p := range players.Player {
		plyrs = append(plyrs, PlayerToGrpc(p))
	}

	return &model.Players{
		Player: plyrs,
	}
}

type Player struct {
	Id             string `json:"id,omitempty"`
	Num            string `json:"num,omitempty"`
	Name           string `json:"name,omitempty"`
	Pos            string `json:"pos,omitempty"`
	PosX           string `json:"posx,omitempty"`
	PosY           string `json:"posy,omitempty"`
	ShotsTotal     string `json:"shots_total,omitempty"`
	ShotsOnGoal    string `json:"shots_on_goal,omitempty"`
	Goals          string `json:"goals,omitempty"`
	Assists        string `json:"assists,omitempty"`
	Offsides       string `json:"offsides,omitempty"`
	FoulsDrawn     string `json:"fouls_drawn,omitempty"`
	FoulsCommitted string `json:"fouls_committed,omitempty"`
	Saves          string `json:"saves,omitempty"`
	YellowCards    string `json:"yellowcards,omitempty"`
	RedCards       string `json:"redcards,omitempty"`
	PenScore       string `json:"pen_score,omitempty"`
	PenMiss        string `json:"pen_miss,omitempty"`
}

func PlayerFromGrpc(player *model.Player) *Player {
	if player == nil {
		return nil
	}

	return &Player{
		Id:             player.Id,
		Num:            player.Num,
		Name:           player.Name,
		Pos:            player.Pos,
		PosX:           player.PosX,
		PosY:           player.PosY,
		ShotsTotal:     player.ShotsTotal,
		ShotsOnGoal:    player.ShotsOnGoal,
		Goals:          player.Goals,
		Assists:        player.Assists,
		Offsides:       player.Offsides,
		FoulsDrawn:     player.FoulsDrawn,
		FoulsCommitted: player.FoulsCommitted,
		Saves:          player.Saves,
		YellowCards:    player.YellowCards,
		RedCards:       player.RedCards,
		PenScore:       player.PenScore,
		PenMiss:        player.PenMiss,
	}
}

func PlayerToGrpc(player *Player) *model.Player {
	if player == nil {
		return nil
	}

	return &model.Player{
		Id:             player.Id,
		Num:            player.Num,
		Name:           player.Name,
		Pos:            player.Pos,
		PosX:           player.PosX,
		PosY:           player.PosY,
		ShotsTotal:     player.ShotsTotal,
		ShotsOnGoal:    player.ShotsOnGoal,
		Goals:          player.Goals,
		Assists:        player.Assists,
		Offsides:       player.Offsides,
		FoulsDrawn:     player.FoulsDrawn,
		FoulsCommitted: player.FoulsCommitted,
		Saves:          player.Saves,
		YellowCards:    player.YellowCards,
		RedCards:       player.RedCards,
		PenScore:       player.PenScore,
		PenMiss:        player.PenMiss,
	}
}

type MatchFactsSlice []MatchFacts

func (m MatchFactsSlice) Len() int {
	return len(m)
}

func (m MatchFactsSlice) Less(i, j int) bool {
	return m[i].GetDateTime().Before(m[j].GetDateTime())
}

func (m MatchFactsSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

package mongo

import (
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"time"
)

type matchFactsEntity struct {
	Id               string      `bson:"_id,omitempty"`
	CompId           string      `bson:"compId,omitempty"`
	FormattedDate    string      `bson:"formattedDate,omitempty"`
	Season           string      `bson:"season,omitempty"`
	Week             string      `bson:"week,omitempty"`
	Venue            string      `bson:"venue,omitempty"`
	VenueId          string      `bson:"venueId,omitempty"`
	VenueCity        string      `bson:"venueCity,omitempty"`
	Status           string      `bson:"status,omitempty"`
	Timer            string      `bson:"timer,omitempty"`
	Time             string      `bson:"time,omitempty"`
	LocalTeamId      string      `bson:"localTeamId,omitempty"`
	LocalTeamName    string      `bson:"localTeamName,omitempty"`
	LocalTeamScore   string      `bson:"localTeamScore,omitempty"`
	VisitorTeamId    string      `bson:"visitorTeamId,omitempty"`
	VisitorTeamName  string      `bson:"visitorTeamName,omitempty"`
	VisitorTeamScore string      `bson:"visitorTeamScore,omitempty"`
	HtScore          string      `bson:"htScore,omitempty"`
	FtScore          string      `bson:"ftScore,omitempty"`
	EtScore          string      `bson:"etScore,omitempty"`
	PenaltyLocal     string      `bson:"penaltyLocal,omitempty"`
	PenaltyVisitor   string      `bson:"penaltyVisitor,omitempty"`
	Events           []*event    `bson:"events,omitempty"`
	Commentary       *commentary `bson:"commentary,omitempty"`
	MatchDate        time.Time   `bson:"matchDate, omitempty"`
}

func fromMatchFacts(matchFacts *model.MatchFacts) *matchFactsEntity {
	var events []*event
	for _, e := range matchFacts.Events {
		events = append(events, fromEvent(e))
	}

	return &matchFactsEntity{
		Id:               matchFacts.Id,
		CompId:           matchFacts.CompId,
		FormattedDate:    matchFacts.FormattedDate,
		Season:           matchFacts.Season,
		Week:             matchFacts.Week,
		Venue:            matchFacts.Venue,
		VenueId:          matchFacts.VenueId,
		VenueCity:        matchFacts.VenueCity,
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
		Commentary:       fromCommentary(matchFacts.Commentary),
		MatchDate:        matchFacts.MatchDate,
	}
}

func toMatchFacts(matchFacts *matchFactsEntity) *model.MatchFacts {
	var events []*model.Event
	for _, e := range matchFacts.Events {
		events = append(events, toEvent(e))
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
		Commentary:       toCommentary(matchFacts.Commentary),
		MatchDate:        matchFacts.MatchDate,
	}
}

type event struct {
	Id       string `bson:"id,omitempty"`
	Type     string `bson:"type,omitempty"`
	Result   string `bson:"result,omitempty"`
	Minute   string `bson:"minute,omitempty"`
	ExtraMin string `bson:"extraMin,omitempty"`
	Team     string `bson:"team,omitempty"`
	Player   string `bson:"player,omitempty"`
	PlayerId string `bson:"playerId,omitempty"`
	Assist   string `bson:"assist,omitempty"`
	AssistId string `bson:"assistId,omitempty"`
}

func fromEvent(e *model.Event) *event {
	if e == nil {
		return nil
	}

	return &event{
		Id:       e.Id,
		Type:     e.Type,
		Result:   e.Result,
		Minute:   e.Minute,
		ExtraMin: e.ExtraMin,
		Team:     e.Team,
		Player:   e.Player,
		PlayerId: e.PlayerId,
		Assist:   e.Assist,
		AssistId: e.AssistId,
	}
}

func toEvent(e *event) *model.Event {
	if e == nil {
		return nil
	}

	return &model.Event{
		Id:       e.Id,
		Type:     e.Type,
		Result:   e.Result,
		Minute:   e.Minute,
		ExtraMin: e.ExtraMin,
		Team:     e.Team,
		Player:   e.Player,
		PlayerId: e.PlayerId,
		Assist:   e.Assist,
		AssistId: e.AssistId,
	}
}

type commentary struct {
	MatchId       string         `bson:"matchId,omitempty"`
	MatchInfo     []*matchInfo   `bson:"matchInfo,omitempty"`
	Lineup        *lineup        `bson:"lineup,omitempty"`
	Subs          *lineup        `bson:"subs,omitempty"`
	Substitutions *substitutions `bson:"substitutions,omitempty"`
	Comments      []*comment     `bson:"comments,omitempty"`
	MatchStats    *matchStats    `bson:"matchStats,omitempty"`
	PlayerStats   *playerStats   `bson:"playerStats,omitempty"`
}

func fromCommentary(c *model.Commentary) *commentary {
	if c == nil {
		return nil
	}

	var matchInfo []*matchInfo
	for _, m := range c.MatchInfo {
		matchInfo = append(matchInfo, fromMatchInfo(m))
	}

	var comments []*comment
	for _, cmt := range c.Comments {
		comments = append(comments, fromComment(cmt))
	}

	return &commentary{
		MatchId:       c.MatchId,
		MatchInfo:     matchInfo,
		Lineup:        fromLineup(c.Lineup),
		Subs:          fromLineup(c.Subs),
		Substitutions: fromSubstitutions(c.Substitutions),
		Comments:      comments,
		MatchStats:    fromMatchStats(c.MatchStats),
		PlayerStats:   fromPlayerStats(c.PlayerStats),
	}
}

func toCommentary(c *commentary) *model.Commentary {
	if c == nil {
		return nil
	}

	var matchInfo []*model.MatchInfo
	for _, m := range c.MatchInfo {
		matchInfo = append(matchInfo, toMatchInfo(m))
	}

	var comments []*model.Comment
	for _, cmt := range c.Comments {
		comments = append(comments, toComment(cmt))
	}

	return &model.Commentary{
		MatchId:       c.MatchId,
		MatchInfo:     matchInfo,
		Lineup:        toLineup(c.Lineup),
		Subs:          toLineup(c.Subs),
		Substitutions: toSubstitutions(c.Substitutions),
		Comments:      comments,
		MatchStats:    toMatchStats(c.MatchStats),
		PlayerStats:   toPlayerStats(c.PlayerStats),
	}
}

type matchInfo struct {
	Stadium    string `bson:"stadium,omitempty"`
	Attendance string `bson:"attendance,omitempty"`
	Referee    string `bson:"referee,omitempty"`
}

func fromMatchInfo(m *model.MatchInfo) *matchInfo {
	if m == nil {
		return nil
	}

	return &matchInfo{
		Stadium:    m.Stadium,
		Attendance: m.Attendance,
		Referee:    m.Referee,
	}
}

func toMatchInfo(m *matchInfo) *model.MatchInfo {
	if m == nil {
		return nil
	}

	return &model.MatchInfo{
		Stadium:    m.Stadium,
		Attendance: m.Attendance,
		Referee:    m.Referee,
	}
}

type lineup struct {
	LocalTeam   []*position `bson:"localTeam,omitempty"`
	VisitorTeam []*position `bson:"visitorTeam,omitempty"`
}

func fromLineup(l *model.Lineup) *lineup {
	if l == nil {
		return nil
	}

	var local []*position
	for _, t := range l.LocalTeam {
		local = append(local, fromPosition(t))
	}

	var visitor []*position
	for _, t := range l.VisitorTeam {
		visitor = append(visitor, fromPosition(t))
	}

	return &lineup{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

func toLineup(l *lineup) *model.Lineup {
	if l == nil {
		return nil
	}

	var local []*model.Position
	for _, t := range l.LocalTeam {
		local = append(local, toPosition(t))
	}

	var visitor []*model.Position
	for _, t := range l.VisitorTeam {
		visitor = append(visitor, toPosition(t))
	}

	return &model.Lineup{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

type position struct {
	Id     string `bson:"id,omitempty"`
	Number string `bson:"number,omitempty"`
	Name   string `bson:"name,omitempty"`
	Pos    string `bson:"pos,omitempty"`
}

func fromPosition(p *model.Position) *position {
	if p == nil {
		return nil
	}

	return &position{
		Id:     p.Id,
		Number: p.Number,
		Name:   p.Name,
		Pos:    p.Pos,
	}
}

func toPosition(p *position) *model.Position {
	if p == nil {
		return nil
	}

	return &model.Position{
		Id:     p.Id,
		Number: p.Number,
		Name:   p.Name,
		Pos:    p.Pos,
	}
}

type substitutions struct {
	LocalTeam   []*substitution `bson:"localTeam,omitempty"`
	VisitorTeam []*substitution `bson:"visitorTeam,omitempty"`
}

func fromSubstitutions(s *model.Substitutions) *substitutions {
	if s == nil {
		return nil
	}

	var local []*substitution
	for _, t := range s.LocalTeam {
		local = append(local, fromSubstitution(t))
	}

	var visitor []*substitution
	for _, t := range s.VisitorTeam {
		visitor = append(visitor, fromSubstitution(t))
	}

	return &substitutions{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

func toSubstitutions(s *substitutions) *model.Substitutions {
	if s == nil {
		return nil
	}

	var local []*model.Substitution
	for _, t := range s.LocalTeam {
		local = append(local, toSubstitution(t))
	}

	var visitor []*model.Substitution
	for _, t := range s.VisitorTeam {
		visitor = append(visitor, toSubstitution(t))
	}

	return &model.Substitutions{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

type substitution struct {
	OffName string `bson:"offName,omitempty"`
	OnName  string `bson:"onName,omitempty"`
	OffId   string `bson:"offId,omitempty"`
	OnId    string `bson:"onId,omitempty"`
	Minute  string `bson:"minute,omitempty"`
	TableId string `bson:"tableId,omitempty"`
}

func fromSubstitution(s *model.Substitution) *substitution {
	if s == nil {
		return nil
	}

	return &substitution{
		OffName: s.OffName,
		OnName:  s.OnName,
		OffId:   s.OffId,
		OnId:    s.OnId,
		Minute:  s.Minute,
		TableId: s.TableId,
	}
}

func toSubstitution(s *substitution) *model.Substitution {
	if s == nil {
		return nil
	}

	return &model.Substitution{
		OffName: s.OffName,
		OnName:  s.OnName,
		OffId:   s.OffId,
		OnId:    s.OnId,
		Minute:  s.Minute,
		TableId: s.TableId,
	}
}

type comment struct {
	Id        string `bson:"id,omitempty"`
	Important string `bson:"important,omitempty"`
	Goal      string `bson:"isGoal,omitempty"`
	Minute    string `bson:"minute,omitempty"`
	Comment   string `bson:"comment,omitempty"`
}

func fromComment(c *model.Comment) *comment {
	if c == nil {
		return nil
	}

	return &comment{
		Id:        c.Id,
		Important: c.Important,
		Goal:      c.Goal,
		Minute:    c.Minute,
		Comment:   c.Comment,
	}
}

func toComment(c *comment) *model.Comment {
	if c == nil {
		return nil
	}

	return &model.Comment{
		Id:        c.Id,
		Important: c.Important,
		Goal:      c.Goal,
		Minute:    c.Minute,
		Comment:   c.Comment,
	}
}

type matchStats struct {
	LocalTeam   []*teamStats `bson:"localTeam,omitempty"`
	VisitorTeam []*teamStats `bson:"visitorTeam,omitempty"`
}

func fromMatchStats(m *model.MatchStats) *matchStats {
	if m == nil {
		return nil
	}

	var local []*teamStats
	for _, t := range m.LocalTeam {
		local = append(local, fromTeamStats(t))
	}

	var visitor []*teamStats
	for _, t := range m.VisitorTeam {
		visitor = append(visitor, fromTeamStats(t))
	}

	return &matchStats{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

func toMatchStats(m *matchStats) *model.MatchStats {
	if m == nil {
		return nil
	}

	var local []*model.TeamStats
	for _, t := range m.LocalTeam {
		local = append(local, toTeamStats(t))
	}

	var visitor []*model.TeamStats
	for _, t := range m.VisitorTeam {
		visitor = append(visitor, toTeamStats(t))
	}

	return &model.MatchStats{
		LocalTeam:   local,
		VisitorTeam: visitor,
	}
}

type teamStats struct {
	ShotsTotal     string `bson:"shotsTotal,omitempty"`
	ShotsOnGoal    string `bson:"shotsOnGoal,omitempty"`
	Fouls          string `bson:"fouls,omitempty"`
	Corners        string `bson:"corners,omitempty"`
	Offsides       string `bson:"offsides,omitempty"`
	PossessionTime string `bson:"possessionTime,omitempty"`
	YellowCards    string `bson:"yellowCards,omitempty"`
	RedCards       string `bson:"redCards,omitempty"`
	Saves          string `bson:"saves,omitempty"`
	TableId        string `bson:"tableId,omitempty"`
}

func fromTeamStats(t *model.TeamStats) *teamStats {
	if t == nil {
		return nil
	}

	return &teamStats{
		ShotsTotal:     t.ShotsTotal,
		ShotsOnGoal:    t.ShotsOnGoal,
		Fouls:          t.Fouls,
		Corners:        t.Corners,
		Offsides:       t.Offsides,
		PossessionTime: t.PossessionTime,
		YellowCards:    t.YellowCards,
		RedCards:       t.RedCards,
		Saves:          t.Saves,
		TableId:        t.TableId,
	}
}

func toTeamStats(t *teamStats) *model.TeamStats {
	if t == nil {
		return nil
	}

	return &model.TeamStats{
		ShotsTotal:     t.ShotsTotal,
		ShotsOnGoal:    t.ShotsOnGoal,
		Fouls:          t.Fouls,
		Corners:        t.Corners,
		Offsides:       t.Offsides,
		PossessionTime: t.PossessionTime,
		YellowCards:    t.YellowCards,
		RedCards:       t.RedCards,
		Saves:          t.Saves,
		TableId:        t.TableId,
	}
}

type playerStats struct {
	LocalTeam   *players `bson:"localTeam,omitempty"`
	VisitorTeam *players `bson:"visitorTeam,omitempty"`
}

func fromPlayerStats(p *model.PlayerStats) *playerStats {
	if p == nil {
		return nil
	}

	return &playerStats{
		LocalTeam:   fromPlayers(p.LocalTeam),
		VisitorTeam: fromPlayers(p.VisitorTeam),
	}
}

func toPlayerStats(p *playerStats) *model.PlayerStats {
	if p == nil {
		return nil
	}

	return &model.PlayerStats{
		LocalTeam:   toPlayers(p.LocalTeam),
		VisitorTeam: toPlayers(p.VisitorTeam),
	}
}

type players struct {
	Player []*player `bson:"player,omitempty"`
}

func fromPlayers(p *model.Players) *players {
	if p == nil {
		return nil
	}

	var plyrs []*player
	for _, p := range p.Player {
		plyrs = append(plyrs, fromPlayer(p))
	}

	return &players{
		Player: plyrs,
	}
}

func toPlayers(p *players) *model.Players {
	if p == nil {
		return nil
	}

	var plyrs []*model.Player
	for _, p := range p.Player {
		plyrs = append(plyrs, toPlayer(p))
	}

	return &model.Players{
		Player: plyrs,
	}
}

type player struct {
	Id             string `bson:"id,omitempty"`
	Num            string `bson:"num,omitempty"`
	Name           string `bson:"name,omitempty"`
	Pos            string `bson:"pos,omitempty"`
	PosX           string `bson:"posX,omitempty"`
	PosY           string `bson:"posY,omitempty"`
	ShotsTotal     string `bson:"shotsTotal,omitempty"`
	ShotsOnGoal    string `bson:"shotsOnGoal,omitempty"`
	Goals          string `bson:"goals,omitempty"`
	Assists        string `bson:"assists,omitempty"`
	Offsides       string `bson:"offsides,omitempty"`
	FoulsDrawn     string `bson:"foulsDrawn,omitempty"`
	FoulsCommitted string `bson:"foulsCommitted,omitempty"`
	Saves          string `bson:"saves,omitempty"`
	YellowCards    string `bson:"yellowCards,omitempty"`
	RedCards       string `bson:"redCards,omitempty"`
	PenScore       string `bson:"penScore,omitempty"`
	PenMiss        string `bson:"penMiss,omitempty"`
}

func fromPlayer(p *model.Player) *player {
	if p == nil {
		return nil
	}

	return &player{
		Id:             p.Id,
		Num:            p.Num,
		Name:           p.Name,
		Pos:            p.Pos,
		PosX:           p.PosX,
		PosY:           p.PosY,
		ShotsTotal:     p.ShotsTotal,
		ShotsOnGoal:    p.ShotsOnGoal,
		Goals:          p.Goals,
		Assists:        p.Assists,
		Offsides:       p.Offsides,
		FoulsDrawn:     p.FoulsDrawn,
		FoulsCommitted: p.FoulsCommitted,
		Saves:          p.Saves,
		YellowCards:    p.YellowCards,
		RedCards:       p.RedCards,
		PenScore:       p.PenScore,
		PenMiss:        p.PenMiss,
	}
}

func toPlayer(p *player) *model.Player {
	if p == nil {
		return nil
	}

	return &model.Player{
		Id:             p.Id,
		Num:            p.Num,
		Name:           p.Name,
		Pos:            p.Pos,
		PosX:           p.PosX,
		PosY:           p.PosY,
		ShotsTotal:     p.ShotsTotal,
		ShotsOnGoal:    p.ShotsOnGoal,
		Goals:          p.Goals,
		Assists:        p.Assists,
		Offsides:       p.Offsides,
		FoulsDrawn:     p.FoulsDrawn,
		FoulsCommitted: p.FoulsCommitted,
		Saves:          p.Saves,
		YellowCards:    p.YellowCards,
		RedCards:       p.RedCards,
		PenScore:       p.PenScore,
		PenMiss:        p.PenMiss,
	}
}

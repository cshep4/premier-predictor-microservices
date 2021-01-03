import * as t from 'google-protobuf/google/protobuf/timestamp_pb';

export class MatchFacts {
    id!: string;
    compId!: string;
    formattedDate!: string;
    season!: string;
    week!: string;
    venue!: string;
    venueId!: string;
    venueCity!: string;
    venueLatitude!: string;
    venueLongitude!: string;
    venueCountry!: string;
    status!: string;
    timer!: string;
    time!: string;
    localTeamId!: string;
    localTeamName!: string;
    localTeamScore!: string;
    visitorTeamId!: string;
    visitorTeamName!: string;
    visitorTeamScore!: string;
    htScore!: string;
    ftScore!: string;
    etScore!: string;
    penaltyLocal!: string;
    penaltyVisitor!: string;
    events!: Event[];
    commentary!: Commentary;
    matchDate!: any;
}

class Event {
    id!: string;
    type!: string;
    result!: string;
    minute!: string;
    extraMin!: string;
    team!: string;
    player!: string;
    playerId!: string;
    assist!: string;
    assistId!: string;
}

class Commentary {
    matchId!: string;
    matchInfo!: MatchInfo;
    lineup!: Lineup;
    subs!: Lineup;
    substitutions!: Substitutions;
    comments!: Comment[];
    matchStats!: MatchStats;
    playerStats!: PlayerStats;
}

class MatchInfo {
    stadium!: string;
    attendance!: string;
    referee!: string;
}

class Lineup {
    localTeam!: Position[];
    visitorTeam!: Position[];
}

class Position {
    id!: string;
    number!: string;
    name!: string;
    pos!: string;
}

class Substitutions {
    localTeam!: Substitution[];
    visitorTeam!: Substitution[];
}

class Substitution {
    offName!: string;
    onName!: string;
    offId!: string;
    onId!: string;
    minute!: string;
    tableId!: string;
}

class Comment {
    id!: string;
    important!: string;
    goal!: string;
    minute!: string;
    comment!: string;
}

class MatchStats {
    localTeam!: TeamStats[];
    visitorTeam!: TeamStats[];
}

class TeamStats {
    shotsTotal!: string;
    shotsOnGoal!: string;
    fouls!: string;
    corners!: string;
    offsides!: string;
    possessionTime!: string;
    yellowCards!: string;
    redCards!: string;
    saves!: string;
    tableId!: string;
}

class PlayerStats {
    localTeam!: Players;
    visitorTeam!: Players;
}

class Players {
    player!: Player[];
}

class Player {
    id!: string;
    num!: string;
    name!: string;
    pos!: string;
    posX!: string;
    posY!: string;
    shotsTotal!: string;
    shotsOnGoal!: string;
    goals!: string;
    assists!: string;
    offsides!: string;
    foulsDrawn!: string;
    foulsCommitted!: string;
    saves!: string;
    yellowCards!: string;
    redCards!: string;
    penScore!: string;
    penMiss!: string;
}

export const matchFactsFromGrpc = (mf: any) => {
    const timestamp: t.Timestamp = new t.Timestamp();
    timestamp.setSeconds(mf.matchDate.seconds);
    timestamp.setNanos(mf.matchDate.nanos);
    const date = timestamp.toDate();

    const matchFacts: MatchFacts = {
        id: mf.id,
        compId: mf.compId,
        formattedDate: mf.formattedDate,
        season: mf.season,
        week: mf.week,
        venue: mf.venue,
        venueId: mf.venueId,
        venueCity: mf.venueCity,
        venueLatitude: mf.venueLatitude,
        venueLongitude: mf.venueLongitude,
        venueCountry: mf.venueCountry,
        status: mf.status,
        timer: mf.timer,
        time: mf.time,
        localTeamId: mf.localTeamId,
        localTeamName: mf.localTeamName,
        localTeamScore: mf.localTeamScore,
        visitorTeamId: mf.visitorTeamId,
        visitorTeamName: mf.visitorTeamName,
        visitorTeamScore: mf.visitorTeamScore,
        htScore: mf.htScore,
        ftScore: mf.ftScore,
        etScore: mf.etScore,
        penaltyLocal: mf.penaltyLocal,
        penaltyVisitor: mf.penaltyVisitor,
        events: mf.events ? mf.events.map((e: Event) => eventFromGrpc(e)) : [],
        commentary: commentaryFromGrpc(mf.commentary),
        matchDate: date.toISOString().split('T')[0],
    };

    return matchFacts;
};

export const eventFromGrpc = (e: any): Event => {
    return {
        id: e.id,
        type: e.type,
        result: e.result,
        minute: e.minute,
        extraMin: e.extraMin,
        team: e.team,
        player: e.player,
        playerId: e.playerId,
        assist: e.assist,
        assistId: e.assistId,
    };
};

export const commentaryFromGrpc = (c: any): Commentary => {
    return {
        matchId: c.matchId,
        matchInfo: matchInfoFromGrpc(c.matchInfo),
        lineup: lineupFromGrpc(c.lineup),
        subs: lineupFromGrpc(c.subs),
        substitutions: substitutionsFromGrpc(c.substitutions),
        comments: c.comments ? c.comments.map((comment: Comment) => commentFromGrpc(comment)) : [],
        matchStats: matchStatsFromGrpc(c.matchStats),
        playerStats: playerStatsFromGrpc(c.playerStats),
    };
};

export const matchInfoFromGrpc = (m: any): MatchInfo => {
    return {
        stadium: m.stadium,
        attendance: m.attendance,
        referee: m.referee,
    };
};

export const lineupFromGrpc = (l: any): Lineup => {
    return {
        localTeam: l.localTeam ? l.localTeam.map((p: any) => positionFromGrpc(p)) : [],
        visitorTeam: l.visitorTeam ? l.visitorTeam.map((p: any) => positionFromGrpc(p)) : [],
    };
};

export const positionFromGrpc = (p: any): Position => {
    return {
        id: p.id,
        number: p.number,
        name: p.name,
        pos: p.pos,
    };
};

export const substitutionsFromGrpc = (s: any): Substitutions => {
    return {
        localTeam: s.localTeam ? s.localTeam.map((p: any) => substitutionFromGrpc(p)): [],
        visitorTeam: s.visitorTeam ? s.visitorTeam.map((p: any) => substitutionFromGrpc(p)): [],
    };
};

export const substitutionFromGrpc = (s: any): Substitution => {
    return {
        offName: s.offName,
        onName: s.onName,
        offId: s.offId,
        onId: s.onId,
        minute: s.minute,
        tableId: s.tableId,
    };
};

export const commentFromGrpc = (p: any): Comment => {
    return {
        id: p.id,
        important: p.important,
        goal: p.goal,
        minute: p.minute,
        comment: p.comment,
    };
};

export const matchStatsFromGrpc = (m: any): MatchStats => {
    return {
        localTeam: m.localTeam ? m.localTeam.map((p: any) => teamStatsFromGrpc(p)) : [],
        visitorTeam: m.visitorTeam ? m.visitorTeam.map((p: any) => teamStatsFromGrpc(p)) : [],
    };
};

export const teamStatsFromGrpc = (t: any): TeamStats => {
    return {
        shotsTotal: t.shotsTotal,
        shotsOnGoal: t.shotsOnGoal,
        fouls: t.fouls,
        corners: t.corners,
        offsides: t.offsides,
        possessionTime: t.possessionTime,
        yellowCards: t.yellowCards,
        redCards: t.redCards,
        saves: t.saves,
        tableId: t.tableId,
    };
};

export const playerStatsFromGrpc = (ps: any): PlayerStats => {
    return {
        localTeam: playersFromGrpc(ps.localTeam),
        visitorTeam: playersFromGrpc(ps.visitorTeam),
    };
};

export const playersFromGrpc = (ps: any): Players => {
    return {
        player: ps.player ? ps.player.map((p: any) => playerFromGrpc(p)) : [],
    };
};

export const playerFromGrpc = (p: any): Player => {
    return {
        id: p.id,
        num: p.num,
        name: p.name,
        pos: p.pos,
        posX: p.posX,
        posY: p.posY,
        shotsTotal: p.shotsTotal,
        shotsOnGoal: p.shotsOnGoal,
        goals: p.goals,
        assists: p.assists,
        offsides: p.offsides,
        foulsDrawn: p.foulsDrawn,
        foulsCommitted: p.foulsCommitted,
        saves: p.saves,
        yellowCards: p.yellowCards,
        redCards: p.redCards,
        penScore: p.penScore,
        penMiss: p.penMiss,
    };
};
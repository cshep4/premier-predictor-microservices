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
    matchDate!: string;
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
    matchInfo!: MatchInfo[];
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
    const timestamp:t.Timestamp = new t.Timestamp();
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
        events: mf.events,
        commentary: mf.commentary,
        matchDate: date.toISOString().split('T')[0],
    };

    return matchFacts;
};

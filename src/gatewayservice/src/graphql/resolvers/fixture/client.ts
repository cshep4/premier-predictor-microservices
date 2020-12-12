import * as grpc from "grpc";

export interface Empty {

}

export interface GetMatchesResponse {
    matches: [Match];
}

export interface Match {
    id: string;
    hTeam: string;
    aTeam: string;
    hGoals: number | null;
    aGoals: number | null;
    played: number;
    dateTime: string | any;
    matchday: number;
}

export interface TeamForm {
    forms: [TeamMatchResult];
}

export interface TeamMatchResult {
    result: string;
    score: string;
    opponent: string;
    location: string;
}

export function toTeamMatchResult(tmr: TeamMatchResult): TeamMatchResult {
    return {
        result: toResult(tmr.result),
        score: tmr.score,
        opponent: tmr.opponent,
        location: toLocation(tmr.location),
    }
}

function toResult(r: any): string {
    switch (r) {
        case 0:
            return "WIN";
        case 1:
            return "DRAW";
    }

    return "LOSS";
}

function toLocation(l: any): string {
    if (l == 0) {
        return "HOME";
    }

    return "AWAY";
}

export interface GetTeamFormResponse {
    teams: Map<string, TeamForm>;
}

export interface FixtureClient {
    getMatches(req: Empty, md: grpc.Metadata, callback: (err: grpc.ServiceError | Error, response: GetMatchesResponse) => void): void

    getTeamForm(req: Empty, md: grpc.Metadata, callback: (err: grpc.ServiceError | Error, response: GetTeamFormResponse) => void): void
}
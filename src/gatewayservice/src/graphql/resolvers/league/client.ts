import * as grpc from "grpc";
import {MatchFacts, matchFactsFromGrpc} from "../../model/live/matchFacts";
import {logger} from "../../log/logger";
import {TokenContext} from "../../model/context/context";

export interface GetOverviewRequest {
    id: string;
}

export interface GetOverviewResponse {
    rank: Long;
    userCount: Long;
    leagues: [LeagueSummary];
}

export interface LeagueSummary {
    leagueName: string;
    pin: Long;
    rank: Long;
}

export interface LeagueClient {
    getOverview(req: GetOverviewRequest, md: grpc.Metadata, callback: (err: grpc.ServiceError | Error, response: GetOverviewResponse) => void): void
}

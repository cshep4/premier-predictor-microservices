import {MatchFacts} from "../../model/live/matchFacts";
import * as grpc from "grpc";
import {Prediction, PredictionSummaryResponse} from "../prediction/client";
import {ClientReadableStream} from "grpc";

export class UpcomingMatchesGRPCResult {
    matches: Map<string, any> = new Map<string, any>();
}

export class UpcomingMatchesResult {
    matches: Matches[] = [];
}

class Matches {
    date!: string;
    matches: MatchFacts[] = [];
}

export interface GetMatchRequest {
    id: string;
}

export interface GetMatchResponse {
    match: MatchFacts;
}

export interface Empty {

}

export interface ListTodaysMatchesResponse {
    matches: [MatchFacts];
}

export interface MatchSummaryResponse {
    match: MatchFacts;
    predictionSummary: PredictionSummaryResponse;
    prediction: Prediction;
}

export interface LiveMatchClient {
    getLiveMatch(req: GetMatchRequest, md: grpc.Metadata, callback: (err: grpc.ServiceError | Error, response: GetMatchResponse) => void): void

    listTodaysMatches(req: Empty, md: grpc.Metadata, callback: (err: grpc.ServiceError | Error, response: ListTodaysMatchesResponse) => void): void

    getUpcomingMatches(req: any, md: grpc.Metadata): ClientReadableStream<UpcomingMatchesGRPCResult>

    getMatchSummary(req: any, md: grpc.Metadata): ClientReadableStream<MatchSummaryResponse>

    getTodaysLiveMatches(req: Empty, md: grpc.Metadata): ClientReadableStream<GetMatchResponse>
}

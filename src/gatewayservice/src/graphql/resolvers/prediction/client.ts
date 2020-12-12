import * as grpc from "grpc";
import {MatchFacts, matchFactsFromGrpc} from "../../model/live/matchFacts";
import {logger} from "../../log/logger";
import {TokenContext} from "../../model/context/context";

export interface PredictionRequest {
    userId: string;
    matchId: string;
}

export interface PredictionResponse {
    userId: string;
    matchId: string;
    hGoals: number;
    aGoals: number;
}

export interface PredictionSummaryResponse {
    homeWin: number;
    draw: number;
    awayWin: number;
}

export interface PredictionSummaryRequest {
    matchId: string;
}

export interface GetUserPredictionRequest {
    userId: string;
}

export interface GetUserPredictionsResponse {
    predictions: [Prediction];
}


export interface Prediction {
    userId: string;
    matchId: string;
    hGoals: number;
    aGoals: number;
}

export interface PredictionClient {
    getPrediction(req: PredictionRequest, md: grpc.Metadata, callback: (err: grpc.ServiceError | Error, response: PredictionResponse) => void): void

    getPredictionSummary(req: PredictionSummaryRequest, md: grpc.Metadata, callback: (err: grpc.ServiceError, response: PredictionSummaryResponse) => void): void

    getUserPredictions(req: GetUserPredictionRequest, md: grpc.Metadata, callback: (err: grpc.ServiceError, response: GetUserPredictionsResponse) => void): void
}

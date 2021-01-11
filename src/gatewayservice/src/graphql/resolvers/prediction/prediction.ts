import * as grpc from "grpc";
import {logger} from "../../log/logger";
import {TokenContext, tokenMetadata} from "../../model/context/context";
import {
    GetUserPredictionRequest, GetUserPredictionsResponse,
    PredictionClient,
    PredictionRequest,
    PredictionResponse,
    PredictionSummaryRequest,
    PredictionSummaryResponse,
    Prediction as UserPrediction
} from "./client";

export class Prediction {
    constructor(private client: PredictionClient) {
    }

    public getPredictionSummary(ctx: TokenContext, req: PredictionSummaryRequest) {
        return new Promise((resolve: any, reject: any) => {
            this.client.getPredictionSummary(req, tokenMetadata(ctx), (err: grpc.ServiceError, res: PredictionSummaryResponse) => {
                if (err) {
                    logger.error({
                        "message": "get_prediction_summary_error",
                        "error": {
                            "code": err.code,
                            "details": err.details,
                            "message": err.message,
                        },
                        "matchId": req.matchId,
                    });
                    return reject(err.message);
                }

                resolve({
                    homeWin: res.homeWin ? res.homeWin : 0,
                    draw: res.draw ? res.draw : 0,
                    awayWin: res.awayWin ? res.awayWin : 0,
                });
            });
        });
    };

    public getPrediction(ctx: TokenContext, req: PredictionRequest) {
        return new Promise((resolve: any, reject: any) => {
            if (!req.userId || !req.matchId) {
                return resolve();
            }
            this.client.getPrediction(req, tokenMetadata(ctx, req.userId), (err: grpc.ServiceError, res: PredictionResponse) => {
                if (err) {
                    if (err.code != grpc.status.NOT_FOUND) {
                        logger.error({
                            "message": "get_prediction_error",
                            "error": {
                                "code": err.code,
                                "details": err.details,
                                "message": err.message,
                            },
                            "matchId": req.matchId,
                            "userId": req.userId,
                        });
                    }
                    return resolve();
                }

                if (!res.userId || !res.matchId) {
                    return resolve();
                }

                return resolve({
                    userId: res.userId,
                    matchId: res.matchId,
                    hGoals: res.hGoals ? res.hGoals : 0,
                    aGoals: res.aGoals ? res.aGoals : 0,
                });
            });
        });
    };

    public getUserPredictions(ctx: TokenContext, req: GetUserPredictionRequest): Promise<[UserPrediction]> {
        return new Promise((resolve: any, reject: any) => {
            this.client.getUserPredictions(req, tokenMetadata(ctx, req.userId), (err: grpc.ServiceError, res: GetUserPredictionsResponse) => {
                if (err) {
                    logger.error({
                        "message": "get_user_predictions_error",
                        "error": {
                            "code": err.code,
                            "details": err.details,
                            "message": err.message,
                        },
                        "userId": req.userId,
                    });
                    return reject(err.message);
                }

                const predictions = res.predictions ? res.predictions
                    .filter(p => p.matchId && p.userId)
                    .map(p => ({
                        userId: p.userId ? p.userId : "",
                        matchId: p.matchId ? p.matchId : "",
                        hGoals: p.hGoals ? p.hGoals : 0,
                        aGoals: p.aGoals ? p.aGoals : 0,
                    })) : [];

                return resolve(predictions);
            });
        });
    };
}

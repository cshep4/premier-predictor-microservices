import * as grpc from "grpc";
import {MatchFacts, matchFactsFromGrpc} from "../../model/live/matchFacts";
import {SubscriptionContext, TokenContext, tokenMetadata} from "../../model/context/context";
import {
    GetMatchRequest,
    GetMatchResponse, ListTodaysMatchesResponse,
    LiveMatchClient, MatchSummaryResponse,
    UpcomingMatchesGRPCResult,
    UpcomingMatchesResult
} from "./client";
import {logger} from "../../log/logger";
import {secondsToDateString} from "../../util/DateUtils";
import {PubSub} from "graphql-subscriptions";
import WebSocket from "ws";
import {ClientReadableStream} from "grpc";
import {handleSubscription, withCancel} from "../../util/subsciptionUtils";


export class LiveMatch {
    constructor(private client: LiveMatchClient, private pubsub: PubSub) {
    }

    public getMatch(ctx: TokenContext, req: GetMatchRequest) {
        return new Promise((resolve: any, reject: any) => {
            this.client.getLiveMatch(req, tokenMetadata(ctx), (err: grpc.ServiceError, res: GetMatchResponse) => {
                if (err) {
                    logger.error({
                        "message": "get_live_match_error",
                        "error": {
                            "code": err.code,
                            "details": err.details,
                            "message": err.message,
                        },
                        "matchId": req.id,
                    });
                    return reject(err.message);
                }

                res.match.matchDate = secondsToDateString(res.match.matchDate.seconds);

                resolve(res.match);
            });
        });
    }

    public listTodaysMatches(ctx: TokenContext) {
        return new Promise((resolve: any, reject: any) => {
            this.client.listTodaysMatches({}, tokenMetadata(ctx), (err: grpc.ServiceError, res: ListTodaysMatchesResponse) => {
                if (err) {
                    logger.error({
                        "message": "list_todays_matches_error",
                        "error": {
                            "code": err.code,
                            "details": err.details,
                            "message": err.message,
                        },
                    });
                    return reject(err.message);
                }

                if (!res.matches) {
                    return resolve([]);
                }

                resolve(res.matches.map(m => {
                    m.matchDate = secondsToDateString(m.matchDate.seconds);
                    return m;
                }));
            });
        });
    }

    public getUpcomingMatches(_obj: any, args: any, ctx: SubscriptionContext) {
        const event = "upcoming";
        const pubsub = new PubSub();

        const call = this.client.getUpcomingMatches({}, tokenMetadata(ctx));

        handleSubscription(ctx.webSocket, call, function (fixtures: UpcomingMatchesGRPCResult) {
            fixtures.matches = new Map(Object.entries(fixtures.matches));
            const result: UpcomingMatchesResult = new UpcomingMatchesResult();
            fixtures.matches.forEach((value: any, key: any) => {
                const m: MatchFacts[] = [];
                value.matches.forEach((match: any) => m.push(matchFactsFromGrpc(match)));
                result.matches.push({
                    date: key,
                    matches: m,
                })
            });
            pubsub.publish(event, {
                upcomingMatches: result
            });
        });

        return withCancel(pubsub.asyncIterator([event]), () => {
            call.cancel();
            console.log(`Subscription closed, do your cleanup`);
        });
    };

    public getMatchSummary(_obj: any, req: any, ctx: SubscriptionContext) {
        const event = req.request.matchId;
        const pubsub = new PubSub();

        const call = this.client.getMatchSummary(req.request, tokenMetadata(ctx));

        handleSubscription(ctx.webSocket, call, function (res: MatchSummaryResponse) {
            pubsub.publish(event, {
                liveMatchSummary: {
                    liveMatch: matchFactsFromGrpc(res.match),
                    predictionSummary: {
                        homeWin: res.predictionSummary.homeWin ? res.predictionSummary.homeWin : 0,
                        draw: res.predictionSummary.draw ? res.predictionSummary.draw : 0,
                        awayWin: res.predictionSummary.awayWin ? res.predictionSummary.awayWin : 0,
                    },
                    prediction: res.prediction ? {
                        userId: res.prediction.userId,
                        matchId: res.prediction.matchId,
                        hGoals: res.prediction.hGoals ? res.prediction.hGoals : 0,
                        aGoals: res.prediction.aGoals ? res.prediction.aGoals : 0,
                    } : null,
                }
            });
        });

        return withCancel(pubsub.asyncIterator([event]), () => {
            call.cancel();
            console.log(`Subscription closed, do your cleanup`);
        });
    }

    public getTodaysLiveMatches(_obj: any, req: any, ctx: SubscriptionContext) {
        const event = "today";
        const pubsub = new PubSub();

        const call = this.client.getTodaysLiveMatches({}, tokenMetadata(ctx));

        handleSubscription(ctx.webSocket, call, function (res: GetMatchResponse) {
            pubsub.publish(event, {
                todaysLiveMatches: {
                    match: matchFactsFromGrpc(res.match),
                }
            });
        });

        return withCancel(pubsub.asyncIterator([event]), () => {
            call.cancel();
            console.log(`Subscription closed, do your cleanup`);
        });
    }
}

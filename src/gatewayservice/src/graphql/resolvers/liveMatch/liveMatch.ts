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
import {secondsToDateString} from "../../util/dateUtils";
import {PubSub} from "graphql-subscriptions";
import {handleSubscription, withCancel} from "../../util/subsciptionUtils";
import {v4 as uuid} from 'uuid';
import schedule from 'node-schedule';
import {ClientReadableStream} from "grpc";


export class LiveMatch {
    readonly todaysLiveMatchesEvent: string = "today";
    private todaysMatchesStream: ClientReadableStream<GetMatchResponse>;
    private matchStreams: Map<string, ClientReadableStream<GetMatchResponse>> = new Map<string, ClientReadableStream<GetMatchResponse>>();

    constructor(private client: LiveMatchClient, private pubsub: PubSub) {
        this.todaysMatchesStream = this.openTodaysLiveMatchesStream({token: ""});

        schedule.scheduleJob('0 1 * * *', () => {
            this.todaysMatchesStream.cancel();
            this.todaysMatchesStream = this.openTodaysLiveMatchesStream({token: ""});
        });
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
                if (value.matches) {
                    value.matches.forEach((match: any) => m.push(matchFactsFromGrpc(match)));
                }
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

    public getLiveMatchSummary(_obj: any, req: any, ctx: SubscriptionContext) {
        const event: string = uuid();

        // get live match and return
        this.client.getLiveMatch({id: req.request.matchId}, tokenMetadata(ctx), (err: grpc.ServiceError, res: GetMatchResponse) => {
            if (err) {
                logger.error({
                    "message": "get_live_match_error",
                    "error": {
                        "code": err.code,
                        "details": err.details,
                        "message": err.message,
                    },
                    "matchId": req.request.matchId,
                });
                ctx.webSocket.close(1011);
                return;
            }

            if (!res) {
                return;
            }

            this.pubsub.publish(event, {
                liveMatchSummary: {
                    liveMatch: res.match ? matchFactsFromGrpc(res.match) : {},
                    userId: req.request.userId,
                    matchId: req.request.matchId,
                }
            });
        });

        if (!this.matchStreams.has(req.request.matchId)) {
            this.matchStreams.set(req.request.matchId, this.openLiveMatchStream(ctx, req.request.matchId));
        }

        return this.pubsub.asyncIterator([event, req.request.matchId]);
    }

    private openLiveMatchStream(ctx: TokenContext, matchId: string): ClientReadableStream<MatchSummaryResponse> {
        const call = this.client.getMatchSummary({matchId: matchId}, tokenMetadata(ctx));

        call.on('data', (res: GetMatchResponse) => {
            this.pubsub.publish(matchId, {
                liveMatchSummary: {
                    liveMatch: res.match ? matchFactsFromGrpc(res.match) : {},
                    matchId: matchId,
                }
            });
        });

        call.on('end', () => {
            logger.info({
                "message": "stream_todays_matches_end",
                "matchId": matchId,
            });
        });

        call.on('error', (err: grpc.ServiceError) => {
            logger.error({
                "message": "get_match_summary_error",
                "matchId": matchId,
                "error": {
                    "code": err.code,
                    "details": err.details,
                    "message": err.message,
                },
            });

            setTimeout(() => {
                this.matchStreams.set(matchId, this.openLiveMatchStream(ctx, matchId));
            }, 1000);
        });

        call.on('status', function (status: grpc.StatusObject) {
            logger.info({
                "message": "get_match_summary_status",
                "matchId": matchId,
                "status": {
                    "code": status.code,
                    "details": status.details,
                },
            });
        });

        return call;
    }

    public getTodaysLiveMatches(_obj: any, req: any, ctx: SubscriptionContext) {
        const event: string = uuid();

        // get today's matches and return
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
                this.pubsub.publish(event, {
                    error: err.message,
                });
                ctx.webSocket.close(1011);
                return;
            }

            if (!res.matches) {
                return;
            }

            res.matches.forEach(m => {
                if (m) {
                    this.pubsub.publish(event, {
                        todaysLiveMatches: {
                            match: matchFactsFromGrpc(m),
                        }
                    });
                }
            });
        });

        return this.pubsub.asyncIterator([event, this.todaysLiveMatchesEvent]);
    }

    private openTodaysLiveMatchesStream(ctx: TokenContext): ClientReadableStream<GetMatchResponse> {
        let call = this.client.getTodaysLiveMatches({}, tokenMetadata(ctx));
        const event = this.todaysLiveMatchesEvent;

        call.on('data', (res: GetMatchResponse) => {
            this.pubsub.publish(event, {
                todaysLiveMatches: {
                    match: res.match ? matchFactsFromGrpc(res.match) : {},
                }
            });
        });

        call.on('end', () => {
            logger.info("stream_todays_matches_end");
        });

        call.on('error', (err: grpc.ServiceError) => {
            logger.error({
                "message": "stream_todays_matches_error",
                "error": {
                    "code": err.code,
                    "details": err.details,
                    "message": err.message,
                },
            });

            if (err.code != grpc.status.CANCELLED) {
                setTimeout(() => {
                    this.todaysMatchesStream = this.openTodaysLiveMatchesStream(ctx);
                }, 1000);
            }
        });

        call.on('status', function (status: grpc.StatusObject) {
            logger.info({
                "message": "stream_todays_matches_status",
                "status": {
                    "code": status.code,
                    "details": status.details,
                },
            });
        });

        return call;
    }
}

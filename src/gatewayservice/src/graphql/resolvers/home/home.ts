import {League} from "../league/league";
import {TokenContext} from "../../model/context/context";
import {GetOverviewResponse} from "../league/client";
import {Prediction} from "../prediction/prediction";
import {Prediction as UserPrediction} from "../prediction/client";
import {Fixture} from "../fixture/fixture";
import {Match} from "../fixture/client";

interface HomeFeedRequest {
    userId: string;
}

export class Home {
    constructor(private league: League,
                private prediction: Prediction,
                private fixture: Fixture) {
    }

    public feed(ctx: TokenContext, req: HomeFeedRequest) {
        return new Promise((resolve: any, reject: any) => {
            this.league.getOverview(ctx, {id: req.userId}).then((res: GetOverviewResponse) => {
                resolve({
                    id: req.userId,
                    userId: req.userId,
                    leagues: res.leagues ? res.leagues.map(l => ({
                        leagueName: l.leagueName,
                        pin: l.pin.toNumber(),
                        rank: l.rank.toNumber(),
                    })) : [],
                    rank: res.rank ? res.rank.toNumber() : -1,
                    messages: [
                        "test message"
                    ],
                })
            }, err => {
                reject(err)
            });
        });
    }

    public upcomingFixturesWithPredictions(ctx: TokenContext, req: HomeFeedRequest) {
        return new Promise((resolve: any, reject: any) => {
            this.fixture.getFixtures(ctx).then((res: [Match]) => {
                const futureFixtures = res
                    .filter(f => Date.parse(f.dateTime) > Date.now())
                    .sort((f1, f2) => Date.parse(f1.dateTime) - Date.parse(f2.dateTime));

                if (futureFixtures.length < 1) {
                    return resolve([]);
                }

                this.prediction.getUserPredictions(ctx, req).then(res => {
                    const fp = futureFixtures.slice(0, 20)
                        .map(f => {
                            const p = this.findPrediction(res, f.id);
                            if (!p) {
                                return f;
                            }

                            return {
                                id: f.id,
                                hTeam: f.hTeam,
                                aTeam: f.aTeam,
                                hGoals: f.hGoals,
                                aGoals: f.aGoals,
                                hPrediction: p.hGoals,
                                aPrediction: p.aGoals,
                                played: f.played,
                                dateTime: f.dateTime,
                                matchday: f.matchday,
                            }
                        });

                    return resolve(fp);
                }, err => {
                    return reject(err);
                });
            }, err => {
                return reject(err);
            });
        });
    }

    private findPrediction(res: [UserPrediction], id: string) {
        for (let i = 0; i < res.length; i++) {
            if (id === res[i].matchId) {
                return res[i];
            }
        }

        return;
    }
}
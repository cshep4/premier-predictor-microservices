import * as grpc from "grpc";
import {logger} from "../../log/logger";
import {TokenContext, tokenMetadata} from "../../model/context/context";
import {
    FixtureClient, GetMatchesResponse, GetTeamFormResponse, Match, TeamMatchResult, toTeamMatchResult,
} from "./client";
import {secondsToDateTimeString} from "../../util/dateUtils";

export class Fixture {
    constructor(private client: FixtureClient) {
    }

    public getFixtures(ctx: TokenContext): Promise<[Match]> {
        return new Promise((resolve: any, reject: any) => {
            this.client.getMatches({}, tokenMetadata(ctx), (err: grpc.ServiceError, res: GetMatchesResponse) => {
                if (err) {
                    logger.error({
                        "message": "get_matches_error",
                        "error": {
                            "code": err.code,
                            "details": err.details,
                            "message": err.message,
                        },
                    });
                    return reject(err.message);
                }

                let fixtures: [Match?] = [];
                for (let i = 0; i < res.matches.length; i++) {
                    const m: Match = {
                        id: res.matches[i].id,
                        hTeam: res.matches[i].hTeam,
                        aTeam: res.matches[i].aTeam,
                        hGoals: null,
                        aGoals: null,
                        played: res.matches[i].played,
                        dateTime: secondsToDateTimeString(res.matches[i].dateTime.seconds),
                        matchday: res.matches[i].matchday,
                    };
                    if (res.matches[i].played != 0) {
                        m.hGoals = res.matches[i].hGoals ? res.matches[i].hGoals : 0;
                        m.aGoals = res.matches[i].aGoals ? res.matches[i].aGoals : 0;
                    }

                    fixtures.push(m);
                }

                resolve(fixtures);
            });
        });
    };

    public getTeamForm(ctx: TokenContext) {
        return new Promise((resolve: any, reject: any) => {
            this.client.getTeamForm({}, tokenMetadata(ctx), (err: grpc.ServiceError, res: GetTeamFormResponse) => {
                if (err) {
                    logger.error({
                        "message": "get_team_form_error",
                        "error": {
                            "code": err.code,
                            "details": err.details,
                            "message": err.message,
                        },
                    });
                    return resolve(err.message);
                }

                const map = new Map(Object.entries(res.teams));

                const forms = [];
                for (let [key, value] of map) {
                    forms.push({
                        team: key,
                        forms: value.forms ? value.forms.map((tmr: TeamMatchResult) => toTeamMatchResult(tmr)) : [],
                    });
                }

                resolve(forms);
            });
        });
    };
}

function* entries(obj: any) {
    for (let key in obj)
        yield [key, obj[key]];
}


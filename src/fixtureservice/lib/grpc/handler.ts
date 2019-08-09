import {Service} from "../service/service";
import {Fixture, fixtureToGrpc} from "../model/fixture";
import {Forms} from "../model/form";
import {logger} from "../utils/utils";

export class Handler {
    constructor(private service: Service) {
    }

    public getMatches(call, callback) {
        this.service.getAllFixtures().then((fixtures: Fixture[]) => {
            const matches = fixtures.map(f => fixtureToGrpc(f));
            this.log("getAllFixtures", true, matches);
            callback(null, {matches: matches});
        }, err => {
            this.log("getAllFixtures", false, err);
            callback(err);
        });
    }

    public getTeamForms(call, callback) {
        this.service.getTeamForms().then((forms: Forms) => {
            this.log("getTeamForms", true, forms);
            callback(null, forms);
        }, err => {
            this.log("getTeamForms", false, err);
            callback(err);
        });
    }

    public getFutureFixtures(call, callback) {
        this.service.getFutureFixtures().then((fixtures: Map<string, string>) => {
            this.log("getFutureFixtures", true, fixtures);
            callback(null, {matches: fixtures});
        }, err => {
            this.log("getFutureFixtures", false, err);
            callback(err);
        });
    }

    private log(method: string, successful: boolean, response: any = null) {
        if (successful) {
            logger.info({
                "req": {
                    "protocol": "grpc",
                    "method": method
                },
                "res": {
                    "data": response
                }
            });
            return;
        }

        logger.error({
            "req": {
                "protocol": "grpc",
                "method": method
            },
            "res": {
                "error": response
            }
        });
    }
}
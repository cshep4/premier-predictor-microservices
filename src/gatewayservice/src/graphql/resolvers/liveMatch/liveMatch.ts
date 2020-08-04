import * as grpc from "grpc";
import {MatchFacts, matchFactsFromGrpc} from "../../model/live/matchFacts";

interface Params {
    token: string;
}

class UpcomingMatchesGRPCResult {
    matches: Map<string, any> = new Map<string, any>();
}

class UpcomingMatchesResult {
    matches: Matches[] = [];
}

class Matches {
    date!: string;
    matches: MatchFacts[] = [];
}

export class LiveMatch {
    constructor(private client: any, private pubsub: any) {
    }

    public getUpcomingMatches(_obj: any, args: any, context: any) {
        const event = "upcoming";
        this.getUpcoming(context, event);
        return this.pubsub.asyncIterator([event]);
    };

    private getUpcoming(context: any, event: string) {
        const pubsub = this.pubsub;
        const metadata = new grpc.Metadata();
        metadata.add('token', context.token);
        const call = this.client.getUpcomingMatches({}, metadata);
        call.on('data', function (fixtures: UpcomingMatchesGRPCResult) {
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
        call.on('end', function () {
            // The server has finished sending
            console.log("end");
        });
        call.on('error', function (err: any) {
            // An error has occurred and the stream has been closed.
            console.log("error: " + err);
        });
        call.on('status', function (status: any) {
            // process status
            console.log("status: " + status);
        });
    }
}

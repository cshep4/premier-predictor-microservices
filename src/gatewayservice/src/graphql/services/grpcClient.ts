import * as grpc from 'grpc';
import * as protoLoader from '@grpc/proto-loader';
import {Listener, Metadata} from 'grpc';

const interceptorAuth: any = (options: any, nextCall: any) =>
    new grpc.InterceptingCall(nextCall(options), {
        start: function (metadata: Metadata, listener: Listener, next: any) {
            console.log(metadata);
            console.log(options);
            metadata.add('x-api-key', 'myapikey');
            next(metadata, listener);
        }
    });

const opts = {
    "grpc.keepalive_time_ms": 6000,
    "grpc.keepalive_timeout_ms": 2000,
    "grpc.keepalive_permit_without_calls": 1,
    // interceptors: [interceptorAuth]
};

function getProto(fileName: string): any {
    const packageDefinition: any = protoLoader.loadSync(__dirname + '../../../../../../proto-gen/model/proto/' + fileName + '.proto');
    return grpc.loadPackageDefinition(packageDefinition).model;
}

const auth = getProto('auth');
const liveMatch = getProto('live');
const prediction = getProto('prediction');
const fixture = getProto('fixture');
const league = getProto('league');
const user = getProto('user');

const authAddr = process.env.AUTH_ADDR ? process.env.AUTH_ADDR : 'localhost:3001';
const fixtureAddr = process.env.FIXTURE_ADDR ? process.env.FIXTURE_ADDR : 'localhost:3006';
const predictionAddr = process.env.PREDICTION_ADDR ? process.env.PREDICTION_ADDR : 'localhost:3007';
const liveMatchAddr = process.env.LIVE_MATCH_ADDR ? process.env.LIVE_MATCH_ADDR : 'localhost:3008';
const leagueAddr = process.env.LEAGUE_ADDR ? process.env.LEAGUE_ADDR : 'localhost:3009';
const userAddr = process.env.USER_ADDR ? process.env.USER_ADDR : 'localhost:3005';

export default {
    auth: () => new auth.AuthService(authAddr, grpc.credentials.createInsecure(), opts),
    fixture: () => new fixture.FixtureService(fixtureAddr, grpc.credentials.createInsecure(), opts),
    prediction: () => new prediction.PredictionService(predictionAddr, grpc.credentials.createInsecure(), opts),
    liveMatch: () => new liveMatch.LiveMatchService(liveMatchAddr, grpc.credentials.createInsecure(), opts),
    league: () => new league.LeagueService(leagueAddr, grpc.credentials.createInsecure(), opts),
    user: () => new user.UserService(userAddr, grpc.credentials.createInsecure(), opts),
};

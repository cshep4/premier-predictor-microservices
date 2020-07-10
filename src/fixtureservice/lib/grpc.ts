import {Handler} from "./grpc/handler";
import {logger} from "./utils/utils";

const PORT = process.env.PORT;

export class Grpc {
    constructor(private handler: Handler) {
    }

    public start(grpc: any) {
        const PROTO_PATH = __dirname + '/../protodefs/fixture.proto';
        const protoLoader = require('@grpc/proto-loader');
        const packageDefinition = protoLoader.loadSync(
            PROTO_PATH,
            {
                keepCase: true,
                longs: String,
                enums: String,
                defaults: true,
                oneofs: true
            });
        const fixtureProto = grpc.loadPackageDefinition(packageDefinition).model;

        const opts = {
            "grpc.keepalive_time_ms": 60000,
            "grpc.keepalive_permit_without_calls": 1
        };
        const server = new grpc.Server(opts);

        server.addService(fixtureProto.FixtureService.service, {
            getTeamForm: (call, callback) => this.handler.getTeamForms(call, callback),
            getMatches: (call, callback) => this.handler.getMatches(call, callback),
            getFutureFixtures: (call, callback) => this.handler.getFutureFixtures(call, callback),
        });

        server.bind(`0.0.0.0:${PORT}`, grpc.ServerCredentials.createInsecure());
        server.start();
        logger.info('GRPC server listening on port ' + PORT);
    }
}

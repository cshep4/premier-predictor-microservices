import {NextFunction, Request, response, Response} from "express";

export class Middleware {
    private client;

    constructor(grpc: any) {
        const PROTO_PATH = __dirname + '/../../protodefs/auth.proto';
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
        const authProto = grpc.loadPackageDefinition(packageDefinition).model;

        const authAddr = process.env.AUTH_ADDR;

        const opts = {
            "grpc.keepalive_time_ms": 60000,
            "grpc.keepalive_timeout_ms": 20000,
            "grpc.keepalive_permit_without_calls" : 1
        };
        this.client = new authProto.AuthService(authAddr, grpc.credentials.createInsecure(), opts);
    }

    public validateHttp(req: Request, res: Response, next: NextFunction) {
        const token = req.header("Authorization");

        this.client.validate({token: token}, (err, response) => {
            if (err) {
                res.status(401)
                    .send(err);
                return;
            }

            next();
        });
    }

    public validateGrpc(req: Request, res: Response, next: NextFunction) {
        const token = req.header("Authorization");

        this.client.validate({token: token}, (err, response) => {
            if (err) {
                res.status(401)
                    .send(err);
                return;
            }

            next();
        });
    }
}
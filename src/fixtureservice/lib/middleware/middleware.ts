import {NextFunction, Request, Response} from "express";

export class Middleware {
    private client;

    constructor(grpc: any) {
        const protoPath = process.env.PROTO_PATH ? __dirname + '/..' + process.env.PROTO_PATH : __dirname + '/../../../../proto-gen/model/proto/';
        const protoFile = protoPath + 'auth.proto';
        const protoLoader = require('@grpc/proto-loader');
        const packageDefinition = protoLoader.loadSync(
            protoFile,
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
            "grpc.keepalive_permit_without_calls": 1
        };
        this.client = new authProto.AuthService(authAddr, grpc.credentials.createInsecure(), opts);
    }

    public validateHttp(req: Request, res: Response, next: NextFunction) {
        // const span = this.tracer.startChildSpan({ name: req.url });

        const token = req.header("Authorization");

        const validateReq = {
            token: token,
            role: 1,
        };
        this.client.validate(validateReq, (err, response) => {
            if (err) {
                res.status(401)
                    .send(err);
                // span.end();
                return;
            }

            next();
            // span.end();
        });
    }

    public validateGrpc(req: Request, res: Response, next: NextFunction) {
        const token = req.header("Authorization");

        const validateReq = {
            token: token,
            role: 1,
        };
        this.client.validate(validateReq, (err, response) => {
            if (err) {
                res.status(401)
                    .send(err);
                return;
            }

            next();
        });
    }
}

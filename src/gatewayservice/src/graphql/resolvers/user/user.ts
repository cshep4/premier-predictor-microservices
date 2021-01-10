import * as grpc from "grpc";
import {logger} from "../../log/logger";
import {TokenContext, tokenMetadata} from "../../model/context/context";
import {
    GetUserRequest, GetUserResponse, UserClient,
} from "./client";

export class User {
    constructor(private client: UserClient) {
    }

    public getUserScore(ctx: TokenContext, req: GetUserRequest): Promise<number> {
        return new Promise((resolve: any, reject: any) => {
            this.client.getUser(req, tokenMetadata(ctx, req.id), (err: grpc.ServiceError, res: GetUserResponse) => {
                if (err) {
                    logger.error({
                        "message": "get_user_error",
                        "error": {
                            "code": err.code,
                            "details": err.details,
                            "message": err.message,
                        },
                        "userId": req.id,
                    });
                    return reject(err.message);
                }

                resolve(res.user.score ? res.user.score : 0);
            });
        });
    };
}

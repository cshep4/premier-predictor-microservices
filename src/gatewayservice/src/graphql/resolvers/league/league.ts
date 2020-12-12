import * as grpc from "grpc";
import {logger} from "../../log/logger";
import {TokenContext, tokenMetadata} from "../../model/context/context";
import {
    GetOverviewRequest, GetOverviewResponse, LeagueClient,
} from "./client";

export class League {
    constructor(private client: LeagueClient) {
    }

    public getOverview(ctx: TokenContext, req: GetOverviewRequest): Promise<GetOverviewResponse> {
        return new Promise((resolve: any, reject: any) => {
            this.client.getOverview(req, tokenMetadata(ctx, req.id), (err: grpc.ServiceError, res: GetOverviewResponse) => {
                if (err) {
                    logger.error({
                        "message": "get_overview_error",
                        "error": {
                            "code": err.code,
                            "details": err.details,
                            "message": err.message,
                        },
                        "userId": req.id,
                    });
                    return reject(err.message);
                }

                resolve(res);
            });
        });
    };
}

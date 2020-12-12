import * as grpc from "grpc";
import WebSocket from "ws";

export interface TokenContext {
    token: string;
}

export interface SubscriptionContext {
    token: string;
    webSocket: WebSocket;
}

export function tokenMetadata(ctx: TokenContext | SubscriptionContext, audience: string = ""): grpc.Metadata {
    const md = new grpc.Metadata();
    md.add('token', ctx.token);
    md.add('audience', audience);

    return md
}
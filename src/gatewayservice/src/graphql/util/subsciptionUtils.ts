import WebSocket from "ws";
import {ClientReadableStream} from "grpc";
import * as grpc from "grpc";

export function withCancel<T>(asyncIterator: AsyncIterator<T | undefined>, onCancel: Function): AsyncIterator<T | undefined> {
    return {
        ...asyncIterator,
        return() {
            onCancel();
            return asyncIterator.return ? asyncIterator.return() : Promise.resolve({value: undefined, done: true})
        }
    };
}

export function handleSubscription(webSocket: WebSocket, call: ClientReadableStream<any>, onCall: (res: any) => void) {
    webSocket.addListener('close', (code: number, message: string) => {
        call.cancel();
        console.log("websocket close");
    });

    webSocket.addListener('error', (err: Error) => {
        call.cancel();
        console.log("websocket error: " + err);
    });

    call.on('data', onCall);

    call.on('end', function () {
        console.log("stream end");
    });

    call.on('error', function (err: any) {
        console.log("stream error: " + err);
    });

    call.on('status', function (status: grpc.StatusObject) {
        if (status.code != grpc.status.CANCELLED) {
            webSocket.close();
        }
        console.log("stream status - code: " + status.code + ", details: " + status.details);
    });
}
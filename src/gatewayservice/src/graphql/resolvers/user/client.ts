import * as grpc from "grpc";

export interface GetUserRequest {
    id: string;
}

export interface GetUserResponse {
    user: User;
}

interface User {
    id: string;
    firstName: string;
    surname: string;
    predictedWinner: string;
    score: number;
    email: string;
    password: string;
    signature: string;
}

export interface UserClient {
    getUser(req: GetUserRequest, md: grpc.Metadata, callback: (err: grpc.ServiceError | Error, response: GetUserResponse) => void): void
}

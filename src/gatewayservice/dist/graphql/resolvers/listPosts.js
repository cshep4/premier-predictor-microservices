"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const grpcClient_1 = __importDefault(require("../services/grpcClient"));
const client = grpcClient_1.default.auth();
exports.default = (root, params) => {
    return new Promise((resolve, reject) => {
        client.validate(params, function (err, response) {
            if (err) {
                return reject(err);
            }
            resolve(response);
        });
    });
};

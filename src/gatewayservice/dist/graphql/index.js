"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const apollo_server_1 = require("apollo-server");
const resolvers_1 = __importDefault(require("./resolvers"));
const schema_1 = __importDefault(require("./api/schema"));
exports.schema = apollo_server_1.makeExecutableSchema({
    resolvers: resolvers_1.default,
    typeDefs: schema_1.default
});

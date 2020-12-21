"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const graphql_1 = require("./graphql");
const express_1 = __importDefault(require("express"));
const apollo_server_express_1 = require("apollo-server-express");
const http = __importStar(require("http"));
const server = new apollo_server_express_1.ApolloServer({
    schema: graphql_1.schema,
    subscriptions: {
        onConnect: (connectionParams, webSocket) => {
            if (connectionParams.authorization) {
                return {
                    connectionParams: connectionParams,
                    webSocket: webSocket,
                    token: connectionParams.authorization,
                };
            }
            return {
                connectionParams: connectionParams,
                webSocket: webSocket,
                token: connectionParams.token,
            };
        },
        onDisconnect: (webSocket, context) => {
            console.log("onDisconnect");
        },
    },
    context: ({ req, connection }) => __awaiter(this, void 0, void 0, function* () {
        if (connection) {
            return connection.context;
        }
        else {
            const token = req.headers.authorization || "";
            return { token };
        }
    }),
});
const app = express_1.default();
server.applyMiddleware({ app });
const httpServer = http.createServer(app);
server.installSubscriptionHandlers(httpServer);
httpServer.listen({ port: 4000 }, () => {
    console.log(`🚀 Server ready at http://localhost:4000${server.graphqlPath}`);
    console.log(`🚀 Subscriptions ready at ws://localhost:4000${server.subscriptionsPath}`);
});

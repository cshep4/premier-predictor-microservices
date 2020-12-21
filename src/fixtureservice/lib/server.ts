import {Http} from "./http";
import {Router} from "./http/router";
import {Repository} from "./repository/repository";
import {Service} from "./service/service";
import {Controller} from "./http/controller";
import {Middleware} from "./middleware/middleware";
import {Grpc} from "./grpc";
import {Handler} from "./grpc/handler";
import {FormFormatter} from "./component/form-formatter";
import {FixtureFormatter} from "./component/fixture-formatter";

export const grpc = require('grpc');

const repository = new Repository();
const fixtureFormatter = new FixtureFormatter();
const formFormatter = new FormFormatter();
const service = new Service(repository, fixtureFormatter, formFormatter);
const middleware = new Middleware(grpc);

const controller = new Controller(service);
const router = new Router(controller, middleware);
const httpServer = new Http(router);

const handler = new Handler(service);
const grpcServer = new Grpc(handler);

httpServer.start();
grpcServer.start(grpc);

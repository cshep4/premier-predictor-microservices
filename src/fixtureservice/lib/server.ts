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
import * as path from "path";

export const grpc = require('grpc');

require('@google-cloud/trace-agent').start({
    projectId: "prempred"
});
require('@google-cloud/profiler').start({
    serviceContext: {
        service: 'fixtureservice',
        version: '1.0.0'
    }
});

const {plugin} = require('@opencensus/instrumentation-grpc');
const tracing = require('@opencensus/nodejs');
const {StackdriverTraceExporter} = require('@opencensus/exporter-stackdriver');

const exporter = new StackdriverTraceExporter({projectId: "prempred"});
tracing.registerExporter(exporter).start();
const tracer = tracing.start({
    samplingRate: 1,
    plugins: {
        http: {
            module: "@opencensus/instrumentation-http",
            config: {
                ignoreIncomingPaths: [
                    /^\/health/,
                ],
            },
        },
    },
}).tracer;

const basedir = path.dirname(require.resolve('grpc'));
plugin.enable(grpc, tracer, "1.0.0", {}, basedir);

const repository = new Repository();
const fixtureFormatter = new FixtureFormatter();
const formFormatter = new FormFormatter();
const service = new Service(repository, fixtureFormatter, formFormatter);
const middleware = new Middleware(grpc, tracer);

const controller = new Controller(service);
const router = new Router(controller, middleware);
const httpServer = new Http(router);

const handler = new Handler(service, tracer);
const grpcServer = new Grpc(handler);

httpServer.start();
grpcServer.start(grpc);

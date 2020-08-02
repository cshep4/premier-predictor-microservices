import fs from 'fs';
import * as grpc from 'grpc';
import * as protoLoader from '@grpc/proto-loader';
import {Listener, Metadata} from 'grpc';

const interceptorAuth:any = (options:any, nextCall:any) =>
  new grpc.InterceptingCall(nextCall(options), {
    start: function(metadata:Metadata, listener:Listener, next:any) {
        console.log(metadata);
        console.log(options);
      metadata.add('x-api-key', 'myapikey');
      next(metadata, listener);
    }
  });

const opts = {
    "grpc.keepalive_time_ms": 60000,
    "grpc.keepalive_timeout_ms": 20000,
    "grpc.keepalive_permit_without_calls": 1,
    // interceptors: [interceptorAuth]
};

function getProto(fileName: string): any {
    const packageDefinition: any = protoLoader.loadSync(__dirname + '../../../../../../proto-gen/model/protodefs/' + fileName + '.proto');
    return  grpc.loadPackageDefinition(packageDefinition).model;
}

const auth = getProto('auth');
const liveMatch = getProto('live');

export default {
    auth: () => new auth.AuthService('localhost:3001', grpc.credentials.createInsecure(), opts),
    liveMatch: () => new liveMatch.LiveMatchService('localhost:3008', grpc.credentials.createInsecure(), opts),
};

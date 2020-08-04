import grpc from '../services/grpcClient';

const client = grpc.auth();

export interface Params {
  token: string;
}

export default (root:any, params: Params) => {
  return new Promise((resolve: any, reject: any) => {
    client.validate(params, function(err: any, response: any) {
      if (err) {
        return reject(err);
      }
      resolve(response);
    });
  });
};

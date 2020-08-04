import grpc from '../services/grpcClient';

const client = grpc.auth();

export default (root:any, params: any) => {
  return new Promise((resolve: any, reject: any) => {
    client.validate(params.data, function(err: any, response: any) {
      if (err) {
        return reject(err.details);
      }
      resolve({message: "post.created", result: response});
    });
  });
};

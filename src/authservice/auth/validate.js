const grpc = require('grpc');

const jwt = require('jsonwebtoken');

const logger = require('../util/util').logger;
const error = require('../util/util').error;

module.exports.validate = (call, callback) => {
    logger.info('received authorise request');
    const request = call.request;
    const token = request.token;

    try {
        const user = jwt.verify(token, process.env.JWT_SECRET);

        logger.info(`user validated: ${JSON.stringify(user)}`);

        callback(null);
    } catch (e) {
        logger.info(e.message);
        callback(error(e.message, grpc.status.UNAUTHENTICATED));
    }
};

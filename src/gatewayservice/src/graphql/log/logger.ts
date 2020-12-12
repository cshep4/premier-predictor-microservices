import pino from "pino";

export const logger = pino({
    name: 'gatewayservice',
    messageKey: 'message',
    changeLevelName: 'severity',
    useLevelLabels: true
});
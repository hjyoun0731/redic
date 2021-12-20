#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "redic.h"

redisContext* NewConnect(char *network, int port) {
    redisContext *c;

    const char *hostname = (sizeof(network) > 0) ? network : "127.0.0.1";
    port = (port > 0) ? port : 6379;
    
    struct timeval timeout = { 1, 500000 }; // 1.5 seconds

    c = redisConnectWithTimeout(hostname, port, timeout);

    if (c != NULL && !c->err) {
        return c;
    }
    return NULL;
}

int CloseConnect(redisContext *c) {
    redisFree(c);
    return 0;
}

const char *RedisCommand(redisContext *c, const char *format) {
    redisReply *reply;
    reply = redisCommand(c,format);
    
    return (char *)reply->str;
}

#ifndef __REDIC_H
# define __REDIC_H

# include "hiredis.h"

redisContext* NewConnect(char *network, int port);
int CloseConnect(redisContext *c);
const char *RedisCommand(redisContext *c, const char *format);

#endif
#include <stdint.h>
#include <stddef.h>

typedef struct account account_t;

extern account_t * from_seed(const uint8_t *n, size_t len);
char * account_private_key(const account_t *);
char * account_view_key(const account_t *);
char * account_address(const account_t *);
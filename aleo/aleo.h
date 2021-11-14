#include <stdint.h>
#include <stddef.h>

typedef struct account account_t;

int last_error_length();
int last_error_message(char *buffer, int length);

extern account_t * from_sk(const char *sk);
extern account_t * from_seed(const uint8_t *n, size_t len, const char *network);
char * account_private_key(const account_t *);
char * account_view_key(const account_t *);
char * account_address(const account_t *);
#include <stdint.h>
#include <stddef.h>

struct Buffer {
  uint8_t *data;
  uintptr_t len;
};

typedef struct Buffer buffer_t;

/* c_error */
int last_error_length();
int last_error_message(char *buffer, int length);

/* account */
typedef struct account account_t;
extern account_t * from_sk(const char *sk);
extern account_t * from_seed(const uint8_t *n, size_t len);
char * account_private_key(const account_t *);
char * account_view_key(const account_t *);
char * account_address(const account_t *);

/* record */
typedef struct record record_t;
extern record_t * new_input_record(const char *addr, uint64_t value, const uint8_t *payload, const uint8_t *randomness, size_t randomness_len);
extern record_t *from_record(const char *addr,
                              uint64_t val,
                              const uint8_t *payload,
                              const char *serial_number_nonce,
                              const char *commitment_randomness);
char *record_owner(const record_t *);
uint64_t record_value(const record_t *);
buffer_t record_payload(const record_t *);
char *record_serial_number_nonce(const record_t *);
char *record_commitment_randomness(const record_t *);
char *record_commitment(const record_t *);
char *record_program_id(const record_t *);
char *encrypt_record(const record_t *, const uint8_t *randomness, size_t randomness_len);
char *decrypt_record(const char *ciphertext, const char *view_key);

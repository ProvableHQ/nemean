use rand::{rngs::StdRng, SeedableRng};
use snarkvm_dpc::{
    network::testnet2::Testnet2, Account, AccountScheme, Address, AleoAmount, Network, Transaction,
};
use snarkvm_utilities::{
    has_duplicates,
    io::{Read, Result as IoResult, Write},
    FromBytes, FromBytesDeserializer, ToBytes, ToBytesSerializer,
};
use std::ffi::CString;
use std::{slice, str::FromStr};

#[no_mangle]
pub extern "C" fn new_coinbase_transaction(
    addr: *const libc::c_char,
    val: u64,
    randomness: *const u8,
    randomness_len: libc::size_t,
) -> *mut libc::c_char {
    // address
    let c_addr = unsafe {
        assert!(!addr.is_null());

        std::ffi::CStr::from_ptr(addr)
    };
    let address = Address::<Testnet2>::from_str(c_addr.to_str().unwrap()).unwrap();

    let c_rng = unsafe {
        assert!(!randomness.is_null());
        slice::from_raw_parts(randomness, randomness_len as usize)
    };

    let mut rng: StdRng = SeedableRng::from_seed(c_rng.try_into().unwrap());

    let transaction =
        Transaction::<Testnet2>::new_coinbase(address, AleoAmount(val as i64), &mut rng).unwrap();
    let tx_bytes = transaction.to_bytes_le().unwrap();
    let serialized_tx = hex::encode(&tx_bytes);
    CString::new(serialized_tx.to_string()).unwrap().into_raw()
}

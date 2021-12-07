use crate::c_error;
use rand::{rngs::StdRng, SeedableRng};
use rand::{thread_rng, Rng};
use rand_chacha::ChaChaRng;
use snarkvm_dpc::{
    network::testnet2::Testnet2, Address, AleoAmount, LedgerProof, PrivateKey, Record, Request,
    Transaction, VirtualMachine,
};
use snarkvm_utilities::{FromBytes, ToBytes};
use std::ffi::{CStr, CString};
use std::{slice, str::FromStr};

#[no_mangle]
pub extern "C" fn new_coinbase_transaction(
    addr: *const libc::c_char,
    val: i64,
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
        Transaction::<Testnet2>::new_coinbase(address, AleoAmount(val as i64), true, &mut rng)
            .unwrap();
    let tx_bytes = transaction.to_bytes_le().unwrap();
    let serialized_tx = hex::encode(&tx_bytes);
    CString::new(serialized_tx.to_string()).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn new_transfer_transaction(
    in_record: *mut Record<Testnet2>,
    ledger_proof_one: *const libc::c_char,
    ledger_proof_two: *const libc::c_char,
    private_key: *const libc::c_char,
    amount: i64,
    fee: i64,
    address: *const libc::c_char,
) -> *mut libc::c_char {
    let seed: u64 = thread_rng().gen();
    let rng = &mut ChaChaRng::seed_from_u64(seed);

    let c_in_record = unsafe {
        assert!(!in_record.is_null());
        &*in_record.clone()
    };

    let c_private_key = unsafe {
        assert!(!private_key.is_null());

        CStr::from_ptr(private_key)
    };
    let sk = PrivateKey::<Testnet2>::from_str(c_private_key.to_str().unwrap()).unwrap();

    let c_address = unsafe {
        assert!(!address.is_null());

        CStr::from_ptr(address)
    };

    let addr = Address::<Testnet2>::from_str(c_address.to_str().unwrap()).unwrap();

    let c_ledger_proof_one = unsafe {
        assert!(!ledger_proof_one.is_null());

        CStr::from_ptr(ledger_proof_one)
    };

    let c_ledger_proof_two = unsafe {
        assert!(!ledger_proof_two.is_null());

        CStr::from_ptr(ledger_proof_two)
    };

    let record_bytes1 = hex::decode(c_ledger_proof_one.to_str().unwrap()).unwrap();
    let record_proof1 = LedgerProof::<Testnet2>::from_bytes_le(&record_bytes1).unwrap();
    let record_bytes2 = hex::decode(c_ledger_proof_two.to_str().unwrap()).unwrap();
    let record_proof2 = LedgerProof::<Testnet2>::from_bytes_le(&record_bytes2).unwrap();
    let ledger_root = record_proof1.ledger_root();

    let state = match Request::<Testnet2>::new_transfer(
        &sk,
        vec![c_in_record.clone()],
        vec![record_proof1, record_proof2],
        addr,
        AleoAmount(amount),
        AleoAmount(fee),
        true,
        rng,
    ) {
        Ok(res) => res,
        Err(_) => {
            c_error::update_last_error(snarkvm_utilities::error(
                "could not create transfer request",
            ));
            return std::ptr::null_mut();
        }
    };

    let vm = VirtualMachine::<Testnet2>::new(ledger_root).unwrap();

    let res = match vm.execute(&state, rng) {
        Ok(res) => res,
        Err(_) => {
            c_error::update_last_error(snarkvm_utilities::error("could not execute transaction"));
            return std::ptr::null_mut();
        }
    };
    let transaction = match res.0.finalize() {
        Ok(res) => res,
        Err(_) => {
            c_error::update_last_error(snarkvm_utilities::error("could not finalize transaction"));
            return std::ptr::null_mut();
        }
    };

    let tx_bytes = transaction.to_bytes_le().unwrap();
    let serialized_tx = hex::encode(&tx_bytes);

    CString::new(serialized_tx.to_string()).unwrap().into_raw()
}

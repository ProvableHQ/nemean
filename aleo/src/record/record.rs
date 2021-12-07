// create a output record
// create an input record
// encrypt a record
// decrypt a record

use crate::c_error;
use rand::{rngs::StdRng, SeedableRng};
use rand::{thread_rng, Rng};
use rand_chacha::ChaChaRng;
use snarkvm_algorithms::EncryptionScheme;
use snarkvm_dpc::{
    network::testnet2::Testnet2, Address, AleoAmount, Network, Payload, Record, ViewKey,
};
use snarkvm_utilities::{FromBytes, ToBytes};
use std::ffi::{CStr, CString};
use std::{slice, str::FromStr};

#[no_mangle]
pub extern "C" fn new_input_record(
    addr: *const libc::c_char,
    val: i64,
    payload: *const u8,
    randomness: *const u8,
    randomness_len: libc::size_t,
) -> *mut Record<Testnet2> {
    // Convert addr into Address
    let c_addr = unsafe {
        assert!(!addr.is_null());

        CStr::from_ptr(addr).to_str().unwrap()
    };

    let address = match Address::<Testnet2>::from_str(c_addr) {
        Ok(address) => address,
        Err(error) => {
            c_error::update_last_error(error);
            return std::ptr::null_mut();
        }
    };

    let c_rng = unsafe {
        assert!(!randomness.is_null());
        slice::from_raw_parts(randomness, randomness_len as usize)
            .try_into()
            .unwrap()
    };

    let mut rng: StdRng = SeedableRng::from_seed(c_rng);

    let c_payload = unsafe {
        assert!(!payload.is_null());

        slice::from_raw_parts(payload, Testnet2::RECORD_PAYLOAD_SIZE_IN_BYTES)
    };

    let record_payload = match Payload::from_bytes_le(&c_payload) {
        Ok(payload) => payload,
        _ => {
            c_error::update_last_error(snarkvm_utilities::error("cannot read from payload"));
            return std::ptr::null_mut();
        }
    };

    let record = match Record::new(
        address,
        AleoAmount::from_aleo(val),
        record_payload,
        *Testnet2::noop_program_id(),
        &mut rng,
    ) {
        Ok(record) => record,
        Err(error) => {
            c_error::update_last_error(error);
            return std::ptr::null_mut();
        }
    };

    Box::into_raw(Box::new(record))
}

#[no_mangle]
pub extern "C" fn from_record(
    addr: *const libc::c_char,
    val: i64,
    payload: *const u8,
) -> *mut Record<Testnet2> {
    let c_addr = unsafe {
        assert!(!addr.is_null());

        CStr::from_ptr(addr)
    };

    let address = Address::<Testnet2>::from_str(c_addr.to_str().unwrap()).unwrap();

    let c_payload = unsafe {
        assert!(!payload.is_null());

        slice::from_raw_parts(payload, Testnet2::RECORD_PAYLOAD_SIZE_IN_BYTES)
    };

    let seed: u64 = thread_rng().gen();
    let rng = &mut ChaChaRng::seed_from_u64(seed);

    let (_randomness, randomizer, record_view_key) =
        Testnet2::account_encryption_scheme().generate_asymmetric_key(&*address, rng);

    let record = match Record::from(
        address,
        AleoAmount(val),
        Payload::from_bytes_le(&c_payload).unwrap(),
        *Testnet2::noop_program_id(),
        randomizer.into(),
        record_view_key.into(),
    ) {
        Ok(record) => record,
        Err(error) => {
            c_error::update_last_error(error);
            return std::ptr::null_mut();
        }
    };

    Box::into_raw(Box::new(record))
}

#[no_mangle]
pub extern "C" fn encrypt_record(ptr: *mut Record<Testnet2>) -> *mut libc::c_char {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let record_ciphertext = &record.ciphertext();

    CString::new(record_ciphertext.to_string())
        .unwrap()
        .into_raw()
}

#[no_mangle]
pub extern "C" fn decrypt_record(
    ciphertext: *const libc::c_char,
    view_key: *const libc::c_char,
) -> *mut Record<Testnet2> {
    let c_ciphertext = unsafe {
        assert!(!ciphertext.is_null());

        CStr::from_ptr(ciphertext)
    };

    let c_view_key = unsafe {
        assert!(!view_key.is_null());

        CStr::from_ptr(view_key)
    };

    let view_key = ViewKey::<Testnet2>::from_str(c_view_key.to_str().unwrap()).unwrap();

    let encrypted_record =
        <Testnet2 as Network>::RecordCiphertext::from_str(c_ciphertext.to_str().unwrap()).unwrap();

    let record = match Record::from_account_view_key(&view_key, &encrypted_record) {
        Ok(rec) => rec,
        _ => {
            c_error::update_last_error(snarkvm_utilities::error("cannot decrypt ciphertext"));
            return std::ptr::null_mut();
        }
    };
    Box::into_raw(Box::new(record))
}

#[no_mangle]
pub extern "C" fn record_owner(ptr: *mut Record<Testnet2>) -> *mut libc::c_char {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let owner = record.owner().to_string();
    CString::new(owner).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn record_value(ptr: *mut Record<Testnet2>) -> i64 {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };
    let value = record.value().as_i64();
    value
}

#[repr(C)]
pub struct Buffer {
    data: *mut u8,
    len: usize,
}

#[no_mangle]
pub extern "C" fn record_payload(ptr: *mut Record<Testnet2>) -> Buffer {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };
    let payload = record.payload();
    let mut buf = payload.to_bytes_le().unwrap();
    let data = buf.as_mut_ptr();
    let len = buf.len();
    std::mem::forget(buf);
    Buffer { data, len }
}

#[no_mangle]
pub extern "C" fn record_commitment_randomness(ptr: *mut Record<Testnet2>) -> *mut libc::c_char {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let commitment_randomness = record.randomizer().to_string();
    CString::new(commitment_randomness).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn record_commitment(ptr: *mut Record<Testnet2>) -> *mut libc::c_char {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let commitment = record.commitment().to_string();
    CString::new(commitment).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn record_program_id(ptr: *mut Record<Testnet2>) -> *mut libc::c_char {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let program_id = record.program_id().to_string();
    CString::new(program_id).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn record_free(ptr: *mut Record<Testnet2>) {
    if ptr.is_null() {
        return;
    }
    unsafe {
        Box::from_raw(ptr);
    }
}

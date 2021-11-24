// create a output record
// create an input record
// encrypt a record
// decrypt a record

use crate::c_error;
use rand::{rngs::StdRng, SeedableRng};
use snarkvm_dpc::{
    network::testnet2::Testnet2, Address, Network, Payload, Record, RecordCiphertext, ViewKey,
};
use snarkvm_utilities::{FromBytes, ToBytes, UniformRand};
use std::ffi::{CStr, CString};
use std::{slice, str::FromStr};

#[no_mangle]
pub extern "C" fn new_input_record(
    addr: *const libc::c_char,
    val: u64,
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

    let record = match Record::new_input(
        address,
        val as u64,
        record_payload,
        *Testnet2::noop_program_id(),
        UniformRand::rand(&mut rng),
        UniformRand::rand(&mut rng),
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
    val: u64,
    payload: *const u8,
    serial_number_nonce: *const libc::c_char,
    commitment_randomness: *const libc::c_char,
) -> *mut Record<Testnet2> {
    let c_addr = unsafe {
        assert!(!addr.is_null());

        CStr::from_ptr(addr)
    };

    let address = Address::<Testnet2>::from_str(c_addr.to_str().unwrap()).unwrap();
    // convert payload

    let c_payload = unsafe {
        assert!(!payload.is_null());

        slice::from_raw_parts(payload, Testnet2::RECORD_PAYLOAD_SIZE_IN_BYTES)
    };

    // convert serial_number_nonce
    let c_serial_number_nonce = unsafe {
        assert!(!serial_number_nonce.is_null());

        CStr::from_ptr(serial_number_nonce)
    };

    let c_commitment_randomness = unsafe {
        assert!(!commitment_randomness.is_null());
        CStr::from_ptr(commitment_randomness)
    };

    let record = match Record::from(
        address,
        val as u64,
        Payload::from_bytes_le(&c_payload).unwrap(),
        *Testnet2::noop_program_id(),
        <Testnet2 as Network>::SerialNumber::from_str(c_serial_number_nonce.to_str().unwrap())
            .unwrap(),
        <Testnet2 as Network>::CommitmentRandomness::from_str(
            c_commitment_randomness.to_str().unwrap(),
        )
        .unwrap(),
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
pub extern "C" fn encrypt_record(
    ptr: *mut Record<Testnet2>,
    randomness: *const u8,
    randomness_len: libc::size_t,
) -> *mut libc::c_char {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let c_rng = unsafe {
        assert!(!randomness.is_null());
        slice::from_raw_parts(randomness, randomness_len as usize)
    };

    let mut rng: StdRng = SeedableRng::from_seed(c_rng.try_into().unwrap());

    let (record_ciphertext, _) = RecordCiphertext::encrypt(&record, &mut rng).unwrap();
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

    let ciphertext = RecordCiphertext::from_str(c_ciphertext.to_str().unwrap()).unwrap();
    let record = match ciphertext.decrypt(&view_key) {
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
pub extern "C" fn record_value(ptr: *mut Record<Testnet2>) -> u64 {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };
    let value = record.value() as u64;
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
pub extern "C" fn record_serial_number_nonce(ptr: *mut Record<Testnet2>) -> *mut libc::c_char {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let serial_number_nonce = record.serial_number_nonce().to_string();
    CString::new(serial_number_nonce).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn record_commitment_randomness(ptr: *mut Record<Testnet2>) -> *mut libc::c_char {
    let record = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let commitment_randomness = record.commitment_randomness().to_string();
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

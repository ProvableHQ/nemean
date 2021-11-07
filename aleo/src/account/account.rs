use snarkvm_dpc::{
    network::testnet2::Testnet2,
    Account,
    AccountScheme,
};
use std::slice;
use std::ffi::CString;
use rand::{rngs::StdRng, SeedableRng};

#[no_mangle]
pub extern "C" fn from_seed(n: *const u8, len: libc::size_t) -> *mut Account<Testnet2> {
    let buf = unsafe {
        assert!(!n.is_null());

        slice::from_raw_parts(n, len as usize)
    };

    let mut rng: StdRng = SeedableRng::from_seed(buf.try_into().unwrap());

    let account = Account::<Testnet2>::new(&mut rng);
    Box::into_raw(Box::new(account))
}

#[no_mangle]
pub extern "C" fn account_private_key(ptr: *mut Account<Testnet2>) -> *mut libc::c_char {
    let account = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let private_key = account.private_key().to_string();
    CString::new(private_key).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn account_view_key(ptr: *mut Account<Testnet2>) -> *mut libc::c_char {
    let account = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let view_key = account.view_key().to_string();
    CString::new(view_key).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn account_address(ptr: *mut Account<Testnet2>) -> *mut libc::c_char {
    let account = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let address = account.address().to_string();
    CString::new(address).unwrap().into_raw()
}
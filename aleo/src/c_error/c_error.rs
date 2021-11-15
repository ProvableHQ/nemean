/*
c_error spawns a local thread to handle error messages such that the caller can later retrieve the last error.
Relies on https://michael-f-bryan.github.io/rust-ffi-guide/errors/index.html
*/

thread_local! {
    static LAST_ERROR: std::cell::RefCell<Option<Box<dyn std::error::Error>>> =  std::cell::RefCell::new(None);
}

pub fn update_last_error<E: std::error::Error + 'static>(err: E) {
    {
        let mut cause = err.source();
        while let Some(parent_err) = cause {
            cause = parent_err.source();
        }
    }

    LAST_ERROR.with(|prev| {
        *prev.borrow_mut() = Some(Box::new(err));
    });
}

/// Retrieve the most recent error, clearing it in the process.
pub fn take_last_error() -> Option<Box<dyn std::error::Error>> {
    LAST_ERROR.with(|prev| prev.borrow_mut().take())
}

#[no_mangle]
pub extern "C" fn last_error_length() -> libc::c_int {
    LAST_ERROR.with(|prev| match *prev.borrow() {
        Some(ref err) => err.to_string().len() as libc::c_int + 1,
        None => 0,
    })
}

#[no_mangle]
pub unsafe extern "C" fn last_error_message(
    buffer: *mut libc::c_char,
    length: libc::c_int,
) -> libc::c_int {
    if buffer.is_null() {
        return -1;
    }

    let last_error = match take_last_error() {
        Some(err) => err,
        None => return 0,
    };

    let error_message = last_error.to_string();

    let buffer = std::slice::from_raw_parts_mut(buffer as *mut u8, length as usize);

    if error_message.len() >= buffer.len() {
        return -1;
    }

    std::ptr::copy_nonoverlapping(
        error_message.as_ptr(),
        buffer.as_mut_ptr(),
        error_message.len(),
    );

    buffer[error_message.len()] = 0;

    error_message.len() as libc::c_int
}

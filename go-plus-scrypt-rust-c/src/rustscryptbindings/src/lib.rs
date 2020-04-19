#![feature(test)]

extern crate libc;
extern crate scrypt;
extern crate test;

use scrypt::{scrypt, ScryptParams};
use std::slice;
use std::ffi::CStr;
use std::os::raw::c_char;

#[no_mangle]
pub extern fn scrypt_key(raw_password: *const c_char, raw_salt: *const c_char, r: u32, p: u32, n: u8, raw_result: *mut u8, result_size: usize) -> i32 {
    let params = match ScryptParams::new(n, r, p) {
        Ok(p) => p,
        Err(_e) => return -1
    };

    unsafe {
        let password = CStr::from_ptr(raw_password);
        let salt = CStr::from_ptr(raw_salt);
        let result = slice::from_raw_parts_mut(raw_result, result_size);

        match scrypt(password.to_bytes(), salt.to_bytes(), &params, result) {
            Ok(_) => 0,
            Err(_) => -1,
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::scrypt_key;
    use std::ffi::CString;
    use test::Bencher;

    #[test]
    fn it_works() {
        let mut result: Vec<u8> = vec![0; 32];
        let password = CString::new("test1").unwrap();
        let salt = CString::new("salt1").unwrap();

        scrypt_key(
            password.as_ptr(),
            salt.as_ptr(),
            8, 1, 16,
            result.as_mut_ptr(),
            32,
        );

        assert_eq!(
            result,
            vec![155, 14, 28, 179, 58, 64, 198, 176, 74, 48, 161, 168, 12, 134, 150, 179, 43, 33, 197, 179, 150, 167, 44, 227, 168, 25, 55, 244, 63, 24, 242, 107]
        );
    }

    #[bench]
    fn bench_scrypt_key(b: &mut Bencher) {
        let mut result: Vec<u8> = vec![0; 32];
        let password = CString::new("test1").unwrap();
        let salt = CString::new("salt1").unwrap();

        b.iter(|| scrypt_key(
            password.as_ptr(),
            salt.as_ptr(),
            8, 1, 16,
            result.as_mut_ptr(),
            32,
        ));
    }
}

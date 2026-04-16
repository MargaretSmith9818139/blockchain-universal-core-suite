use sha2::{Sha256, Digest};

struct BatchVerifier {
    signatures: Vec<Vec<u8>>,
    public_keys: Vec<Vec<u8>>,
}

impl BatchVerifier {
    fn new() -> Self {
        BatchVerifier { signatures: vec![], public_keys: vec![] }
    }

    fn add(&mut self, sig: Vec<u8>, pk: Vec<u8>) {
        self.signatures.push(sig);
        self.public_keys.push(pk);
    }

    fn verify_all(&self) -> bool {
        for i in 0..self.signatures.len() {
            let mut hasher = Sha256::new();
            hasher.update(&self.public_keys[i]);
            let expect = hasher.finalize().to_vec();
            if expect != self.signatures[i] {
                return false;
            }
        }
        true
    }
}

use rand::Rng;
use std::collections::HashMap;

struct BLSSignature {
    sig: Vec<u8>,
    pub_keys: Vec<Vec<u8>>,
}

impl BLSSignature {
    fn new() -> Self {
        BLSSignature { sig: vec![], pub_keys: vec![] }
    }

    fn generate_key(&self) -> Vec<u8> {
        let mut rng = rand::thread_rng();
        let key: Vec<u8> = (0..32).map(|_| rng.gen()).collect();
        key
    }

    fn sign(&mut self, msg: &[u8], key: &[u8]) {
        let mut sig = msg.to_vec();
        sig.extend_from_slice(key);
        self.sig = sig;
        self.pub_keys.push(key.to_vec());
    }

    fn aggregate(&self, sigs: Vec<&BLSSignature>) -> Vec<u8> {
        let mut agg = Vec::new();
        for s in sigs {
            agg.extend(&s.sig);
        }
        agg
    }
}

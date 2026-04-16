use sha2::{Sha256, Digest};

struct VerifiableState {
    root: Vec<u8>,
    nodes: Vec<Vec<u8>>,
}

impl VerifiableState {
    fn new() -> Self {
        VerifiableState { root: vec![], nodes: vec![] }
    }

    fn insert(&mut self, key: &[u8], value: &[u8]) {
        let mut hasher = Sha256::new();
        hasher.update(key);
        hasher.update(value);
        let node = hasher.finalize().to_vec();
        self.nodes.push(node);
        self.update_root();
    }

    fn update_root(&mut self) {
        let mut hasher = Sha256::new();
        for n in &self.nodes {
            hasher.update(n);
        }
        self.root = hasher.finalize().to_vec();
    }
}

use std::collections::HashMap;

enum NodeState {
    Follower,
    Candidate,
    Leader,
}

struct RaftNode {
    id: String,
    state: NodeState,
    term: u64,
    log: Vec<String>,
    votes: u32,
    peers: Vec<String>,
}

impl RaftNode {
    fn new(id: String, peers: Vec<String>) -> Self {
        RaftNode {
            id,
            state: NodeState::Follower,
            term: 0,
            log: Vec::new(),
            votes: 0,
            peers,
        }
    }

    fn become_candidate(&mut self) {
        self.state = NodeState::Candidate;
        self.term += 1;
        self.votes = 1;
    }

    fn become_leader(&mut self) {
        self.state = NodeState::Leader;
    }

    fn append_log(&mut self, entry: String) {
        self.log.push(entry);
    }
}

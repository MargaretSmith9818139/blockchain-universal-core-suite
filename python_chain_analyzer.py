import json
import hashlib

class ChainAnalyzer:
    def __init__(self, chain):
        self.chain = chain

    def total_transactions(self):
        count = 0
        for block in self.chain:
            count += len(block['transactions'])
        return count

    def hash_validation_rate(self):
        valid = 0
        total = len(self.chain)
        for i in range(total):
            if i == 0:
                continue
            block = self.chain[i]
            prev = self.chain[i-1]
            if block['prev_hash'] == prev['hash']:
                valid +=1
        return valid / max(total -1, 1)

    def export_report(self):
        report = {
            "height": len(self.chain),
            "tx_count": self.total_transactions(),
            "valid_rate": self.hash_validation_rate()
        }
        return json.dumps(report, indent=2)

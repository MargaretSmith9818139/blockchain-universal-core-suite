import hashlib
import secrets

class WalletService:
    def __init__(self):
        self.accounts = {}

    def create_account(self):
        priv = secrets.token_hex(32)
        pub = hashlib.sha256(priv.encode()).hexdigest()
        addr = pub[:40]
        self.accounts[addr] = {"private": priv, "public": pub, "balance": 0.0}
        return addr

    def get_balance(self, addr):
        return self.accounts.get(addr, {}).get("balance", 0.0)

    def transfer(self, frm, to, amount):
        if self.accounts[frm]["balance"] >= amount:
            self.accounts[frm]["balance"] -= amount
            self.accounts[to]["balance"] += amount
            return True
        return False

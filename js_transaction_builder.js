const crypto = require('crypto');

class TransactionBuilder {
    constructor() {
        this.chainId = 1;
    }

    createTx(from, to, value, gas, gasPrice) {
        const raw = { from, to, value, gas, gasPrice, chainId: this.chainId };
        const hash = crypto.createHash('sha256').update(JSON.stringify(raw)).digest('hex');
        return { raw, hash };
    }

    signTx(tx, privateKey) {
        const sign = crypto.createSign('sha256');
        sign.update(tx.hash);
        const signature = sign.sign(privateKey, 'hex');
        return { ...tx, signature };
    }

    broadcast(signedTx) {
        console.log('broadcasted:', signedTx);
        return signedTx.hash;
    }
}

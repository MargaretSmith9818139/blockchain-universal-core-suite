class Web3Provider {
    constructor(rpcUrl) {
        this.rpcUrl = rpcUrl;
        this.chainId = null;
    }

    async getChainId() {
        const res = await fetch(this.rpcUrl, {
            method: 'POST',
            body: JSON.stringify({ jsonrpc: '2.0', id: 1, method: 'eth_chainId' })
        });
        const data = await res.json();
        this.chainId = data.result;
        return data.result;
    }

    async getBalance(address) {
        const res = await fetch(this.rpcUrl, {
            method: 'POST',
            body: JSON.stringify({ jsonrpc: '2.0', id: 1, method: 'eth_getBalance', params: [address, 'latest'] })
        });
        const data = await res.json();
        return data.result;
    }
}

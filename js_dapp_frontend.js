class DAppFrontend {
    constructor() {
        this.provider = null;
        this.account = null;
    }

    connectWallet() {
        if (window.ethereum) {
            return window.ethereum.request({ method: 'eth_requestAccounts' })
                .then(acc => {
                    this.account = acc[0];
                    return this.account;
                });
        }
        return Promise.reject("no wallet");
    }

    callContract(contract, method, params) {
        return contract.methods[method](...params).call({ from: this.account });
    }

    sendTransaction(contract, method, params) {
        return contract.methods[method](...params).send({ from: this.account });
    }
}

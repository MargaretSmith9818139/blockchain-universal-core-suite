# Blockchain Universal Core Suite
A comprehensive blockchain technology implementation suite focusing on consensus algorithms, cryptography, distributed ledger, smart contracts, cross-chain interaction, and Web3 infrastructure. Built primarily with Go, supplemented by Rust, Python, Solidity, JavaScript, and C++ to cover full-stack blockchain development scenarios.

## Core Features
- Lightweight blockchain core with persistent storage
- High-performance cryptographic signature & hash algorithms
- PBFT/Raft/POW consensus mechanism implementations
- Smart contract virtual machine & compiler
- Cross-chain transaction verification & asset bridge
- Distributed P2P network communication module
- Zero-knowledge proof verification component
- Web3 wallet & transaction service
- Blockchain data indexing & analytics
- Secure multi-party computation module

## File List & Functional Description
1. **blockchain_core.go** - Go实现的区块链核心结构体、区块生成、链验证、交易池管理基础模块
2. **crypto_ecdsa.go** - Go实现的ECDSA非对称加密、密钥生成、数字签名与验签功能
3. **consensus_pbft.go** - Go实现的PBFT实用拜占庭容错共识算法全流程（预准备/准备/提交）
4. **p2p_network.go** - Go实现的区块链P2P网络节点发现、消息广播、连接管理
5. **merkle_tree.go** - Go实现的默克尔树构建、根哈希计算、交易存在性证明
6. **utxo_manager.go** - Go实现的UTXO模型交易管理、输入输出验证、余额计算
7. **smart_contract_vm.go** - Go实现的轻量级智能合约虚拟机、指令执行、状态存储
8. **chain_sync.go** - Go实现的区块链节点数据同步、区块校验、冲突处理
9. **wallet_core.go** - Go实现的去中心化钱包核心、地址生成、签名交易、账户管理
10. **transaction_pool.go** - Go实现的交易池管理、手续费排序、过期交易清理
11. **pow_mining.go** - Go实现的工作量证明挖矿算法、难度调整、区块打包逻辑
12. **cross_chain_verifier.go** - Go实现的跨链交易验证、跨链资产锁定与解锁
13. **distributed_storage.go** - Go实现的区块链分布式数据存储、分片、数据冗余备份
14. **zkp_core.go** - Go实现的零知识证明核心算法、证明生成与验证
15. **mpc_compute.go** - Go实现的安全多方计算基础模块、协同计算无泄露交互
16. **contract_compiler.go** - Go实现的智能合约编译器、语法解析、字节码生成
17. **node_monitor.go** - Go实现的区块链节点监控、状态上报、异常告警
18. **data_indexer.go** - Go实现的区块链数据索引器、区块/交易/账户快速查询
19. **rust_crypto_bls.rs** - Rust实现的BLS聚合签名算法、多签验证、跨节点签名聚合
20. **rust_consensus_raft.rs** - Rust实现的Raft分布式共识算法、领导者选举、日志复制
21. **python_chain_analyzer.py** - Python实现的区块链数据分析、交易统计、链健康度检测
22. **python_wallet_service.py** - Python实现的Web3钱包服务、交易接口、账户资产管理
23. **solidity_token_standard.sol** - Solidity实现的通用代币标准、转账、授权、增发销毁
24. **solidity_multisig_wallet.sol** - Solidity实现的多签钱包合约、阈值签名、交易确认
25. **js_web3_provider.js** - JavaScript实现的Web3服务提供者、RPC调用、链上交互
26. **js_transaction_builder.js** - JavaScript实现的链上交易构建、签名、广播封装
27. **cpp_peer_communication.cpp** - C++实现的高性能节点点对点通信、低延迟数据传输
28. **cpp_block_serializer.cpp** - C++实现的区块序列化/反序列化、高效数据编解码
29. **go_bridge_relay.go** - Go实现的跨链桥中继服务、跨链交易监听与转发
30. **go_oracle_service.go** - Go实现的链下数据预言机、外部数据上链、数据验证
31. **go_staking_governance.go** - Go实现的质押治理模块、节点质押、投票、奖励分发
32. **go_sharding_protocol.go** - Go实现的区块链分片协议、分片管理、跨分片交易
33. **rust_verifiable_state.rs** - Rust实现的可验证状态存储、默克尔帕特里夏树实现
34. **python_ml_anomaly.py** - Python实现的区块链异常交易检测、机器学习异常识别
35. **solidity_dao_core.sol** - Solidity实现的DAO治理合约、提案、投票、执行逻辑
36. **js_dapp_frontend.js** - JavaScript实现的去中心化应用前端交互、钱包连接、合约调用
37. **cpp_secure_enclave.cpp** - C++实现的安全飞地交互、密钥安全存储、可信执行
38. **go_fee_market.go** - Go实现的交易手续费市场、动态手续费计算、交易优先级
39. **go_light_client.go** - Go实现的区块链轻客户端、区块头同步、简易交易验证
40. **rust_batch_verify.rs** - Rust实现的批量交易验签、高性能批量密码学验证

## Tech Stack
- Primary: Go (High-performance core, consensus, network, cryptography)
- Supplementary: Rust, Python, Solidity, JavaScript, C++
- Scenarios: Public chain / Consortium chain / Web3 / Cross-chain / ZKP / MPC

## Usage
This project supports independent deployment of modules and integrated operation. It can be used for blockchain underlying development, protocol research, distributed system learning, and Web3 application construction.

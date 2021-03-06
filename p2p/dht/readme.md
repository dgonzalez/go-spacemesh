## SPACEMESH-DHT
In this package we'd like to implement a modern [DHT](https://en.wikipedia.org/wiki/Distributed_hash_table) with the following main properties:

1. All comm between nodes is secure and authenticated (using spacemesh-swarm). This is sometimes known as s-kademlia.
2. [Coral](https://en.wikipedia.org/wiki/Coral_Content_Distribution_Network) improvements to Kademlia content distribution. aka [mainlin dht](https://en.wikipedia.org/wiki/Mainline_DHT). 
Implemented in libp2p-kad-dht and bittorent clients
3. Avoid the complexity and possible race conditions involved in `go-eth-p2p-kad` and `libp2p-kad-dht`. 
libp2p open-issues list is somewhat scray and libp2p-kad-dht is tightly coupled with its main client - ipfs.
4. Have as little external deps as possible - e.g. multi-address, specific key-id schemes, etc...

This package contains heavily modified code inspired by `libp2p-kad-dht` and from `libp2p-kbucket` - the license file for both of these packages is included in this package as required by the MIT licensing terms.


# Wasp & Chain Info

HORNET:
    ports:
      - "14265:14265/tcp" # API
      - "8091:8091/tcp" # Faucet
      - "9029:9029/tcp" # INX

WASP:
    env:
      - WASP_API_URL = http://wasp.nearfieldquery.io
      - L1_API_URL = http://hornet.nearfieldquery.io
    ports:
      - "2112:2112/tcp" # Prometheus
      - "4000:4000/tcp" # Peering
      - "9090:9090/tcp" # WebAPI
      - "6060:6060/tcp" # Profiling

Address: tst1qp452409f9lmerlw6xmlyu5edkwv0q7vu9s6p7us2rvhu7yqy02gyex7rzd

* committee size = 1
* quorum = 1
* members: 0x4eabcfb03f99bfd9040e56a1f91d8d88767ca70d89c49ff6a879879e979e47b6 (me)

Creating new chain

* Owner address:    0x417df6d2c7d9a8e3400f8ced50d4198f6a1a386aad67488b221a4f5d47d13246
* State controller: 0x6b4555e5497fbc8feed1b7f272996d9cc783cce161a0fb9050d97e788023d482
* committee size = 2
* quorum = 0
2023-09-02T21:43:13-06:00       INFO    nc      Posted blockID 0xac6bf3c7e0263d74f8ec8b5f5493418aa183936bfb0457ce21f5126e7567abbe
2023-09-02T21:43:13-06:00       INFO    nc      Posted transaction id 0x7460afa48610756f6c8584b5858fca80360f679e72a7da44dc5b2c0da6d596da
Chain has been created successfully on the Tangle.
* ChainID: tst1pqvk3kkw9mm7ncw7pyvpngaxm2nx0tzlwr5wthddp5pj86tvssl353hqccc
* State address: tst1qp452409f9lmerlw6xmlyu5edkwv0q7vu9s6p7us2rvhu7yqy02gyex7rzd
* committee size = 2
* quorum = 0
Make sure to activate the chain on all committee nodes
Chain: tst1pqvk3kkw9mm7ncw7pyvpngaxm2nx0tzlwr5wthddp5pj86tvssl353hqccc (nearfieldquery-evm)

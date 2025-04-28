keplerd query auth account "$(keplerd keys show alice -a)" -o json

# ditto1cug2mhrhuzw99rag90w4m0e8jpj5lzhw5hhhgr
keplerd keys add alice-multisig --multisig alice --multisig-threshold 1

keplerd tx committee send-report \
  --from alice-multisig \
  --fees 20ditto \
  --chain-id kepler \
  --generate-only > tx.json

keplerd tx sign --from alice --multisig=ditto1cug2mhrhuzw99rag90w4m0e8jpj5lzhw5hhhgr tx.json --chain-id kepler > tx-signed-alice.json

keplerd tx multisign tx.json alice-multisig tx-signed-alice.json --chain-id kepler > tx-signed.json
keplerd tx broadcast tx-signed.json --chain-id kepler --gas auto

keplerd query tx {tx_hash_here}

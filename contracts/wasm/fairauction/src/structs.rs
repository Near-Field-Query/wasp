// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

#[derive(Clone)]
pub struct Auction {
    pub creator        : ScAgentID, // issuer of start_auction transaction
    pub deposit        : u64, // deposit by auction owner to cover the SC fees
    pub description    : String, // auction description
    pub duration       : u32, // auction duration in minutes
    pub highest_bid    : u64, // the current highest bid amount
    pub highest_bidder : ScAgentID, // the current highest bidder
    pub minimum_bid    : u64, // minimum bid amount
    pub num_tokens     : u64, // number of tokens for sale
    pub owner_margin   : u64, // auction owner's margin in promilles
    pub token          : ScTokenID, // token of tokens for sale
    pub when_started   : u64, // timestamp when auction started
}

impl Auction {
    pub fn from_bytes(bytes: &[u8]) -> Auction {
        let mut dec = WasmDecoder::new(bytes);
        Auction {
            creator        : agent_id_decode(&mut dec),
            deposit        : uint64_decode(&mut dec),
            description    : string_decode(&mut dec),
            duration       : uint32_decode(&mut dec),
            highest_bid    : uint64_decode(&mut dec),
            highest_bidder : agent_id_decode(&mut dec),
            minimum_bid    : uint64_decode(&mut dec),
            num_tokens     : uint64_decode(&mut dec),
            owner_margin   : uint64_decode(&mut dec),
            token          : token_id_decode(&mut dec),
            when_started   : uint64_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		agent_id_encode(&mut enc, &self.creator);
		uint64_encode(&mut enc, self.deposit);
		string_encode(&mut enc, &self.description);
		uint32_encode(&mut enc, self.duration);
		uint64_encode(&mut enc, self.highest_bid);
		agent_id_encode(&mut enc, &self.highest_bidder);
		uint64_encode(&mut enc, self.minimum_bid);
		uint64_encode(&mut enc, self.num_tokens);
		uint64_encode(&mut enc, self.owner_margin);
		token_id_encode(&mut enc, &self.token);
		uint64_encode(&mut enc, self.when_started);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableAuction {
    pub(crate) proxy: Proxy,
}

impl ImmutableAuction {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> Auction {
        Auction::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableAuction {
    pub(crate) proxy: Proxy,
}

impl MutableAuction {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &Auction) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> Auction {
        Auction::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct Bid {
    pub amount    : u64, // cumulative amount of bids from same bidder
    pub index     : u32, // index of bidder in bidder list
    pub timestamp : u64, // timestamp of most recent bid
}

impl Bid {
    pub fn from_bytes(bytes: &[u8]) -> Bid {
        let mut dec = WasmDecoder::new(bytes);
        Bid {
            amount    : uint64_decode(&mut dec),
            index     : uint32_decode(&mut dec),
            timestamp : uint64_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		uint64_encode(&mut enc, self.amount);
		uint32_encode(&mut enc, self.index);
		uint64_encode(&mut enc, self.timestamp);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableBid {
    pub(crate) proxy: Proxy,
}

impl ImmutableBid {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> Bid {
        Bid::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableBid {
    pub(crate) proxy: Proxy,
}

impl MutableBid {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &Bid) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> Bid {
        Bid::from_bytes(&self.proxy.get())
    }
}

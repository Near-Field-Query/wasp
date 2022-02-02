// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ImmutableMintSupplyParams extends wasmlib.ScMapID {
    description(): wasmlib.ScImmutableString {
		return new wasmlib.ScImmutableString(this.mapID, wasmlib.Key32.fromString(sc.ParamDescription));
	}

    userDefined(): wasmlib.ScImmutableString {
		return new wasmlib.ScImmutableString(this.mapID, wasmlib.Key32.fromString(sc.ParamUserDefined));
	}
}

export class MutableMintSupplyParams extends wasmlib.ScMapID {
    description(): wasmlib.ScMutableString {
		return new wasmlib.ScMutableString(this.mapID, wasmlib.Key32.fromString(sc.ParamDescription));
	}

    userDefined(): wasmlib.ScMutableString {
		return new wasmlib.ScMutableString(this.mapID, wasmlib.Key32.fromString(sc.ParamUserDefined));
	}
}

export class ImmutableTransferOwnershipParams extends wasmlib.ScMapID {
    color(): wasmlib.ScImmutableColor {
		return new wasmlib.ScImmutableColor(this.mapID, wasmlib.Key32.fromString(sc.ParamColor));
	}
}

export class MutableTransferOwnershipParams extends wasmlib.ScMapID {
    color(): wasmlib.ScMutableColor {
		return new wasmlib.ScMutableColor(this.mapID, wasmlib.Key32.fromString(sc.ParamColor));
	}
}

export class ImmutableUpdateMetadataParams extends wasmlib.ScMapID {
    color(): wasmlib.ScImmutableColor {
		return new wasmlib.ScImmutableColor(this.mapID, wasmlib.Key32.fromString(sc.ParamColor));
	}
}

export class MutableUpdateMetadataParams extends wasmlib.ScMapID {
    color(): wasmlib.ScMutableColor {
		return new wasmlib.ScMutableColor(this.mapID, wasmlib.Key32.fromString(sc.ParamColor));
	}
}

export class ImmutableGetInfoParams extends wasmlib.ScMapID {
    color(): wasmlib.ScImmutableColor {
		return new wasmlib.ScImmutableColor(this.mapID, wasmlib.Key32.fromString(sc.ParamColor));
	}
}

export class MutableGetInfoParams extends wasmlib.ScMapID {
    color(): wasmlib.ScMutableColor {
		return new wasmlib.ScMutableColor(this.mapID, wasmlib.Key32.fromString(sc.ParamColor));
	}
}

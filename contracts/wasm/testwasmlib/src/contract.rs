// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;
use crate::*;

pub struct AddressMapOfAddressArrayAppendCall {
	pub func: ScFunc,
	pub params: MutableAddressMapOfAddressArrayAppendParams,
}

pub struct AddressMapOfAddressArrayClearCall {
	pub func: ScFunc,
	pub params: MutableAddressMapOfAddressArrayClearParams,
}

pub struct AddressMapOfAddressArraySetCall {
	pub func: ScFunc,
	pub params: MutableAddressMapOfAddressArraySetParams,
}

pub struct AddressMapOfAddressMapClearCall {
	pub func: ScFunc,
	pub params: MutableAddressMapOfAddressMapClearParams,
}

pub struct AddressMapOfAddressMapSetCall {
	pub func: ScFunc,
	pub params: MutableAddressMapOfAddressMapSetParams,
}

pub struct ArrayOfAddressArrayAppendCall {
	pub func: ScFunc,
	pub params: MutableArrayOfAddressArrayAppendParams,
}

pub struct ArrayOfAddressArrayClearCall {
	pub func: ScFunc,
}

pub struct ArrayOfAddressArraySetCall {
	pub func: ScFunc,
	pub params: MutableArrayOfAddressArraySetParams,
}

pub struct ArrayOfAddressMapClearCall {
	pub func: ScFunc,
}

pub struct ArrayOfAddressMapSetCall {
	pub func: ScFunc,
	pub params: MutableArrayOfAddressMapSetParams,
}

pub struct ArrayOfStringArrayAppendCall {
	pub func: ScFunc,
	pub params: MutableArrayOfStringArrayAppendParams,
}

pub struct ArrayOfStringArrayClearCall {
	pub func: ScFunc,
}

pub struct ArrayOfStringArraySetCall {
	pub func: ScFunc,
	pub params: MutableArrayOfStringArraySetParams,
}

pub struct ArrayOfStringMapClearCall {
	pub func: ScFunc,
}

pub struct ArrayOfStringMapSetCall {
	pub func: ScFunc,
	pub params: MutableArrayOfStringMapSetParams,
}

pub struct ParamTypesCall {
	pub func: ScFunc,
	pub params: MutableParamTypesParams,
}

pub struct RandomCall {
	pub func: ScFunc,
}

pub struct StringMapOfStringArrayAppendCall {
	pub func: ScFunc,
	pub params: MutableStringMapOfStringArrayAppendParams,
}

pub struct StringMapOfStringArrayClearCall {
	pub func: ScFunc,
	pub params: MutableStringMapOfStringArrayClearParams,
}

pub struct StringMapOfStringArraySetCall {
	pub func: ScFunc,
	pub params: MutableStringMapOfStringArraySetParams,
}

pub struct StringMapOfStringMapClearCall {
	pub func: ScFunc,
	pub params: MutableStringMapOfStringMapClearParams,
}

pub struct StringMapOfStringMapSetCall {
	pub func: ScFunc,
	pub params: MutableStringMapOfStringMapSetParams,
}

pub struct TakeAllowanceCall {
	pub func: ScFunc,
}

pub struct TakeBalanceCall {
	pub func: ScFunc,
	pub results: ImmutableTakeBalanceResults,
}

pub struct TriggerEventCall {
	pub func: ScFunc,
	pub params: MutableTriggerEventParams,
}

pub struct AddressMapOfAddressArrayLengthCall {
	pub func: ScView,
	pub params: MutableAddressMapOfAddressArrayLengthParams,
	pub results: ImmutableAddressMapOfAddressArrayLengthResults,
}

pub struct AddressMapOfAddressArrayValueCall {
	pub func: ScView,
	pub params: MutableAddressMapOfAddressArrayValueParams,
	pub results: ImmutableAddressMapOfAddressArrayValueResults,
}

pub struct AddressMapOfAddressMapValueCall {
	pub func: ScView,
	pub params: MutableAddressMapOfAddressMapValueParams,
	pub results: ImmutableAddressMapOfAddressMapValueResults,
}

pub struct ArrayOfAddressArrayLengthCall {
	pub func: ScView,
	pub results: ImmutableArrayOfAddressArrayLengthResults,
}

pub struct ArrayOfAddressArrayValueCall {
	pub func: ScView,
	pub params: MutableArrayOfAddressArrayValueParams,
	pub results: ImmutableArrayOfAddressArrayValueResults,
}

pub struct ArrayOfAddressMapValueCall {
	pub func: ScView,
	pub params: MutableArrayOfAddressMapValueParams,
	pub results: ImmutableArrayOfAddressMapValueResults,
}

pub struct ArrayOfStringArrayLengthCall {
	pub func: ScView,
	pub results: ImmutableArrayOfStringArrayLengthResults,
}

pub struct ArrayOfStringArrayValueCall {
	pub func: ScView,
	pub params: MutableArrayOfStringArrayValueParams,
	pub results: ImmutableArrayOfStringArrayValueResults,
}

pub struct ArrayOfStringMapValueCall {
	pub func: ScView,
	pub params: MutableArrayOfStringMapValueParams,
	pub results: ImmutableArrayOfStringMapValueResults,
}

pub struct BigIntAddCall {
	pub func: ScView,
	pub params: MutableBigIntAddParams,
	pub results: ImmutableBigIntAddResults,
}

pub struct BigIntDivCall {
	pub func: ScView,
	pub params: MutableBigIntDivParams,
	pub results: ImmutableBigIntDivResults,
}

pub struct BigIntModCall {
	pub func: ScView,
	pub params: MutableBigIntModParams,
	pub results: ImmutableBigIntModResults,
}

pub struct BigIntMulCall {
	pub func: ScView,
	pub params: MutableBigIntMulParams,
	pub results: ImmutableBigIntMulResults,
}

pub struct BigIntShlCall {
	pub func: ScView,
	pub params: MutableBigIntShlParams,
	pub results: ImmutableBigIntShlResults,
}

pub struct BigIntShrCall {
	pub func: ScView,
	pub params: MutableBigIntShrParams,
	pub results: ImmutableBigIntShrResults,
}

pub struct BigIntSubCall {
	pub func: ScView,
	pub params: MutableBigIntSubParams,
	pub results: ImmutableBigIntSubResults,
}

pub struct BlockRecordCall {
	pub func: ScView,
	pub params: MutableBlockRecordParams,
	pub results: ImmutableBlockRecordResults,
}

pub struct BlockRecordsCall {
	pub func: ScView,
	pub params: MutableBlockRecordsParams,
	pub results: ImmutableBlockRecordsResults,
}

pub struct CheckAgentIDCall {
	pub func: ScView,
	pub params: MutableCheckAgentIDParams,
}

pub struct GetRandomCall {
	pub func: ScView,
	pub results: ImmutableGetRandomResults,
}

pub struct IotaBalanceCall {
	pub func: ScView,
	pub results: ImmutableIotaBalanceResults,
}

pub struct StringMapOfStringArrayLengthCall {
	pub func: ScView,
	pub params: MutableStringMapOfStringArrayLengthParams,
	pub results: ImmutableStringMapOfStringArrayLengthResults,
}

pub struct StringMapOfStringArrayValueCall {
	pub func: ScView,
	pub params: MutableStringMapOfStringArrayValueParams,
	pub results: ImmutableStringMapOfStringArrayValueResults,
}

pub struct StringMapOfStringMapValueCall {
	pub func: ScView,
	pub params: MutableStringMapOfStringMapValueParams,
	pub results: ImmutableStringMapOfStringMapValueResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn address_map_of_address_array_append(_ctx: &dyn ScFuncCallContext) -> AddressMapOfAddressArrayAppendCall {
        let mut f = AddressMapOfAddressArrayAppendCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ADDRESS_MAP_OF_ADDRESS_ARRAY_APPEND),
            params: MutableAddressMapOfAddressArrayAppendParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn address_map_of_address_array_clear(_ctx: &dyn ScFuncCallContext) -> AddressMapOfAddressArrayClearCall {
        let mut f = AddressMapOfAddressArrayClearCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ADDRESS_MAP_OF_ADDRESS_ARRAY_CLEAR),
            params: MutableAddressMapOfAddressArrayClearParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn address_map_of_address_array_set(_ctx: &dyn ScFuncCallContext) -> AddressMapOfAddressArraySetCall {
        let mut f = AddressMapOfAddressArraySetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ADDRESS_MAP_OF_ADDRESS_ARRAY_SET),
            params: MutableAddressMapOfAddressArraySetParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn address_map_of_address_map_clear(_ctx: &dyn ScFuncCallContext) -> AddressMapOfAddressMapClearCall {
        let mut f = AddressMapOfAddressMapClearCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ADDRESS_MAP_OF_ADDRESS_MAP_CLEAR),
            params: MutableAddressMapOfAddressMapClearParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn address_map_of_address_map_set(_ctx: &dyn ScFuncCallContext) -> AddressMapOfAddressMapSetCall {
        let mut f = AddressMapOfAddressMapSetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ADDRESS_MAP_OF_ADDRESS_MAP_SET),
            params: MutableAddressMapOfAddressMapSetParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn array_of_address_array_append(_ctx: &dyn ScFuncCallContext) -> ArrayOfAddressArrayAppendCall {
        let mut f = ArrayOfAddressArrayAppendCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_ADDRESS_ARRAY_APPEND),
            params: MutableArrayOfAddressArrayAppendParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    // Clear all the arrays of the array
    pub fn array_of_address_array_clear(_ctx: &dyn ScFuncCallContext) -> ArrayOfAddressArrayClearCall {
        ArrayOfAddressArrayClearCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_ADDRESS_ARRAY_CLEAR),
        }
    }

    pub fn array_of_address_array_set(_ctx: &dyn ScFuncCallContext) -> ArrayOfAddressArraySetCall {
        let mut f = ArrayOfAddressArraySetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_ADDRESS_ARRAY_SET),
            params: MutableArrayOfAddressArraySetParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn array_of_address_map_clear(_ctx: &dyn ScFuncCallContext) -> ArrayOfAddressMapClearCall {
        ArrayOfAddressMapClearCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_ADDRESS_MAP_CLEAR),
        }
    }

    pub fn array_of_address_map_set(_ctx: &dyn ScFuncCallContext) -> ArrayOfAddressMapSetCall {
        let mut f = ArrayOfAddressMapSetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_ADDRESS_MAP_SET),
            params: MutableArrayOfAddressMapSetParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn array_of_string_array_append(_ctx: &dyn ScFuncCallContext) -> ArrayOfStringArrayAppendCall {
        let mut f = ArrayOfStringArrayAppendCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_STRING_ARRAY_APPEND),
            params: MutableArrayOfStringArrayAppendParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    // Clear all the arrays of the array
    pub fn array_of_string_array_clear(_ctx: &dyn ScFuncCallContext) -> ArrayOfStringArrayClearCall {
        ArrayOfStringArrayClearCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_STRING_ARRAY_CLEAR),
        }
    }

    pub fn array_of_string_array_set(_ctx: &dyn ScFuncCallContext) -> ArrayOfStringArraySetCall {
        let mut f = ArrayOfStringArraySetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_STRING_ARRAY_SET),
            params: MutableArrayOfStringArraySetParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn array_of_string_map_clear(_ctx: &dyn ScFuncCallContext) -> ArrayOfStringMapClearCall {
        ArrayOfStringMapClearCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_STRING_MAP_CLEAR),
        }
    }

    pub fn array_of_string_map_set(_ctx: &dyn ScFuncCallContext) -> ArrayOfStringMapSetCall {
        let mut f = ArrayOfStringMapSetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_ARRAY_OF_STRING_MAP_SET),
            params: MutableArrayOfStringMapSetParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn param_types(_ctx: &dyn ScFuncCallContext) -> ParamTypesCall {
        let mut f = ParamTypesCall {
            func: ScFunc::new(HSC_NAME, HFUNC_PARAM_TYPES),
            params: MutableParamTypesParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn random(_ctx: &dyn ScFuncCallContext) -> RandomCall {
        RandomCall {
            func: ScFunc::new(HSC_NAME, HFUNC_RANDOM),
        }
    }

    pub fn string_map_of_string_array_append(_ctx: &dyn ScFuncCallContext) -> StringMapOfStringArrayAppendCall {
        let mut f = StringMapOfStringArrayAppendCall {
            func: ScFunc::new(HSC_NAME, HFUNC_STRING_MAP_OF_STRING_ARRAY_APPEND),
            params: MutableStringMapOfStringArrayAppendParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn string_map_of_string_array_clear(_ctx: &dyn ScFuncCallContext) -> StringMapOfStringArrayClearCall {
        let mut f = StringMapOfStringArrayClearCall {
            func: ScFunc::new(HSC_NAME, HFUNC_STRING_MAP_OF_STRING_ARRAY_CLEAR),
            params: MutableStringMapOfStringArrayClearParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn string_map_of_string_array_set(_ctx: &dyn ScFuncCallContext) -> StringMapOfStringArraySetCall {
        let mut f = StringMapOfStringArraySetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_STRING_MAP_OF_STRING_ARRAY_SET),
            params: MutableStringMapOfStringArraySetParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn string_map_of_string_map_clear(_ctx: &dyn ScFuncCallContext) -> StringMapOfStringMapClearCall {
        let mut f = StringMapOfStringMapClearCall {
            func: ScFunc::new(HSC_NAME, HFUNC_STRING_MAP_OF_STRING_MAP_CLEAR),
            params: MutableStringMapOfStringMapClearParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn string_map_of_string_map_set(_ctx: &dyn ScFuncCallContext) -> StringMapOfStringMapSetCall {
        let mut f = StringMapOfStringMapSetCall {
            func: ScFunc::new(HSC_NAME, HFUNC_STRING_MAP_OF_STRING_MAP_SET),
            params: MutableStringMapOfStringMapSetParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn take_allowance(_ctx: &dyn ScFuncCallContext) -> TakeAllowanceCall {
        TakeAllowanceCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TAKE_ALLOWANCE),
        }
    }

    pub fn take_balance(_ctx: &dyn ScFuncCallContext) -> TakeBalanceCall {
        let mut f = TakeBalanceCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TAKE_BALANCE),
            results: ImmutableTakeBalanceResults { proxy: Proxy::nil() },
        };
        ScFunc::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn trigger_event(_ctx: &dyn ScFuncCallContext) -> TriggerEventCall {
        let mut f = TriggerEventCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TRIGGER_EVENT),
            params: MutableTriggerEventParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn address_map_of_address_array_length(_ctx: &dyn ScViewCallContext) -> AddressMapOfAddressArrayLengthCall {
        let mut f = AddressMapOfAddressArrayLengthCall {
            func: ScView::new(HSC_NAME, HVIEW_ADDRESS_MAP_OF_ADDRESS_ARRAY_LENGTH),
            params: MutableAddressMapOfAddressArrayLengthParams { proxy: Proxy::nil() },
            results: ImmutableAddressMapOfAddressArrayLengthResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn address_map_of_address_array_value(_ctx: &dyn ScViewCallContext) -> AddressMapOfAddressArrayValueCall {
        let mut f = AddressMapOfAddressArrayValueCall {
            func: ScView::new(HSC_NAME, HVIEW_ADDRESS_MAP_OF_ADDRESS_ARRAY_VALUE),
            params: MutableAddressMapOfAddressArrayValueParams { proxy: Proxy::nil() },
            results: ImmutableAddressMapOfAddressArrayValueResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn address_map_of_address_map_value(_ctx: &dyn ScViewCallContext) -> AddressMapOfAddressMapValueCall {
        let mut f = AddressMapOfAddressMapValueCall {
            func: ScView::new(HSC_NAME, HVIEW_ADDRESS_MAP_OF_ADDRESS_MAP_VALUE),
            params: MutableAddressMapOfAddressMapValueParams { proxy: Proxy::nil() },
            results: ImmutableAddressMapOfAddressMapValueResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn array_of_address_array_length(_ctx: &dyn ScViewCallContext) -> ArrayOfAddressArrayLengthCall {
        let mut f = ArrayOfAddressArrayLengthCall {
            func: ScView::new(HSC_NAME, HVIEW_ARRAY_OF_ADDRESS_ARRAY_LENGTH),
            results: ImmutableArrayOfAddressArrayLengthResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn array_of_address_array_value(_ctx: &dyn ScViewCallContext) -> ArrayOfAddressArrayValueCall {
        let mut f = ArrayOfAddressArrayValueCall {
            func: ScView::new(HSC_NAME, HVIEW_ARRAY_OF_ADDRESS_ARRAY_VALUE),
            params: MutableArrayOfAddressArrayValueParams { proxy: Proxy::nil() },
            results: ImmutableArrayOfAddressArrayValueResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn array_of_address_map_value(_ctx: &dyn ScViewCallContext) -> ArrayOfAddressMapValueCall {
        let mut f = ArrayOfAddressMapValueCall {
            func: ScView::new(HSC_NAME, HVIEW_ARRAY_OF_ADDRESS_MAP_VALUE),
            params: MutableArrayOfAddressMapValueParams { proxy: Proxy::nil() },
            results: ImmutableArrayOfAddressMapValueResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn array_of_string_array_length(_ctx: &dyn ScViewCallContext) -> ArrayOfStringArrayLengthCall {
        let mut f = ArrayOfStringArrayLengthCall {
            func: ScView::new(HSC_NAME, HVIEW_ARRAY_OF_STRING_ARRAY_LENGTH),
            results: ImmutableArrayOfStringArrayLengthResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn array_of_string_array_value(_ctx: &dyn ScViewCallContext) -> ArrayOfStringArrayValueCall {
        let mut f = ArrayOfStringArrayValueCall {
            func: ScView::new(HSC_NAME, HVIEW_ARRAY_OF_STRING_ARRAY_VALUE),
            params: MutableArrayOfStringArrayValueParams { proxy: Proxy::nil() },
            results: ImmutableArrayOfStringArrayValueResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn array_of_string_map_value(_ctx: &dyn ScViewCallContext) -> ArrayOfStringMapValueCall {
        let mut f = ArrayOfStringMapValueCall {
            func: ScView::new(HSC_NAME, HVIEW_ARRAY_OF_STRING_MAP_VALUE),
            params: MutableArrayOfStringMapValueParams { proxy: Proxy::nil() },
            results: ImmutableArrayOfStringMapValueResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn big_int_add(_ctx: &dyn ScViewCallContext) -> BigIntAddCall {
        let mut f = BigIntAddCall {
            func: ScView::new(HSC_NAME, HVIEW_BIG_INT_ADD),
            params: MutableBigIntAddParams { proxy: Proxy::nil() },
            results: ImmutableBigIntAddResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn big_int_div(_ctx: &dyn ScViewCallContext) -> BigIntDivCall {
        let mut f = BigIntDivCall {
            func: ScView::new(HSC_NAME, HVIEW_BIG_INT_DIV),
            params: MutableBigIntDivParams { proxy: Proxy::nil() },
            results: ImmutableBigIntDivResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn big_int_mod(_ctx: &dyn ScViewCallContext) -> BigIntModCall {
        let mut f = BigIntModCall {
            func: ScView::new(HSC_NAME, HVIEW_BIG_INT_MOD),
            params: MutableBigIntModParams { proxy: Proxy::nil() },
            results: ImmutableBigIntModResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn big_int_mul(_ctx: &dyn ScViewCallContext) -> BigIntMulCall {
        let mut f = BigIntMulCall {
            func: ScView::new(HSC_NAME, HVIEW_BIG_INT_MUL),
            params: MutableBigIntMulParams { proxy: Proxy::nil() },
            results: ImmutableBigIntMulResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn big_int_shl(_ctx: &dyn ScViewCallContext) -> BigIntShlCall {
        let mut f = BigIntShlCall {
            func: ScView::new(HSC_NAME, HVIEW_BIG_INT_SHL),
            params: MutableBigIntShlParams { proxy: Proxy::nil() },
            results: ImmutableBigIntShlResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn big_int_shr(_ctx: &dyn ScViewCallContext) -> BigIntShrCall {
        let mut f = BigIntShrCall {
            func: ScView::new(HSC_NAME, HVIEW_BIG_INT_SHR),
            params: MutableBigIntShrParams { proxy: Proxy::nil() },
            results: ImmutableBigIntShrResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn big_int_sub(_ctx: &dyn ScViewCallContext) -> BigIntSubCall {
        let mut f = BigIntSubCall {
            func: ScView::new(HSC_NAME, HVIEW_BIG_INT_SUB),
            params: MutableBigIntSubParams { proxy: Proxy::nil() },
            results: ImmutableBigIntSubResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn block_record(_ctx: &dyn ScViewCallContext) -> BlockRecordCall {
        let mut f = BlockRecordCall {
            func: ScView::new(HSC_NAME, HVIEW_BLOCK_RECORD),
            params: MutableBlockRecordParams { proxy: Proxy::nil() },
            results: ImmutableBlockRecordResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn block_records(_ctx: &dyn ScViewCallContext) -> BlockRecordsCall {
        let mut f = BlockRecordsCall {
            func: ScView::new(HSC_NAME, HVIEW_BLOCK_RECORDS),
            params: MutableBlockRecordsParams { proxy: Proxy::nil() },
            results: ImmutableBlockRecordsResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn check_agent_id(_ctx: &dyn ScViewCallContext) -> CheckAgentIDCall {
        let mut f = CheckAgentIDCall {
            func: ScView::new(HSC_NAME, HVIEW_CHECK_AGENT_ID),
            params: MutableCheckAgentIDParams { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn get_random(_ctx: &dyn ScViewCallContext) -> GetRandomCall {
        let mut f = GetRandomCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_RANDOM),
            results: ImmutableGetRandomResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn iota_balance(_ctx: &dyn ScViewCallContext) -> IotaBalanceCall {
        let mut f = IotaBalanceCall {
            func: ScView::new(HSC_NAME, HVIEW_IOTA_BALANCE),
            results: ImmutableIotaBalanceResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn string_map_of_string_array_length(_ctx: &dyn ScViewCallContext) -> StringMapOfStringArrayLengthCall {
        let mut f = StringMapOfStringArrayLengthCall {
            func: ScView::new(HSC_NAME, HVIEW_STRING_MAP_OF_STRING_ARRAY_LENGTH),
            params: MutableStringMapOfStringArrayLengthParams { proxy: Proxy::nil() },
            results: ImmutableStringMapOfStringArrayLengthResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn string_map_of_string_array_value(_ctx: &dyn ScViewCallContext) -> StringMapOfStringArrayValueCall {
        let mut f = StringMapOfStringArrayValueCall {
            func: ScView::new(HSC_NAME, HVIEW_STRING_MAP_OF_STRING_ARRAY_VALUE),
            params: MutableStringMapOfStringArrayValueParams { proxy: Proxy::nil() },
            results: ImmutableStringMapOfStringArrayValueResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn string_map_of_string_map_value(_ctx: &dyn ScViewCallContext) -> StringMapOfStringMapValueCall {
        let mut f = StringMapOfStringMapValueCall {
            func: ScView::new(HSC_NAME, HVIEW_STRING_MAP_OF_STRING_MAP_VALUE),
            params: MutableStringMapOfStringMapValueParams { proxy: Proxy::nil() },
            results: ImmutableStringMapOfStringMapValueResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }
}

module example::example;

use example::fund;

fun init(ctx: &mut TxContext) {
	fund::create_fund(ctx);
}

public fun ex01(_ctx: &mut TxContext) {
}

public fun ex02(_value : address, _ctx: &mut TxContext) {
}

public fun ex03(_value : bool, _ctx: &mut TxContext) {
}

public fun ex04(_value : u8, _ctx: &mut TxContext) {
}

public fun ex05(_value : u16, _ctx: &mut TxContext) {
}

public fun ex06(_value : u32, _ctx: &mut TxContext) {
}

public fun ex07(_value : u64, _ctx: &mut TxContext) {
}

public fun ex08(_value : u128, _ctx: &mut TxContext) {
}

public fun ex09(_value : u256, _ctx: &mut TxContext) {
}

public fun ex10(_value : address, _ctx: &mut TxContext) {
}

public fun ex11(_value : address, _ctx: &mut TxContext) {
}

public fun ex12(_value : vector<address>, _ctx: &mut TxContext) {
}

public fun ex13(_value : vector<bool>, _ctx: &mut TxContext) {
}

public fun ex14(_value : vector<u8>, _ctx: &mut TxContext) {
}

public fun ex15(_value : vector<u16>, _ctx: &mut TxContext) {
}

public fun ex16(_value : vector<u32>, _ctx: &mut TxContext) {
}

public fun ex17(_value : vector<u64>, _ctx: &mut TxContext) {
}

public fun ex18(_value : vector<u128>, _ctx: &mut TxContext) {
}

public fun ex19(_value : vector<u256>, _ctx: &mut TxContext) {
}

public fun ex20(_v1 : vector<u8>, _v2 : bool, _v3 :u32, _ctx: &mut TxContext) {
}

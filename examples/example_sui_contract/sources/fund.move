module example::fund;

use sui::{clock::Clock};

public struct Fund has key, store {
	id: UID,
	counter: u64,
}

public(package) fun create_fund(ctx: &mut TxContext) {
    let f = Fund {
        id: object::new(ctx),
        counter: 0,
    };
    transfer::share_object(f);
}

public fun ex2(f: &mut Fund, _ctx: &mut TxContext) {
    f.counter = f.counter + 1;
}

public fun ex3(f: &mut Fund, _ctx: &mut TxContext) {
    f.counter = f.counter + 1;
}

public fun ex4(f: &mut Fund, _clock: &Clock, _ctx: &mut TxContext) {
    f.counter = f.counter + 1;
}

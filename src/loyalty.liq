(* Auction contract from https://www.michelson-lang.com/contract-a-day.html *)

[%%version 0.15]

type loyalty = {
  customer : key_hash;
  loyalty_points : int;
}

let%init storage (customer : key_hash) =
  let loyalty.customer = customer in
  let loyalty.loyalty_points = 1000

let%entry main
    (parameter : unit)
    (storage : storage)
  : unit * storage =

  let loyalty.loyalty_points = (loyalty.loyalty_points + (Current.amount () * 0.1))
    
  ((), storage)

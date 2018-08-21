[
	{
		email: "wanlong@gmail.com"

		// available_balance + Locked balance
        total_balances: [
            jpy: 100
		 	usdt: 1000
        ]

		// (deposite - withdrawal) - Locked balance
		available_balances: [
		 	jpy: 100
		 	usdt: 1000
		]

		deposits: [
			{date, volume, currency}
		]

		withdrawal: [
			{date, volume, currency}
		]

		// when I place order or withdraw
		LockedBalances: [
			{created_at: xxx, currency: jpy, volume: 330, operation_type: withdraw|order, operation_id}
			{created_at: xxx, currency: usdt, volume: 10, operation_type: withdraw|order, operation_id}
		]

		// Situations:
		// 1. place buy|sell order
		// 2. buy order got partial|complete deal
		// 3. sell order got partial|complete deal
		//
        active_orders: [
			{
				id, created_at, done_at,
				direction: bid|ask (sell),
				type: market|limited,
				pair: XRP_JPY,
				volume: 10, // make sure this use has 3.3 * 10 amount of JPY
				unit_price: 3.3, // order will be deal when 3.3 bought
				deal_volume: 10, // snapshot
				// ??? average_deal_unit_price: 3.3
				trades: [
					{
						id, created_at, order_id 
						deal_volume: 9,
						pair_code: XRP_JPY,
						dealed_unit_price: 3.2,
						fee: 9 * 0.001 // unit XRP
					}, {
						id, created_at, order_id 
						deal_volume: 1,
						pair_code: XRP_JPY, 
						dealed_unit_price: 3.3,
						fee: 1 * 0.001 // unit XRP
					}
				]
			}, {
				id, created_at, done_at
				direction: bid|ask (buy),
				type: market|limited
				pair: XRP_JPY,
				volume: 10, // make sure this use has 3.3 * 10 amount of JPY
				unit_price: 3.3, // order will be deal when 3.3 bought
				deal_volume: 10, // snapshot
				// ??? average_deal_unit_price: 3.3
				trades: [
					{
						id, created_at, order_id 
						deal_volume: 9,
						pair_code: XRP_JPY,
						dealed_unit_price: 3.2,
						fee: 9 * 0.001 // unit XRP
					}, {
						id, created_at, order_id 
						deal_volume: 1,
						pair_code: XRP_JPY, 
						dealed_unit_price: 3.3,
						fee: 1 * 0.001 // unit XRP
					}
				]
			}
        ]
	}
]
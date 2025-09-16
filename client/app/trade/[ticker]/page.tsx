"use client";

import React, { useState } from "react";
import Chart from "@/components/trade/Chart";
import Orderbook from "@/components/trade/Orderbook";
import TickerStats from "@/components/trade/TickerStats";
import TradingWindow from "@/components/trade/TradingWindow";
import UserStats from "@/components/trade/UserStats";
import { PlaceOrderResponse } from "@/types/order";

interface TradePageProps {
  params: Promise<{ ticker: string }>;
}

export default function Trade({ params }: TradePageProps) {
  const [ticker, setTicker] = React.useState<string>("");
  const [newOrder, setNewOrder] = useState<PlaceOrderResponse | undefined>(
    undefined,
  );

  // Resolve params
  React.useEffect(() => {
    params.then(({ ticker }) => {
      // Convert URL format BTC_USD to BTC/USD for internal use
      const formattedTicker = ticker.replace("_", "/");
      setTicker(formattedTicker);
    });
  }, [params]);

  if (!ticker) {
    return (
      <div className="bg-secondary h-full p-5 flex items-center justify-center">
        <div className="w-6 h-6 border-2 border-blue-400 border-t-transparent rounded-full animate-spin"></div>
      </div>
    );
  }

  return (
    <div className="bg-secondary h-full p-5">
      <div className="flex flex-row gap-2">
        <div className=" flex flex-col gap-3">
          <TickerStats ticker={ticker} />
          <div className="flex flex-row gap-3">
            <div className="w-[71%]">
              <Chart ticker={ticker} />
            </div>
            <div className="w-2/8">
              <Orderbook ticker={ticker} />
            </div>
          </div>
          <div className="">
            <UserStats ticker={ticker} newOrder={newOrder} />
          </div>
        </div>

        <div className="w-3/8">
          <TradingWindow ticker={ticker} onOrderPlaced={setNewOrder} />
        </div>
      </div>
    </div>
  );
}

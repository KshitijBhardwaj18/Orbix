"use client";
import React from "react";

const data = {
  ticker: "BTC/USD",
  symbol: "/symbols/btc",
  lastPrice: "111,934.8",
  change24hpoints: "+941.3",
  change24hpercentage: "+0.90%",
  volume24h: "7,515,187.33",
  high24h: "112,922",
  low24h: "110,922.6",
};

interface TickerStatsProps {
  ticker: string;
}

const TickerStats = ({ ticker }: { ticker: string }) => {
  return (
    <div className="bg-primary  w-full rounded-xl shadow-2xl">
      <div className="flex flex-row items-center justify-start gap-8 p-2">
        <div className="flex items-center justify-center rounded-2xl bg-neutral-800 p-2 text-white text-sm gap-2">
          <img alt="btc" src="/symbols/btc.webp" className="size-6" />
          <p>
            <span>{data.ticker.split("/")[0]}</span>
            <span className="text-neutral-400">
              /{data.ticker.split("/")[1]}
            </span>
          </p>
        </div>

        <div className="flex flex-col items-center justify-center  text-xl">
          <span className="text-red-600  ">{data.lastPrice}</span>
          <span className="text-white text-lg">${data.lastPrice}</span>
        </div>

        <div className="flex flex-col items-center justify-center ">
          <span className="text-neutral-400 text-sm">24H Change</span>
          <div className="text-green-400 flex flex-row gap-1 ml-5 text-sm">
            <span>{data.change24hpercentage}</span>
            <span>{data.change24hpoints}</span>
          </div>
        </div>

        <div className="flex flex-col items-center justify-center ">
          <span className="text-neutral-400 text-sm">24H High</span>
          <div className="text-white text-sm flex flex-row gap-1 ">
            <span>{data.high24h}</span>
          </div>
        </div>

        <div className="flex flex-col items-center justify-center ">
          <span className="text-neutral-400 text-sm">24H Low</span>
          <div className="text-white text-sm flex flex-row gap-1 ">
            <span>{data.low24h}</span>
          </div>
        </div>

        <div className="flex flex-col  justify-center ">
          <span className="text-neutral-400 text-sm">24H Volume(USD)</span>
          <div className="text-white text-sm flex flex-row gap-1 text-left ">
            <span>{data.volume24h}</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default TickerStats;

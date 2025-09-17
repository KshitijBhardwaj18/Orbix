"use client";

import { useEffect, useState } from "react";
import { BidTable } from "./BidTable";
import { AskTable } from "./AskTable";
import { getDepth } from "@/utils/http";

interface DepthProps {
  market: string;
  precision?: number;
}

export function Depth({ market, precision = 0.00001 }: DepthProps) {
  const [bids, setBids] = useState<[string, string][]>([]);
  const [asks, setAsks] = useState<[string, string][]>([]);
  const [currentPrice, setCurrentPrice] = useState<string>("4296.38");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchDepthData = async () => {
      try {
        setLoading(true);
        const marketParam = market.replace("/", "_");
        const d = await getDepth(marketParam);

        console.log("🔍 Depth data received:", d);
        console.log("📊 Bids:", d.bids?.length || 0, "items");
        console.log("📈 Asks:", d.asks?.length || 0, "items");

        // Set the data
        setBids(d.bids || []);
        setAsks(d.asks || []);

        // Calculate current price from bid-ask spread
        if (d.bids?.length && d.asks?.length) {
          const bestBid = parseFloat(d.bids[0][0]);
          const bestAsk = parseFloat(d.asks[0][0]);
          const midPrice = ((bestBid + bestAsk) / 2).toFixed(5); // Show 5 decimal places
          setCurrentPrice(midPrice);
          console.log("💰 Current price set to:", midPrice);
        }
      } catch (error) {
        console.error("❌ Failed to fetch depth data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchDepthData();

    // Poll every 5 seconds for live updates
    // const interval = setInterval(fetchDepthData, 5000);
    // return () => clearInterval(interval);
  }, [market]);

  if (loading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="w-6 h-6 border-2 border-blue-400 border-t-transparent rounded-full animate-spin"></div>
        <span className="ml-2 text-gray-400">Loading orderbook...</span>
      </div>
    );
  }

  return (
    <div className="h-full flex flex-col">
      {/* Header */}
      <div className="px-4 py-2 border-b border-gray-800/30 flex-shrink-0">
        <div className="flex justify-between text-xs font-medium">
          <div className="text-gray-400">Price (USD)</div>
          <div className="text-gray-400">Size (SWTCH)</div>
          <div className="text-gray-400">Total (SWTCH)</div>
        </div>
      </div>

      {/* Content - Using precise height calculations */}
      <div className="flex-1 flex flex-col min-h-0">
        {/* Asks (Sell orders) - Top half - Scrolled to bottom initially */}
        <div className="flex-1 min-h-0 max-h-[45%] overflow-hidden">
          {asks.length > 0 ? (
            <AskTable asks={asks} precision={precision} />
          ) : (
            <div className="p-4 text-center text-gray-500 text-sm">
              No asks available
            </div>
          )}
        </div>

        {/* Current Price - Fixed in exact middle */}
        <div className="h-12 px-4 border-y border-gray-800/30 bg-gray-900/50 flex-shrink-0 flex items-center justify-center">
          <span className="text-xl font-bold text-red-400 font-mono">
            ${currentPrice}
          </span>
        </div>

        {/* Bids (Buy orders) - Bottom half with enforced scrolling */}
        <div className="flex-1 min-h-0 max-h-[45%] overflow-hidden">
          <div className="h-full overflow-y-auto custom-scrollbar">
            {bids.length > 0 ? (
              <BidTable bids={bids} precision={precision} />
            ) : (
              <div className="p-4 text-center text-gray-500 text-sm">
                No bids available
              </div>
            )}
          </div>
        </div>

        {/* Volume Distribution */}
        <div className="px-4 py-2 border-t border-gray-800/30 flex-shrink-0 h-10">
          <div className="flex h-6">
            <div className="flex-1 bg-green-500/20 rounded-l flex items-center justify-center">
              <span className="text-green-400 text-xs font-bold">62%</span>
            </div>
            <div className="flex-1 bg-red-500/20 rounded-r flex items-center justify-center">
              <span className="text-red-400 text-xs font-bold">38%</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

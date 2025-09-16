"use client";

import React, { useState } from "react";
import { Depth } from "../depth/Depth";

interface OrderbookProps {
  ticker: string;
}

const Orderbook = ({ ticker }: OrderbookProps) => {
  const [activeTab, setActiveTab] = useState<"book" | "trades">("book");
  const [precision, setPrecision] = useState(0.00001);

  return (
    <div className="h-[600px] bg-primary w-full rounded-xl overflow-hidden">
      {/* Header with tabs */}
      <div className="p-4 border-b border-gray-800/50">
        <div className="flex items-center justify-between mb-4">
          {/* Tabs */}
          <div className="flex">
            <button
              className={`px-4 py-2 text-sm font-medium rounded-lg transition-colors ${
                activeTab === "book"
                  ? "bg-gray-800 text-white"
                  : "text-gray-400 hover:text-gray-300"
              }`}
              onClick={() => setActiveTab("book")}
            >
              Book
            </button>
            <button
              className={`px-4 py-2 text-sm font-medium rounded-lg transition-colors ml-2 ${
                activeTab === "trades"
                  ? "bg-gray-800 text-white"
                  : "text-gray-400 hover:text-gray-300"
              }`}
              onClick={() => setActiveTab("trades")}
            >
              Trades
            </button>
          </div>

          {/* Precision controls */}
          {activeTab === "book" && (
            <div className="flex items-center space-x-1">
              <button
                className="w-6 h-6 flex items-center justify-center text-gray-400 hover:text-white transition-colors"
                onClick={() => setPrecision(Math.max(precision / 10, 0.00001))}
              >
                −
              </button>
              <span className="text-white text-sm font-mono px-2">
                {precision.toFixed(5)}
              </span>
              <button
                className="w-6 h-6 flex items-center justify-center text-gray-400 hover:text-white transition-colors"
                onClick={() => setPrecision(Math.min(precision * 10, 1))}
              >
                +
              </button>
            </div>
          )}
        </div>
      </div>

      {/* Content */}
      <div className="h-[calc(100%-80px)]">
        {activeTab === "book" ? (
          <Depth market={ticker} precision={precision} />
        ) : (
          <div className="p-4 text-gray-400 text-center">
            Trades component coming soon...
          </div>
        )}
      </div>
    </div>
  );
};

export default Orderbook;

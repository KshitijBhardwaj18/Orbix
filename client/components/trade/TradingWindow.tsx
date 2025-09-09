"use client";

import React from "react";
import { useState } from "react";
import { cn } from "@/lib/cn";

interface TradeProps {
  tricker: string;
}

const Trade = ({ ticker }: { ticker: string }) => {
  const [orderType, setOrderType] = useState("");
  const [price, setPrice] = useState("");
  const [quantity, setQuantity] = useState("");

  return (
    <div className="w-full  bg-primary rounded-xl p-4">
      <div className="flex flex-col gap-2">
        <div className="flex flex-row bg-neutral-7 rounded-xl bg-[#202126] font-bold text-neutral-400">
          <button
            className={cn(
              "flex items-center justify-center p-4  rounded-xl   flex-1 hover:text-green-400",
              orderType === "buy" && "bg-[#232d2c] text-green-400",
            )}
            onClick={() => setOrderType("buy")}
          >
            Buy
          </button>
          <button
            className={cn(
              orderType === "sell" && "bg-[#39242b] text-red-600  ",
              "flex items-center justify-center p-4 rounded-xl flex-1 hover:text-red-400",
            )}
            onClick={() => setOrderType("sell")}
          >
            {" "}
            Sell
          </button>
        </div>
      </div>
    </div>
  );
};

export default Trade;

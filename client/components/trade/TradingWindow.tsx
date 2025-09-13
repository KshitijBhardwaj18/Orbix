"use client";

import React from "react";
import { useState } from "react";
import { cn } from "@/lib/cn";

interface TradeProps {
  tricker: string;
}

const Trade = ({ ticker }: { ticker: string }) => {
  const [orderType, setOrderType] = useState<string>("BUY");
  const [price, setPrice] = useState<number>(0);
  const [quantity, setQuantity] = useState<number>(0);
  const [orderValue, setOrderValue] = useState<number>(0);

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    console.log({ orderType, price, quantity });
  };

  return (
    <div className="w-full  bg-primary rounded-xl p-4">
      <div className="flex flex-col gap-4 text-neutral-400">
        <div className="flex flex-row bg-neutral-7 rounded-xl bg-[#202126] font-bold ">
          <button
            className={cn(
              "flex items-center justify-center p-4  rounded-xl   flex-1 hover:text-green-400",
              orderType === "BUY" && "bg-[#232d2c] text-green-400",
            )}
            onClick={() => setOrderType("BUY")}
          >
            Buy
          </button>
          <button
            className={cn(
              orderType === "SELL" && "bg-[#39242b] text-red-600  ",
              "flex items-center justify-center p-4 rounded-xl flex-1 hover:text-red-400",
            )}
            onClick={() => setOrderType("SELL")}
          >
            {" "}
            Sell
          </button>
        </div>

        <div className="flex flex-col ">
          <form className="flex flex-col gap-6" onSubmit={onSubmit}>
            <div className="flex flex-col  justify-start gap-2 ">
              <label className="text-sm">Price</label>

              <div className="relative flex bg-[#202126] flex-1  rounded-xl ">
                <img
                  className="size-7 absolute right-4 top-2"
                  alt="usd"
                  src={"/symbols/usd.svg"}
                />
                <input
                  type="number"
                  className="p-3 w-full pr-12 [appearance:textfield] 
                  [&::-webkit-outer-spin-button]:appearance-none 
                  [&::-webkit-inner-spin-button]:appearance-none rounded-xl"
                  placeholder="0"
                  value={price}
                  onChange={(e) => setPrice(Number(e.target.value))}
                ></input>
              </div>
            </div>

            <div className="flex flex-col  justify-start gap-2 ">
              <label className="text-sm">Quantity</label>

              <div className="relative flex bg-[#202126] flex-1  rounded-xl ">
                <img
                  className="size-7 absolute right-4 top-2"
                  alt="usd"
                  src={"/symbols/btc.webp"}
                />
                <input
                  type="number"
                  className="p-3 w-full pr-12 [appearance:textfield] 
                  [&::-webkit-outer-spin-button]:appearance-none 
                  [&::-webkit-inner-spin-button]:appearance-none rounded-xl"
                  placeholder="0"
                  value={quantity}
                  onChange={(e) => setQuantity(Number(e.target.value))}
                ></input>
              </div>
            </div>

            <div className="flex flex-col  justify-start gap-2 ">
              <label className="text-sm">Order Value</label>

              <div className="relative flex bg-[#202126] flex-1  rounded-xl ">
                <img
                  className="size-7 absolute right-4 top-2"
                  alt="usd"
                  src={"/symbols/usd.svg"}
                />
                <input
                  type="number"
                  className="p-3 w-full pr-12 [appearance:textfield] 
                  [&::-webkit-outer-spin-button]:appearance-none 
                  [&::-webkit-inner-spin-button]:appearance-none rounded-xl"
                  placeholder="0"
                  value={orderValue}
                  onChange={(e) => setOrderValue(Number(e.target.value))}
                ></input>
              </div>
            </div>

            <div className="flex flex-1 items-center justify-center">
              <button
                className={cn(
                  "cursor-pointer flex items-center justify-center flex-1 bg-white text-black py-2  shadow-xl rounded-xl",
                  orderType == "BUY" && "hover:bg-green-400 hover:text-white",
                  orderType == "SELL" && "hover:bg-red-500 hover:text-white",
                )}
              >
                Place Order
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Trade;

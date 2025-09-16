"use client";

import React, { useState } from "react";
import OrderTabs from "./OrderTabs";
import { PlaceOrderResponse } from "@/types/order";

interface UserStatsProps {
  ticker: string;
  newOrder?: PlaceOrderResponse;
}

const UserStats: React.FC<UserStatsProps> = ({ ticker, newOrder }) => {
  return (
    <div className="w-full h-full">
      <OrderTabs ticker={ticker} newOrder={newOrder} />
    </div>
  );
};

export default UserStats;

"use client";

import React, { useState, useEffect } from "react";
import { api } from "@/lib/axios";
import { PlaceOrderResponse } from "@/types/order";

interface OrderTabsProps {
  ticker: string;
  newOrder?: PlaceOrderResponse; // Passed from TradingWindow when new order is placed
}

type TabType = "open" | "history";

const OrderTabs: React.FC<OrderTabsProps> = ({ ticker, newOrder }) => {
  const [activeTab, setActiveTab] = useState<TabType>("open");
  const [openOrders, setOpenOrders] = useState<PlaceOrderResponse[]>([]);
  const [orderHistory, setOrderHistory] = useState<PlaceOrderResponse[]>([]);
  const [loading, setLoading] = useState(false);

  // Fetch open orders
  const fetchOpenOrders = async () => {
    setLoading(true);
    try {
      const response = await api.get(`/orders/open?market=${ticker}`);
      setOpenOrders(response.data.orders || []);
    } catch (error) {
      console.error("Failed to fetch open orders:", error);
    } finally {
      setLoading(false);
    }
  };

  // Initial load
  useEffect(() => {
    fetchOpenOrders();
  }, [ticker]);

  // Add new order to open orders when order is placed
  useEffect(() => {
    if (newOrder) {
      setOpenOrders((prev) => [newOrder, ...prev]);
    }
  }, [newOrder]);

  const formatTime = (dateString: string) => {
    return new Date(dateString).toLocaleTimeString("en-US", {
      hour12: false,
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
    });
  };

  const formatPrice = (price?: string) => {
    if (!price) return "Market";
    return `$${parseFloat(price).toLocaleString("en-US", {
      minimumFractionDigits: 2,
      maximumFractionDigits: 8,
    })}`;
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case "FILLED":
        return "text-green-400 bg-green-400/10";
      case "PARTIAL":
        return "text-yellow-400 bg-yellow-400/10";
      case "PENDING":
        return "text-blue-400 bg-blue-400/10";
      case "CANCELLED":
        return "text-gray-400 bg-gray-400/10";
      case "REJECTED":
        return "text-red-400 bg-red-400/10";
      default:
        return "text-gray-400 bg-gray-400/10";
    }
  };

  const getSideColor = (side: string) => {
    return side === "BUY" ? "text-green-400" : "text-red-400";
  };

  const renderOrderList = (orders: PlaceOrderResponse[]) => {
    if (loading) {
      return (
        <div className="flex justify-center items-center py-8">
          <div className="w-6 h-6 border-2 border-blue-400 border-t-transparent rounded-full animate-spin"></div>
        </div>
      );
    }

    if (orders.length === 0) {
      return (
        <div className="text-center text-gray-400 py-8">
          <div className="text-2xl mb-2">📊</div>
          <p>No {activeTab} orders</p>
          <p className="text-sm">
            {activeTab === "open"
              ? "Your active orders will appear here"
              : "Your order history will appear here"}
          </p>
        </div>
      );
    }

    return (
      <div className="space-y-2">
        {orders.map((order, index) => (
          <div
            key={order.id || index}
            className="bg-[#202126] rounded-lg p-3 border border-gray-800/50 hover:border-gray-700/50 transition-colors"
          >
            {/* Header Row */}
            <div className="flex justify-between items-start mb-2">
              <div className="flex items-center gap-2">
                <span className={`font-medium ${getSideColor(order.side)}`}>
                  {order.side}
                </span>
                <span className="text-white font-medium text-sm">
                  {order.market_id}
                </span>
              </div>
              <div className="flex items-center gap-2">
                <span
                  className={`px-2 py-1 rounded text-xs font-medium ${getStatusColor(
                    order.status,
                  )}`}
                >
                  {order.status}
                </span>
                <span className="text-gray-400 text-xs">
                  {formatTime(order.created_at)}
                </span>
              </div>
            </div>

            {/* Order Details Grid */}
            <div className="grid grid-cols-2 gap-3 text-sm">
              <div>
                <span className="text-gray-400">Price:</span>
                <span className="text-white ml-2">
                  {formatPrice(order.price)}
                </span>
              </div>
              <div>
                <span className="text-gray-400">Type:</span>
                <span className="text-white ml-2">{order.type}</span>
              </div>
              <div>
                <span className="text-gray-400">Quantity:</span>
                <span className="text-white ml-2">{order.quantity}</span>
              </div>
              <div>
                <span className="text-gray-400">Filled:</span>
                <span className="text-white ml-2">{order.filledQuantity}</span>
              </div>
            </div>

            {/* Progress Bar for Partial Fills */}
            {order.status === "PARTIAL" && (
              <div className="mt-3">
                <div className="flex justify-between text-xs text-gray-400 mb-1">
                  <span>Fill Progress</span>
                  <span>
                    {(
                      (parseFloat(order.filledQuantity) /
                        parseFloat(order.quantity)) *
                      100
                    ).toFixed(1)}
                    %
                  </span>
                </div>
                <div className="w-full bg-gray-700 rounded-full h-2">
                  <div
                    className="bg-yellow-400 h-2 rounded-full transition-all duration-300"
                    style={{
                      width: `${
                        (parseFloat(order.filledQuantity) /
                          parseFloat(order.quantity)) *
                        100
                      }%`,
                    }}
                  ></div>
                </div>
              </div>
            )}

            {/* Action buttons for open orders */}
            {activeTab === "open" && order.status === "PENDING" && (
              <div className="mt-3 flex gap-2">
                <button className="flex-1 bg-red-600/20 text-red-400 py-1 px-3 rounded text-sm hover:bg-red-600/30 transition-colors">
                  Cancel
                </button>
                <button className="flex-1 bg-blue-600/20 text-blue-400 py-1 px-3 rounded text-sm hover:bg-blue-600/30 transition-colors">
                  Modify
                </button>
              </div>
            )}
          </div>
        ))}
      </div>
    );
  };

  return (
    <div className="w-full bg-primary rounded-xl p-4">
      {/* Tab Headers */}
      <div className="flex border-b border-gray-800 mb-4">
        <button
          className={`flex-1 py-2 px-4 text-sm font-medium transition-colors ${
            activeTab === "open"
              ? "text-blue-400 border-b-2 border-blue-400"
              : "text-gray-400 hover:text-gray-300"
          }`}
          onClick={() => setActiveTab("open")}
        >
          Open Orders ({openOrders.length})
        </button>
        <button
          className={`flex-1 py-2 px-4 text-sm font-medium transition-colors ${
            activeTab === "history"
              ? "text-blue-400 border-b-2 border-blue-400"
              : "text-gray-400 hover:text-gray-300"
          }`}
          onClick={() => setActiveTab("history")}
        >
          Order History ({orderHistory.length})
        </button>
      </div>

      {/* Refresh Button */}
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-white font-medium">
          {activeTab === "open" ? "Open Orders" : "Order History"}
        </h3>
        <button
          onClick={fetchOpenOrders}
          className="text-gray-400 hover:text-white transition-colors"
          disabled={loading}
        >
          <svg
            className="w-4 h-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
            />
          </svg>
        </button>
      </div>

      {/* Tab Content */}
      <div className="max-h-96 overflow-y-auto">
        {activeTab === "open"
          ? renderOrderList(openOrders)
          : renderOrderList(orderHistory)}
      </div>
    </div>
  );
};

export default OrderTabs;

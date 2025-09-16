"use client";

import React, { useState, useEffect, useCallback } from "react";
import { api } from "@/lib/axios";
import { PlaceOrderResponse } from "@/types/order";
import { cn } from "@/lib/cn";

interface OrderTabsProps {
  ticker: string;
  newOrder?: PlaceOrderResponse;
}

type TabType = "open" | "history";

const OrderTabs: React.FC<OrderTabsProps> = ({ ticker, newOrder }) => {
  const [activeTab, setActiveTab] = useState<TabType>("open");
  const [openOrders, setOpenOrders] = useState<PlaceOrderResponse[]>([]);
  const [orderHistory, setOrderHistory] = useState<PlaceOrderResponse[]>([]);
  const [loading, setLoading] = useState(false);
  const [isPolling, setIsPolling] = useState(true);

  // Fetch open orders
  const fetchOpenOrders = useCallback(async () => {
    if (!isPolling) return;

    setLoading(true);
    try {
      // Convert ticker format for API call
      const marketParam = ticker.replace("/", "_");
      const response = await api.get(`/orders/open?market=${marketParam}`);
      setOpenOrders(response.data.orders || []);
    } catch (error) {
      console.error("Failed to fetch open orders:", error);
    } finally {
      setLoading(false);
    }
  }, [ticker, isPolling]);

  // Initial load
  useEffect(() => {
    fetchOpenOrders();
  }, [fetchOpenOrders]);

  // Polling effect - every 10 seconds
  useEffect(() => {
    if (!isPolling) return;

    const interval = setInterval(() => {
      fetchOpenOrders();
    }, 10000); // Poll every 10 seconds

    return () => clearInterval(interval);
  }, [fetchOpenOrders, isPolling]);

  // Add new order to open orders when order is placed
  useEffect(() => {
    if (newOrder) {
      setOpenOrders((prev) => [newOrder, ...prev]);
    }
  }, [newOrder]);

  // Stop polling when component unmounts
  useEffect(() => {
    return () => setIsPolling(false);
  }, []);

  const formatTime = (dateString: string) => {
    return new Date(dateString).toLocaleTimeString("en-US", {
      hour12: false,
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
    });
  };

  const formatPrice = (price?: string) => {
    if (!price) return "Market";
    return parseFloat(price).toLocaleString("en-US", {
      minimumFractionDigits: 2,
      maximumFractionDigits: 8,
    });
  };

  const formatQuantity = (quantity: string) => {
    return parseFloat(quantity).toLocaleString("en-US", {
      minimumFractionDigits: 0,
      maximumFractionDigits: 8,
    });
  };

  const getStatusBadge = (status: string) => {
    const baseClasses = "px-2 py-1 rounded-full text-xs font-medium";
    switch (status) {
      case "FILLED":
        return `${baseClasses} text-green-400 bg-green-400/10`;
      case "PARTIAL":
        return `${baseClasses} text-yellow-400 bg-yellow-400/10`;
      case "PENDING":
        return `${baseClasses} text-blue-400 bg-blue-400/10`;
      case "CANCELLED":
        return `${baseClasses} text-gray-400 bg-gray-400/10`;
      case "REJECTED":
        return `${baseClasses} text-red-400 bg-red-400/10`;
      default:
        return `${baseClasses} text-gray-400 bg-gray-400/10`;
    }
  };

  const getSideColor = (side: string) => {
    return side === "BUY" ? "text-green-400" : "text-red-400";
  };

  const renderTableHeader = () => (
    <div className="grid grid-cols-8 gap-4 p-3 border-b border-gray-800/50 text-gray-400 text-sm font-medium">
      <div>Time</div>
      <div>Side</div>
      <div>Price</div>
      <div>Quantity</div>
      <div>Filled</div>
      <div>Remaining</div>
      <div>Status</div>
      <div>Actions</div>
    </div>
  );

  const renderOrderRow = (order: PlaceOrderResponse, index: number) => {
    const fillPercentage =
      (parseFloat(order.filledQuantity) / parseFloat(order.quantity)) * 100;

    return (
      <div
        key={order.id || index}
        className="grid grid-cols-8 gap-4 p-3 border-b border-gray-800/20 hover:bg-gray-800/20 transition-colors text-sm"
      >
        {/* Time */}
        <div className="text-gray-400">
          <div>{formatTime(order.created_at)}</div>
          <div className="text-xs">{formatDate(order.created_at)}</div>
        </div>

        {/* Side */}
        <div className={`font-medium ${getSideColor(order.side)}`}>
          {order.side}
        </div>

        {/* Price */}
        <div className="text-white">{formatPrice(order.price)}</div>

        {/* Quantity */}
        <div className="text-white">{formatQuantity(order.quantity)}</div>

        {/* Filled */}
        <div className="text-white">
          {formatQuantity(order.filledQuantity)}
          {order.status === "PARTIAL" && (
            <div className="text-xs text-gray-400">
              {fillPercentage.toFixed(1)}%
            </div>
          )}
        </div>

        {/* Remaining */}
        <div className="text-white">
          {formatQuantity(order.remaining_quantity)}
        </div>

        {/* Status */}
        <div>
          <span className={getStatusBadge(order.status)}>{order.status}</span>
          {order.status === "PARTIAL" && (
            <div className="w-full bg-gray-700 rounded-full h-1 mt-1">
              <div
                className="bg-yellow-400 h-1 rounded-full transition-all duration-300"
                style={{ width: `${fillPercentage}%` }}
              />
            </div>
          )}
        </div>

        {/* Actions */}
        <div className="flex gap-1">
          {activeTab === "open" && order.status === "PENDING" && (
            <>
              <button className="px-2 py-1 bg-red-600/20 text-red-400 rounded text-xs hover:bg-red-600/30 transition-colors">
                Cancel
              </button>
              <button className="px-2 py-1 bg-blue-600/20 text-blue-400 rounded text-xs hover:bg-blue-600/30 transition-colors">
                Edit
              </button>
            </>
          )}
          {activeTab === "history" && (
            <button className="px-2 py-1 bg-gray-600/20 text-gray-400 rounded text-xs hover:bg-gray-600/30 transition-colors">
              Details
            </button>
          )}
        </div>
      </div>
    );
  };

  const renderOrderTable = (orders: PlaceOrderResponse[]) => {
    if (loading && orders.length === 0) {
      return (
        <div className="flex justify-center items-center py-12">
          <div className="w-6 h-6 border-2 border-blue-400 border-t-transparent rounded-full animate-spin"></div>
          <span className="ml-2 text-gray-400">Loading orders...</span>
        </div>
      );
    }

    if (orders.length === 0) {
      return (
        <div className="text-center text-gray-400 py-12">
          <div className="text-4xl mb-3">📊</div>
          <p className="text-lg font-medium mb-1">No {activeTab} orders</p>
          <p className="text-sm">
            {activeTab === "open"
              ? "Your active orders will appear here"
              : "Your completed orders will appear here"}
          </p>
        </div>
      );
    }

    return (
      <div className="bg-[#1a1a1a] rounded-lg overflow-hidden">
        {renderTableHeader()}
        <div className="max-h-80 overflow-y-auto custom-scrollbar">
          {orders.map((order, index) => renderOrderRow(order, index))}
        </div>
      </div>
    );
  };

  const handleRefresh = () => {
    fetchOpenOrders();
  };

  return (
    <div className="w-full bg-primary rounded-xl p-4">
      {/* Tab Headers */}
      <div className="flex border-b border-gray-800 mb-4">
        <button
          className={cn(
            "flex-1 py-3 px-4 text-sm font-medium transition-all duration-200 relative",
            activeTab === "open"
              ? "text-blue-400"
              : "text-gray-400 hover:text-gray-300",
          )}
          onClick={() => setActiveTab("open")}
        >
          Open Orders
          {openOrders.length > 0 && (
            <span className="ml-2 px-2 py-0.5 bg-blue-500/20 text-blue-400 rounded-full text-xs">
              {openOrders.length}
            </span>
          )}
          {activeTab === "open" && (
            <div className="absolute bottom-0 left-0 right-0 h-0.5 bg-blue-400 rounded-t-sm" />
          )}
        </button>
        <button
          className={cn(
            "flex-1 py-3 px-4 text-sm font-medium transition-all duration-200 relative",
            activeTab === "history"
              ? "text-blue-400"
              : "text-gray-400 hover:text-gray-300",
          )}
          onClick={() => setActiveTab("history")}
        >
          Order History
          {orderHistory.length > 0 && (
            <span className="ml-2 px-2 py-0.5 bg-blue-500/20 text-blue-400 rounded-full text-xs">
              {orderHistory.length}
            </span>
          )}
          {activeTab === "history" && (
            <div className="absolute bottom-0 left-0 right-0 h-0.5 bg-blue-400 rounded-t-sm" />
          )}
        </button>
      </div>

      {/* Controls */}
      <div className="flex justify-between items-center mb-4">
        <div className="flex items-center gap-3">
          <h3 className="text-white font-medium">
            {activeTab === "open" ? "Open Orders" : "Order History"}
          </h3>
          {isPolling && activeTab === "open" && (
            <div className="flex items-center text-green-400 text-xs">
              <div className="w-2 h-2 bg-green-400 rounded-full animate-pulse mr-1"></div>
              Live
            </div>
          )}
        </div>
        <div className="flex items-center gap-2">
          <button
            onClick={handleRefresh}
            disabled={loading}
            className={cn(
              "p-2 rounded-lg transition-colors",
              loading
                ? "text-gray-500 cursor-not-allowed"
                : "text-gray-400 hover:text-white hover:bg-gray-800/50",
            )}
          >
            <svg
              className={cn("w-4 h-4", loading && "animate-spin")}
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
          <button
            onClick={() => setIsPolling(!isPolling)}
            className={cn(
              "px-3 py-1 rounded-lg text-xs font-medium transition-colors",
              isPolling
                ? "bg-green-500/20 text-green-400 hover:bg-green-500/30"
                : "bg-gray-500/20 text-gray-400 hover:bg-gray-500/30",
            )}
          >
            {isPolling ? "Auto-refresh ON" : "Auto-refresh OFF"}
          </button>
        </div>
      </div>

      {/* Tab Content */}
      {activeTab === "open"
        ? renderOrderTable(openOrders)
        : renderOrderTable(orderHistory)}
    </div>
  );
};

export default OrderTabs;

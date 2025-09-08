// Create client/components/home/MarketCards.tsx
import React from "react";

const marketData = [
  // New tokens
  {
    category: "new",
    symbol: "WLFI",
    name: "World Liberty Financial",
    price: 0.21016,
    change: "-11.38%",
    icon: "/symbols/wlfi.webp",
  },
  {
    category: "new",
    symbol: "APT",
    name: "Aptos",
    price: 4.31,
    change: "+1.70%",
    icon: "/symbols/apt.webp",
  },
  {
    category: "new",
    symbol: "DOGE",
    name: "Dogecoin",
    price: 0.23229,
    change: "+6.85%",
    icon: "/symbols/doge.webp",
  },
  {
    category: "new",
    symbol: "SEI",
    name: "Sei",
    price: 0.29544,
    change: "+0.30%",
    icon: "/symbols/sei.webp",
  },
  {
    category: "new",
    symbol: "ES",
    name: "Ethena",
    price: 0.13792,
    change: "+0.19%",
    icon: "/symbols/ena.webp",
  },
  // Top Gainers
  {
    category: "gainers",
    symbol: "WLD",
    name: "Worldcoin",
    price: 1.27,
    change: "+24.19%",
    icon: "/symbols/wld.webp",
  },
  {
    category: "gainers",
    symbol: "PENGU",
    name: "Pudgy Penguins",
    price: 0.03111,
    change: "+7.28%",
    icon: "/symbols/pengu.webp",
  },
  {
    category: "gainers",
    symbol: "DOGE",
    name: "Dogecoin",
    price: 0.23229,
    change: "+6.85%",
    icon: "/symbols/doge.webp",
  },
  {
    category: "gainers",
    symbol: "BONK",
    name: "Bonk",
    price: 0.00002161,
    change: "+6.45%",
    icon: "/symbols/bonk.webp", 
  },
  {
    category: "gainers",
    symbol: "RENDER",
    name: "Render",
    price: 3.57,
    change: "+3.03%",
    icon: "/symbols/render.webp",
  },

  {
    category: "popular",
    symbol: "SOL",
    name: "Solana",
    price: 207.96,
    change: "+2.44%",
    icon: "/symbols/sol.webp",
  },
  {
    category: "popular",
    symbol: "ETH",
    name: "Ethereum",
    price: 4290.31,
    change: "-0.13%",
    icon: "/symbols/eth.webp",
  },
  {
    category: "popular",
    symbol: "BTC",
    name: "Bitcoin",
    price: 111281.3,
    change: "+0.47%",
    icon: "/symbols/btc.webp",
  },
  {
    category: "popular",
    symbol: "SUI",
    name: "Sui",
    price: 3.41,
    change: "+1.11%",
    icon: "/symbols/sui.webp",
  },
  {
    category: "popular",
    symbol: "USDT",
    name: "Tether",
    price: 1.0,
    change: "-0.01%",
    icon: "/symbols/usdt.webp",
  },
];

function MarketTrends() {
  const formatPrice = (price: number) => {
    if (price >= 1000) {
      return `$${price.toLocaleString("en-US", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}`;
    } else if (price >= 1) {
      return `$${price.toFixed(2)}`;
    } else {
      return `$${price.toFixed(8)}`.replace(/\.?0+$/, "");
    }
  };

  const formatChange = (change: string) => {
    const isPositive = change.startsWith("+");
    const isNegative = change.startsWith("-");
    return {
      value: change,
      className: isPositive
        ? "text-green-400"
        : isNegative
          ? "text-red-400"
          : "text-gray-400",
    };
  };

  const categories = [
    {
      key: "new",
      title: "New",
      markets: marketData.filter((m) => m.category === "new"),
    },
    {
      key: "gainers",
      title: "Top Gainers",
      markets: marketData.filter((m) => m.category === "gainers"),
    },
    {
      key: "popular",
      title: "Popular",
      markets: marketData.filter((m) => m.category === "popular"),
    },
  ];

  return (
    <div className="mb-6 grid grid-cols-1 gap-3 lg:grid-cols-3">
      {categories.map((category) => (
        <div
          key={category.key}
          className="bg-primary rounded-2xl p-2 shadow-2xl"
        >
          <h3 className="mb-2 text-sm font-semibold text-white">
            {category.title}
          </h3>
          <div className="flex m-0 p-0 flex-col gap-0">
            {category.markets.map((market, index) => (
              <div
                key={index}
                className="flex cursor-pointer items-center justify-between rounded-lg p-2 transition-colors hover:bg-gray-800/20"
              >
                <div className="flex items-center gap-3">
                  <div className="flex h-6 w-6 items-center justify-center overflow-hidden rounded-full bg-gray-800">
                    <img
                      src={market.icon}
                      alt={market.name}
                      className="h-6 w-6 object-contain"
                    />
                    <div
                      className="h-6 w-6 items-center justify-center rounded-full bg-gray-600 text-xs font-medium text-white"
                      style={{ display: "none" }}
                    >
                      {market.symbol.charAt(0)}
                    </div>
                  </div>
                  <div>
                    <div className="text-sm font-medium text-white">
                      {market.symbol}
                    </div>
                  </div>
                </div>
                <div className="flex flex-row gap-10">
                  <div className="font-mono text-sm text-white">
                    {formatPrice(market.price)}
                  </div>
                  <div
                    className={`text-xs ${formatChange(market.change).className}`}
                  >
                    {formatChange(market.change).value}
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
}

export default MarketTrends;

import React from 'react'

const data = [
    {
      name: "Bitcoin",
      symbol: "BTC/USD",
      price: 111227.7,
      "24h_volume": "4.3M",
      market_cap: "2.2T",
      "24h_change": "+0.60%",
      icon: "/symbols/btc.webp"
    },
    {
      name: "Ethereum",
      symbol: "ETH/USD",
      price: 4296.38,
      "24h_volume": "9.3M",
      market_cap: "518.1B",
      "24h_change": "+0.22%",
      icon: "/symbols/eth.webp"
    },
    {
      name: "USDT",
      symbol: "USDT/USD",
      price: 1.0001,
      "24h_volume": "908.8K",
      market_cap: "168.9B",
      "24h_change": "-0.01%",
      icon: "/symbols/usdt.webp"
    },
    {
      name: "Solana",
      symbol: "SOL/USD",
      price: 207.27,
      "24h_volume": "16.5M",
      market_cap: "112.3B",
      "24h_change": "+2.35%",
      icon: "/symbols/sol.webp"
    },
    {
      name: "Dogecoin",
      symbol: "DOGE/USD",
      price: 0.23229,
      "24h_volume": "52.4K",
      market_cap: "35B",
      "24h_change": "+7.01%",
      icon: "/symbols/doge.webp"
    },
    {
      name: "Chainlink",
      symbol: "LINK/USD",
      price: 22.427,
      "24h_volume": "28.4K",
      market_cap: "15.2B",
      "24h_change": "+0.68%",
      icon: "/symbols/link.svg"
    },
    {
      name: "Sui",
      symbol: "SUI/USD",
      price: 3.397,
      "24h_volume": "2.1M",
      market_cap: "12.1B",
      "24h_change": "+0.84%",
      icon: "/symbols/sui.webp"
    },
    {
      name: "Shiba Inu",
      symbol: "SHIB/USD",
      price: 0.00001255,
      "24h_volume": "37K",
      market_cap: "7.4B",
      "24h_change": "+1.60%",
      icon: "/symbols/shib.svg"
    },
    {
      name: "Render",
      symbol: "RENDER/USD",
      price: 3.572,
      "24h_volume": "$348",
      market_cap: "$1.9B",
      "24h_change": "+3.03%",
      icon: "/symbols/render.webp" // Fallback - you can add render.webp if you have it
    },
    {
      name: "Sei",
      symbol: "SEI/USD",
      price: 0.29544,
      "24h_volume": "$15.2K",
      market_cap: "$1.8B",
      "24h_change": "-0.04%",
      icon: "/symbols/sei.webp" // Fallback - you can add sei.webp if you have it
    },
    {
      name: "Ondo",
      symbol: "ONDO/USD",
      price: 0.9155,
      "24h_volume": "$5.7K",
      market_cap: "$2.9B",
      "24h_change": "+1.00%",
      icon: "/symbols/ondo.svg" // Fallback - you can add ondo.webp if you have it
    },
    {
      name: "Worldcoin",
      symbol: "WLD/USD",
      price: 1.2728,
      "24h_volume": "$10K",
      market_cap: "$2.5B",
      "24h_change": "+24.19%",
      icon: "/symbols/wld.webp" // Fallback - you can add wld.webp if you have it
    },
    {
      name: "Pudgy Penguins",
      symbol: "PENGU/USD",
      price: 0.03131,
      "24h_volume": "$53.9K",
      market_cap: "$2B",
      "24h_change": "+7.97%",
      icon: "/symbols/pengu.webp" // Fallback - you can add pengu.webp if you have it
    },
    {
      name: "Pepe",
      symbol: "PEPE/USD",
      price: 0.00001001,
      "24h_volume": "$2.1K",
      market_cap: "$4.2B",
      "24h_change": "+2.30%",
      icon: "/symbols/pepe.svg" // Using doge as similar meme coin
    },
    {
      name: "Aptos",
      symbol: "APT/USD",
      price: 4.322,
      "24h_volume": "$506.9K",
      market_cap: "$3B",
      "24h_change": "+1.91%",
      icon: "/symbols/apt.webp" // Fallback - you can add apt.webp if you have it
    },
    {
      name: "POL (ex-MATIC)",
      symbol: "POL/USD",
      price: 0.2779,
      "24h_volume": "$6.6K",
      market_cap: "$2.9B",
      "24h_change": "-1.21%",
      icon: "/symbols/pol.webp" // Fallback - you can add pol.webp if you have it
    },
    {
      name: "Uniswap",
      symbol: "UNI/USD",
      price: 9.40,
      "24h_volume": "$1K",
      market_cap: "$5.7B",
      "24h_change": "+0.43%",
      icon: "/symbols/uni.webp" // Fallback - you can add uni.webp if you have it
    },
    {
      name: "Ethena",
      symbol: "ENA/USD",
      price: 0.765,
      "24h_volume": "$41.6K",
      market_cap: "$5.3B",
      "24h_change": "+4.08%",
      icon: "/symbols/ena.webp" // Fallback - you can add ena.webp if you have it
    },
    {
      name: "Aave",
      symbol: "AAVE/USD",
      price: 299.59,
      "24h_volume": "$8.1K",
      market_cap: "$4.6B",
      "24h_change": "-0.24%",
      icon: "/symbols/aave.svg" // Fallback - you can add aave.webp if you have it
    },


];

function MarketTable() {
  const formatPrice = (price: number) => {
    return `$${price.toLocaleString('en-US', {
      minimumFractionDigits: price < 1 ? 8 : 2,
      maximumFractionDigits: price < 1 ? 8 : 2
    })}`;
  };

  const formatChange = (change: string) => {
    const isPositive = change.startsWith('+');
    const isNegative = change.startsWith('-');
    return {
      value: change,
      className: isPositive ? 'text-green-400' : isNegative ? 'text-red-400' : 'text-gray-400'
    };
  };

  return (
    <div className="w-full bg-primary shadow-2xl rounded-2xl p-3">
      <table className="w-full">
        <thead>
          <tr className="border-b border-gray-800">
            <th className="text-left py-3 px-4 text-gray-400 font-medium text-sm">Name</th>
            <th className="text-right py-3 px-4 text-gray-400 font-medium text-sm">Price</th>
            <th className="text-right py-3 px-4 text-gray-400 font-medium text-sm">24h Volume</th>
            <th className="text-right py-3 px-4 text-gray-400 font-medium text-sm">
              <div className="flex items-center justify-end gap-1">
                <span>↓</span>
                <span>Market Cap</span>
              </div>
            </th>
            <th className="text-right py-3 px-4 text-gray-400 font-medium text-sm">24h Change</th>
          </tr>
        </thead>
        <tbody>
          {data.map((crypto, index) => {
            const changeData = formatChange(crypto["24h_change"]);
            return (
              <tr 
                key={index} 
                className="border-b border-gray-800/50 hover:bg-gray-800/20 transition-colors cursor-pointer"
              >
                <td className="py-4 px-4">
                  <div className="flex items-center gap-3">
                    <div className="w-8 h-8 rounded-full overflow-hidden flex items-center justify-center bg-gray-800">
                      <img 
                        src={crypto.icon} 
                        alt={crypto.name}
                        className="w-8 h-8 object-contain"
                      
                      />
                      <div 
                        className="w-8 h-8 rounded-full bg-gray-600 items-center justify-center text-white text-xs font-medium hidden"
                        style={{ display: 'none' }}
                      >
                        {crypto.name.charAt(0)}
                      </div>
                    </div>
                    <div>
                      <div className="text-white font-medium">{crypto.name}</div>
                      <div className="text-gray-400 text-sm">{crypto.symbol}</div>
                    </div>
                  </div>
                </td>
                <td className="py-4 px-4 text-right text-white font-mono">
                  {formatPrice(crypto.price)}
                </td>
                <td className="py-4 px-4 text-right text-white">
                  {crypto["24h_volume"]}
                </td>
                <td className="py-4 px-4 text-right text-white">
                  {crypto.market_cap}
                </td>
                <td className="py-4 px-4 text-right">
                  <span className={changeData.className}>
                    {changeData.value}
                  </span>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}

export default MarketTable;
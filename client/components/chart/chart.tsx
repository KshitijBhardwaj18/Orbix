import { useEffect, useRef } from "react";
import { ChartManager } from "@/utils/chartmanager/ChartManger";

export interface KLine {
  close: string;
  end: string;
  high: string;
  low: string;
  open: string;
  quoteVolume: string;
  start: string;
  trades: string;
  volume: string;
}

// Hardcoded sample data for BTC/USD with realistic price movements
const generateSampleKlineData = (): KLine[] => {
  const basePrice = 42500; // Starting price around 42,500 USD
  const data: KLine[] = [];
  const now = new Date().getTime();
  const oneHour = 60 * 60 * 1000; // 1 hour in milliseconds

  // Generate 7 days of hourly data (168 hours)
  for (let i = 168; i >= 0; i--) {
    const timestamp = now - i * oneHour;
    const startTime = new Date(timestamp - oneHour).toISOString();
    const endTime = new Date(timestamp).toISOString();

    // Create realistic price movement
    const volatility = 100; // Price can move ±100 USD per hour
    const trend = Math.sin(i / 24) * 200; // Daily cycle with 200 USD amplitude
    const randomChange = (Math.random() - 0.5) * volatility;

    const open = basePrice + trend + randomChange;
    const close = open + (Math.random() - 0.5) * 50; // ±25 USD movement within hour
    const high = Math.max(open, close) + Math.random() * 30; // Up to 30 USD above
    const low = Math.min(open, close) - Math.random() * 30; // Up to 30 USD below

    // Random volume between 100-1000 BTC
    const volume = (100 + Math.random() * 900).toFixed(2);
    const quoteVolume = (parseFloat(volume) * ((high + low) / 2)).toFixed(2);

    data.push({
      open: open.toFixed(2),
      high: high.toFixed(2),
      low: low.toFixed(2),
      close: close.toFixed(2),
      volume: volume,
      quoteVolume: quoteVolume,
      start: startTime,
      end: endTime,
      trades: Math.floor(50 + Math.random() * 200).toString(), // 50-250 trades per hour
    });
  }

  return data;
};

export function TradeView({ market }: { market: string }) {
  const chartRef = useRef<HTMLDivElement>(null);
  const chartManagerRef = useRef<ChartManager>(null);

  useEffect(() => {
    const init = async () => {
      // Use hardcoded sample data instead of API call
      const klineData: KLine[] = generateSampleKlineData();

      console.log("📊 Using sample candlestick data for market:", market);
      console.log("📈 Generated", klineData.length, "data points");

      if (chartRef.current) {
        if (chartManagerRef.current) {
          chartManagerRef.current.destroy();
        }

        console.log("🕯️ Sample kline data:", klineData.slice(0, 5)); // Log first 5 entries

        const chartManager = new ChartManager(
          chartRef.current,
          [
            ...klineData?.map((x) => ({
              close: parseFloat(x.close),
              high: parseFloat(x.high),
              low: parseFloat(x.low),
              open: parseFloat(x.open),
              timestamp: new Date(x.end).getTime(), // Convert to timestamp
            })),
          ].sort((x, y) => (x.timestamp < y.timestamp ? -1 : 1)) || [],
          {
            background: "#0e0f14",
            color: "white",
          },
        );

        //@ts-ignore
        chartManagerRef.current = chartManager;
      }
    };

    init();
  }, [market, chartRef]);

  return (
    <div className="w-full h-full bg-primary rounded-xl p-4">
      <div className="mb-4">
        <h3 className="text-white text-lg font-semibold">{market} Chart</h3>
        <p className="text-gray-400 text-sm">
          1 Hour Candlesticks (Sample Data)
        </p>
      </div>
      <div
        ref={chartRef}
        style={{
          height: "480px",
          width: "100%",
          borderRadius: "8px",
          overflow: "hidden",
        }}
      />
    </div>
  );
}

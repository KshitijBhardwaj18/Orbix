import Chart from "@/components/trade/Chart";
import Orderbook from "@/components/trade/Orderbook";
import TickerStats from "@/components/trade/TickerStats";
import TradeWindow from "@/components/trade/TradingWindow";

export default async function Trade({
  params,
}: {
  params: Promise<{ ticker: string }>;
}) {
  const { ticker } = await params;

  return (
    <div className="bg-secondary h-screen p-5">
      <div className="flex flex-row gap-2">
        <div className="w-9/10 flex flex-col gap-3">
          <TickerStats ticker={ticker} />
          <div className="flex flex-row gap-3">
            <div className="w-3/4">
              <Chart ticker={ticker} />
            </div>
            <div className="w-2/8">
              <Orderbook ticker={ticker} />
            </div>
          </div>
        </div>

        <div className="w-2/8">
            <TradeWindow ticker={ticker}/>
        </div>
      </div>
    </div>
  );
}

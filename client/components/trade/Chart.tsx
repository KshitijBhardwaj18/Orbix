import React from "react";
import { TradeView } from "../chart/chart";

interface ChartProps {
  ticker: string;
}

const Chart = ({ ticker }: { ticker: string }) => {
  return (
    <div className="w-full bg-primary h-[600px] rounded-xl">
      <TradeView market={ticker} />
    </div>
  );
};
export default Chart;

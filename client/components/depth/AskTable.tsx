import { useEffect, useRef } from "react";

interface AskTableProps {
  asks: [string, string][];
  precision?: number;
}

export const AskTable = ({ asks, precision = 0.00001 }: AskTableProps) => {
  const scrollRef = useRef<HTMLDivElement>(null);

  // Auto-scroll to bottom whenever asks data changes
  useEffect(() => {
    if (scrollRef.current) {
      scrollRef.current.scrollTop = scrollRef.current.scrollHeight;
    }
  }, [asks]);

  if (!asks || asks.length === 0) {
    return (
      <div className="p-4 text-center text-gray-500 text-sm">No ask orders</div>
    );
  }

  let currentTotal = 0;
  // Take more items to ensure scrolling
  const relevantAsks = asks.slice(0, 25);

  // Sort asks by price (lowest first) then reverse to show highest at top
  const sortedAsks = relevantAsks
    .sort(([a], [b]) => parseFloat(a) - parseFloat(b))
    .reverse(); // Highest asks at top, lowest asks at bottom (closest to current price)

  const asksWithTotal: [string, string, number][] = sortedAsks.map(
    ([price, quantity]) => {
      currentTotal += Number(quantity);
      return [price, quantity, currentTotal];
    },
  );

  const maxTotal = sortedAsks.reduce(
    (acc, [_, quantity]) => acc + Number(quantity),
    0,
  );

  return (
    <div
      ref={scrollRef}
      className="w-full h-full overflow-y-auto custom-scrollbar"
    >
      {asksWithTotal.map(([price, quantity, total], index) => (
        <Ask
          maxTotal={maxTotal}
          key={`${price}-${index}`}
          price={price}
          quantity={quantity}
          total={total}
          precision={precision}
        />
      ))}
    </div>
  );
};

function Ask({
  price,
  quantity,
  total,
  maxTotal,
  precision,
}: {
  price: string;
  quantity: string;
  total: number;
  maxTotal: number;
  precision: number;
}) {
  const formatPrice = (price: string) => {
    const num = parseFloat(price);
    if (isNaN(num)) return "0.00000";
    // Calculate decimal places based on precision
    const decimalPlaces =
      precision < 0.001 ? 5 : precision < 0.01 ? 4 : precision < 0.1 ? 3 : 2;
    return num.toFixed(decimalPlaces);
  };

  const formatQuantity = (quantity: string) => {
    const num = parseFloat(quantity);
    if (isNaN(num)) return "0.00";
    if (num >= 1000) return `${(num / 1000).toFixed(2)}K`;
    return num.toFixed(2);
  };

  const formatTotal = (total: number) => {
    if (isNaN(total)) return "0.00";
    if (total >= 1000) return `${(total / 1000).toFixed(2)}K`;
    return total.toFixed(2);
  };

  return (
    <div className="relative group hover:bg-gray-800/30 transition-colors cursor-pointer h-6 w-full flex-shrink-0">
      {/* Background bar for volume visualization */}
      <div
        className="absolute inset-0 bg-red-500/10 transition-all duration-300"
        style={{
          width: `${maxTotal > 0 ? (100 * total) / maxTotal : 0}%`,
        }}
      />

      {/* Content */}
      <div className="relative px-4 py-0.5 flex justify-between items-center text-sm h-full w-full">
        <div className="text-red-400 font-mono w-20 text-left">
          {formatPrice(price)}
        </div>
        <div className="text-white font-mono w-16 text-right">
          {formatQuantity(quantity)}
        </div>
        <div className="text-gray-400 font-mono w-16 text-right">
          {formatTotal(total)}
        </div>
      </div>

      {/* Hover effect border */}
      <div className="absolute right-0 top-0 bottom-0 w-1 bg-red-400 opacity-0 group-hover:opacity-100 transition-opacity" />
    </div>
  );
}

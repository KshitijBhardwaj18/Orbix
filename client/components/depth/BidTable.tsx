interface BidTableProps {
  bids: [string, string][];
  precision?: number;
}

export const BidTable = ({ bids, precision = 0.00001 }: BidTableProps) => {
  if (!bids || bids.length === 0) {
    return (
      <div className="p-4 text-center text-gray-500 text-sm">No bid orders</div>
    );
  }

  let currentTotal = 0;
  // Take more items to ensure scrolling
  const relevantBids = bids.slice(0, 25);

  // Sort bids by price (highest first) - highest bids should be closest to current price
  const sortedBids = relevantBids.sort(
    ([a], [b]) => parseFloat(b) - parseFloat(a),
  );

  const bidsWithTotal: [string, string, number][] = sortedBids.map(
    ([price, quantity]) => {
      currentTotal += Number(quantity);
      return [price, quantity, currentTotal];
    },
  );

  const maxTotal = sortedBids.reduce(
    (acc, [_, quantity]) => acc + Number(quantity),
    0,
  );

  return (
    <div className="w-full">
      {bidsWithTotal.map(([price, quantity, total], index) => (
        <Bid
          maxTotal={maxTotal}
          total={total}
          key={`${price}-${index}`}
          price={price}
          quantity={quantity}
          precision={precision}
        />
      ))}
    </div>
  );
};

function Bid({
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
    return num.toFixed(2); // Show 2 decimals for readability
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
        className="absolute inset-0 bg-green-500/10 transition-all duration-300"
        style={{
          width: `${maxTotal > 0 ? (100 * total) / maxTotal : 0}%`,
        }}
      />

      {/* Content */}
      <div className="relative px-4 py-0.5 flex justify-between items-center text-sm h-full w-full">
        <div className="text-green-400 font-mono w-20 text-left">
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
      <div className="absolute right-0 top-0 bottom-0 w-1 bg-green-400 opacity-0 group-hover:opacity-100 transition-opacity" />
    </div>
  );
}

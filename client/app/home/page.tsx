import MarketTable from "@/components/home/MarketTable";
import React from "react";
import MarketTrends from "@/components/home/MarketTrends";
import Thumbnail from "@/components/home/Thumbnail";

export default function Home() {
  return (
    <div className="bg-secondary h-full px-25 py-5">
      <Thumbnail/>  
      <MarketTrends />
      <MarketTable />
    </div>
  );
}

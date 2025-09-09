import React from 'react'

interface OrderbookProps {
    ticker: string 
}

const Orderbook = ({ticker}: {ticker:string}) => {
  return (
    <div className="h-[600px] bg-primary w-full rounded-xl ">Orderbook</div>
  )
}

export default Orderbook
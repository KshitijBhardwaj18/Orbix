import React from 'react'

interface ChartProps {
    ticker : string;
}

const Chart = ({ticker}:{ticker:string}) => {
  return (
    <div className='w-full bg-primary h-[600px] rounded-xl'>Chart</div>
  )
}

export default Chart
import {
  ColorType,
  createChart,
  CrosshairMode,
  UTCTimestamp,
} from "lightweight-charts";

export class ChartManager {
  private candleSeries: any;
  private lastUpdateTime: number = 0;
  private chart: any;
  private currentBar: {
    open: number | null;
    high: number | null;
    low: number | null;
    close: number | null;
  } = {
    open: null,
    high: null,
    low: null,
    close: null,
  };

  constructor(
    ref: HTMLElement,
    initialData: any[],
    layout: { background: string; color: string },
  ) {
    this.chart = createChart(ref, {
      autoSize: true,
      overlayPriceScales: {
        ticksVisible: true,
        borderVisible: true,
      },
      crosshair: {
        mode: CrosshairMode.Normal,
      },
      rightPriceScale: {
        visible: true,
        ticksVisible: true,
        entireTextOnly: true,
      },
      grid: {
        horzLines: {
          visible: false,
        },
        vertLines: {
          visible: false,
        },
      },
      layout: {
        background: {
          type: ColorType.Solid,
          color: layout.background,
        },
        textColor: "white",
      },
    });

    // Use addCandlestickSeries (without the 's')
    this.candleSeries = this.chart.addCandlestickSeries({
      upColor: "#26a69a",
      downColor: "#ef5350",
      borderVisible: false,
      wickUpColor: "#26a69a",
      wickDownColor: "#ef5350",
    });

    // Set initial data
    if (initialData && initialData.length > 0) {
      this.candleSeries.setData(
        initialData.map((data: any) => ({
          time: (data.timestamp / 1000) as UTCTimestamp,
          open: data.open,
          high: data.high,
          low: data.low,
          close: data.close,
        })),
      );
    }
  }

  public update(updatedPrice: {
    close: number;
    low: number;
    high: number;
    open: number;
    newCandleInitiated?: boolean;
    time?: number;
  }) {
    if (!this.lastUpdateTime) {
      this.lastUpdateTime = new Date().getTime();
    }

    this.candleSeries.update({
      time: (this.lastUpdateTime / 1000) as UTCTimestamp,
      close: updatedPrice.close,
      low: updatedPrice.low,
      high: updatedPrice.high,
      open: updatedPrice.open,
    });

    if (updatedPrice.newCandleInitiated && updatedPrice.time) {
      this.lastUpdateTime = updatedPrice.time;
    }
  }

  public destroy() {
    this.chart.remove();
  }
}

import { api } from "@/lib/axios";
import { Depth } from "@/types/depth";


export const getDepth = async (market:string) : Promise<Depth> => {

    market = market.replace("/","_")

    const response = await api.get<Depth>(`/market/getdepth/${market}`);

    return response.data
    
}
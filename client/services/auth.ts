import { RegisterRequest, RegisterResponse } from "@/types/auth";
import { api } from "@/lib/axios";

export async function registerUser(data:RegisterRequest){
    const response =  await api<RegisterResponse>("auth/register",{data})
    return response.data
}
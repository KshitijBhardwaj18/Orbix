import { RegisterRequest, RegisterResponse } from "@/types/auth";
import { api } from "@/lib/axios";

export interface User {
    username: string;
    email: string;

}

export async function registerUser(data:RegisterRequest){
    const response =  await api.post<RegisterResponse>("auth/register",data)
    return response;
}

export async function getCurrentUser(){
    const response = await api.get<User>("/user/me")
    if(response.data && 'error' in response.data){
        return null;
    }
    return response.data;
}

export async function logoutUser(){
    const response = await api.post("/auth/logout")
    return response
}
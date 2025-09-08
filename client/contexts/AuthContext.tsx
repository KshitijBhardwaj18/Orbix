'use client';
import {createContext, useContext, useEffect, useState} from 'react';
import {User, getCurrentUser, logoutUser} from "@/services/auth"

interface AuthContextType {
    user: User | null;
    loading: boolean;
    isAuthenticated: boolean;
    logout: () => Promise<void>;
    refetch:() => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({children}: {children: React.ReactNode}) {
    const [user, setUser] = useState<User | null>(null);
    const [loading,setLoading] = useState(true);

    const checkAuthStatus = async () => {
        try{
            const userData = await getCurrentUser();
            
            setUser(userData)
        }catch{
            setUser(null);
        }finally{
            setLoading(false)
        }
    };

    const logout = async () => {
        try {
            await logoutUser();
            setUser(null);
            window.location.href = '/signin'
        }catch(error) {
            console.log('Logout Failed', error)
        }
    };

    useEffect(() => {
        checkAuthStatus()
    }, [])

    return (
        <AuthContext.Provider value={{user,loading,isAuthenticated: !!user,logout,refetch: checkAuthStatus}}>
            {children}
        </AuthContext.Provider>
    )
}

export function useAuth() {
    const context = useContext(AuthContext)
    if(!context){
        throw new Error('useAuth must be used with in Auth Provider')
    }
    return context;
}
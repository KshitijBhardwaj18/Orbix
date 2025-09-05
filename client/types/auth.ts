export type RegisterResponse = {
    message: string;
    user_id: string;
}

export type RegisterRequest = {
    username: string;
    email: string;
    password: string;
}

export type LoginResponse = {
    email: string;
    password: string;
}
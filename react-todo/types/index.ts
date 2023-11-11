export type Task = {
    id: number;
    tilte: string;
    created_at: string;
    updated_at: string;
}

export type CsrfToken = {
    crsf_token: string; 
}

export type Credentials = {
    email: string;
    password: string;
}

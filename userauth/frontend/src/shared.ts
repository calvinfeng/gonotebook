
export type Message = {
    id: number
    body: string
    sender: User
    receiver: User
}

export type User = {
    id: string
    name: string
    email: string
    jwt: string
}
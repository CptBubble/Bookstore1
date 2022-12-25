export interface IBook {
    UUID: string
    Name: string
    Saleprice: number
    Year: number
    Type: string
    Srokgodnost: number
    Color: string
    Description: string
    Image: string
}

export interface ICart {
    BookUUID: string
    Quantity: number
}
export interface IOrder {
    UUID: string
    Book: string
    Quantity: number
    UserUUID: string
    Date: string
    Status: string
}
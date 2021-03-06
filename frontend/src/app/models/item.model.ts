export interface ItemModel {
    ID: string
    Title: string
    Subscription: string
    SubscriptionName: string
    Description?: string
    Link: string
    Published: Date
    New: boolean
    Saved: boolean
    Provider: string
    Type: string
    Thumbnail: string
    Embed: string
    Favorite: boolean
}

export type ItemCollection = ItemModel[];

export interface PaginatedItemCollection {
    Page: number
    Total: number
    Items: ItemCollection
}

export interface PaginatedRequest {
    Page: number
    ItemsPerPage: number
    NodeID?: string
    LeafIDs?: string[]
    Filter?: FilterRequest
}

export interface FilterRequest {
    New?: boolean
    Favorite?: boolean
    Type?: string[]
    Period?: {
        Start?: Date
        End?: Date
    }
}
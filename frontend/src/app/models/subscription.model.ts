import { LeafModel } from 'src/treex/model';


// SubscriptionModel ...
export interface SubscriptionModel {
    ID: string
    XURL?: string
    Title: string
    Description: string
    URL?: string
    LastUpdate?: Date
}


export const subscriptionToLeaf = (_l: SubscriptionModel): LeafModel => {
    return {
        id: _l.ID,
        label: _l.Title,
        description: _l.Description,
    } as LeafModel
}

export const leafToSubscription = (l: LeafModel): SubscriptionModel => {
    return {
        ID: l.id,
        Description: l.description,
        Title: l.label
    }
}


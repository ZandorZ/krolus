export interface LeafModel {
    id?: string
    parent?: string
    label?: string
    description?: string
    icon?: string
    color?: string
    items_count?: number
    new_items_count?: number
}

export interface NodeModel extends LeafModel {
    nodes_count: number
    children: NodeModel[]
    leaves?: LeafModel[]
    expanded: boolean
}

export type TreexNode = NodeModel | LeafModel;

export type LoadingDictionary = Dictionary<boolean>;

export type Dictionary<T> = {
    [id: string]: T
}


export const isNode = (data: TreexNode): data is NodeModel => {
    return (<NodeModel>data).nodes_count !== undefined;
}

export const isDescendent = (node: NodeModel, id: string): boolean => {
    if (!!node.children && !!node.children.find(n => n.id == id)) {
        return true;
    }

    for (let i in node.children) {
        if (isDescendent(node.children[i], id)) {
            return true
        }
    }

    return false;
}
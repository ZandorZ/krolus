import { TreexNodeHeader } from "../state/store";


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

export const getHeadersFromPath = (node: NodeModel, path: string): TreexNodeHeader[] => {

    let local: any[] = [];
    path.split(".children.").slice(1).forEach((id) => {
        if (id.includes(".leaves.")) {
            const temp = id.split('.leaves.');
            const parent = node.children.find(n => n.id == temp[0]);
            local.push({ id: parent.id, label: parent.label, leaf: false });
            const leaf = parent.leaves.find(l => l.id == temp[1]);
            local.push({ id: leaf.id, label: leaf.label, leaf: true });
        } else {
            const n = node.children.find(n => n.id == id);
            local.push({ id: n.id, label: n.label, leaf: false });
            node = n;
        }
    });

    return local;
}
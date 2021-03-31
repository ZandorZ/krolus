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

export const getPath = (parent: NodeModel, id: string): string => {

    if (!!parent.children && !!parent.children.find(n => n.id == id)) {
        return ".children." + id;
    }

    if (!!parent.leaves && !!parent.leaves.find(l => l.id == id)) {
        return ".leaves." + id;
    }

    if (!!parent.children) {
        for (let i = 0; i < parent.children.length; i++) {
            const path = getPath(parent.children[i], id);
            if (path.length > 0) {
                return ".children." + parent.children[i].id + path;
            }
        }
    }

    return "";
}

export const getHeadersFromPath = (node: NodeModel, path: string): TreexNodeHeader[] => {

    let headers: TreexNodeHeader[] = [
        {
            id: node.id,
            label: node.label,
            leaf: false,
            path: '/',
            description: '/'
        }
    ];
    let local = "root";

    path.split(".children.").slice(1).forEach((id) => {
        if (id.includes(".leaves.")) {
            const temp = id.split('.leaves.');
            const parent = node.children.find(n => n.id == temp[0]);
            local += ".children." + temp[0];
            headers.push({ id: parent.id, label: parent.label, description: parent.description, leaf: false, path: local });
            const leaf = parent.leaves.find(l => l.id == temp[1]);
            local += ".leaves." + temp[1];
            headers.push({ id: leaf.id, label: leaf.label, description: leaf.description, leaf: true, path: local });
        } else if (!!node.children) {
            const n = node.children.find(n => n.id == id);
            local += ".children." + id
            headers.push({ id: n.id, label: n.label, description: n.description, leaf: false, path: local });
            node = n;
        }
    });

    return headers;
}

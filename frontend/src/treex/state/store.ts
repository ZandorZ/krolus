import { LoadingDictionary, NodeModel, TreexNode } from '../model';


export type TreexNodeHeader = {
    id: string
    label: string
    description: string
    leaf: boolean
    path?: string
}

export class TreexState {
    root: NodeModel;
    selected: TreexNode;
    selectedPath: string;
    dragged: TreexNode;
    loading: LoadingDictionary = {};
}

export interface ITreexStore {
    //TODO: signature
}

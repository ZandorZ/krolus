import { LoadingDictionary, NodeModel, TreexNode } from '../model';


export type TreexNodeHeader = {
    id: string
    label: string
    leaf: boolean
}

export class TreexState {
    root: NodeModel;
    selected: TreexNode;
    selectedPath: string;
    selectedHeader: TreexNodeHeader[] = [];
    dragged: TreexNode;
    loading: LoadingDictionary = {};
}

export interface ITreexStore {

}

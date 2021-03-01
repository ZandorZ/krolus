import { LoadingDictionary, NodeModel, TreexNode } from '../model';


export class TreexState {
    root: NodeModel;
    selected: TreexNode;
    dragged: TreexNode;
    loading: LoadingDictionary = {};
}

export interface ITreexStore {

}

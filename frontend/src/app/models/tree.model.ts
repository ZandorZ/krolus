import { TreexNode } from 'src/treex/model';
import { NodeDictionary } from 'src/treex/state/store';
import { DirectoryCollection, DirectoryModel, directoryToNode } from './directory.model';
import { SubscriptionCollection, SubscriptionModel, subscriptionToLeaf } from './subscription.model';


export type DirectoryDictionary = {
    [path: string]: DirectoryCollection
}

export type SubscriptionDictionary = {
    [path: string]: SubscriptionCollection
}

export class TreeState {
    Selected: SubscriptionModel | DirectoryModel
    Directories: DirectoryDictionary
    Subscriptions: SubscriptionDictionary
}

export const convertStates = (state: TreeState): NodeDictionary => {

    const nodes = {};

    for (let path in state.Directories) {
        if (state.Directories[path] != null) {
            nodes[path] = state.Directories[path]
                .map(node => directoryToNode(node))
                .sort(sortNodes);
        }
    }

    for (let path in state.Subscriptions) {
        if (state.Subscriptions[path] != null) {
            const subs = state.Subscriptions[path]
                .map(node => subscriptionToLeaf(node))
                .sort(sortNodes);

            if (nodes[path]) {
                nodes[path] = [...nodes[path], ...subs];
            } else {
                nodes[path] = subs;
            }

        }
    }

    return nodes;
}


const sortNodes = (a: TreexNode, b: TreexNode) => {
    if (a.label.toLowerCase() > b.label.toLowerCase()) {
        return 1;
    }
    if (a.label.toLowerCase() < b.label.toLowerCase()) {
        return -1;
    }
    return 0;
};
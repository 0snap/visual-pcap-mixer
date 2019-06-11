function parseTraceFileDictionary(traceFiles, defaultAdded = false) {
    return Object.entries(traceFiles).reduce((acc, entry) => {
        const [id, tf] = entry;
        tf.added = defaultAdded;
        acc[id] = tf;
        return acc;
    }, {});
}

function traceFiles(state = {}, action) {
    const newState = {...state};
    switch (action.type) {
    case 'ADD_MSA_ENTRY':
        if (action.entry.type === 'traceFile') newState[action.entry.id].added = true;
        return newState;
    case 'REMOVE_MSA_ENTRY':
        if (newState[action.entryId]) newState[action.entryId].added = false;
        return newState;
    case 'LOAD_TRACEFILES_SUCCESS':
    return parseTraceFileDictionary(action.traceFiles);
    case 'LOAD_MSA_SUCCESS':
        const loadedTfsFromMsa = Object.entries(action.multistepattacks).reduce((acc, entry) => {
            const msa = entry[1];
            return {...acc, ...parseTraceFileDictionary(msa.traceFiles)};
        }, {});
        return {...newState, ...loadedTfsFromMsa};
    case 'CREATE_MSA_SUCCESS':
        return {...newState, ...parseTraceFileDictionary(action.msa.traceFiles)};
    case 'LOAD_MSA_TO_TIMELINE':
        return parseTraceFileDictionary(state);
    default:
        return state
    }
}
export default traceFiles